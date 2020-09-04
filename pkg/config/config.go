package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	// HOMEPATH - application config file
	HOMEPATH = "/etc/ecsw/ecsw.yml"
	// APPHOME - folder with workspaces
	APPHOME = "/etc/ecsw/"
)

// Workspace - workspace configuration
type Workspace struct {
	Name    string `yaml:"name"`
	Cluster string `yaml:"cluster"`
	Region  string `yaml:"region"`
}

// Config - configuration map
type Config struct {
	homePath string
	appHome  string

	CurrentWorkspace     Workspace
	CurrentWorkspaceName string `yaml:"workspace"`
}

// CreateWorkspace - write workspace file
func CreateWorkspace(name, cluster, region string) error {
	workspace := Workspace{
		Name:    name,
		Cluster: cluster,
		Region:  region,
	}

	yaml, err := yaml.Marshal(&workspace)
	fmt.Println(string(yaml))
	if err != nil {
		return err
	}

	yamlFile := fmt.Sprintf("%s/%s.yml", APPHOME, name)

	f, err := os.Create(yamlFile)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(yaml); err != nil {
		return err
	}

	return nil
}

// ChangeWorkspace - change current workspace
func ChangeWorkspace(name string) error {
	f, err := os.OpenFile(HOMEPATH, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	config := Config{
		CurrentWorkspaceName: name,
	}

	yaml, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	if _, err := f.Write(yaml); err != nil {
		return err
	}

	return nil
}

func getWorkspace(file string) (string, error) {
	var conf Config
	var workspace string

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return workspace, err
	}

	if err := yaml.Unmarshal(yamlFile, &conf); err != nil {
		return workspace, err
	}

	workspace = conf.CurrentWorkspaceName

	return workspace, nil
}

func readClusterFile(file string) (Workspace, error) {
	var workspace Workspace

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return workspace, err
	}

	if err := yaml.Unmarshal(yamlFile, &workspace); err != nil {
		return workspace, err
	}

	return workspace, nil
}

// NewConfig - create new config
func NewConfig() *Config {
	conf := &Config{
		homePath: HOMEPATH,
		appHome:  APPHOME,
	}

	// get current workspace
	wkName, err := getWorkspace(conf.homePath)
	if err != nil {
		fmt.Printf("Unable to open workspace file reason: %s\n", err)
		panic(err)
	}
	conf.CurrentWorkspaceName = wkName

	wkFile := fmt.Sprintf("%s/%s.yml", conf.appHome, conf.CurrentWorkspaceName)

	// add config
	wkConf, err := readClusterFile(wkFile)
	if err != nil {
		fmt.Printf("Unable to use workspace reason: %s\n", err)
		panic(err)
	}
	conf.CurrentWorkspace = wkConf

	return conf
}
