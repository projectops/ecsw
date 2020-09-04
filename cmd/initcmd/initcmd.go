package initcmd

import (
	"errors"
	"flag"
	"fmt"

	"github.com/projectops/ecsw/pkg/config"
)

// InitCommand - the flagset init
type InitCommand struct {
	fs *flag.FlagSet

	name    string
	cluster string
	region  string
}

// NewInitCmd - create the new init flagset
func NewInitCmd() *InitCommand {
	cmd := &InitCommand{
		fs: flag.NewFlagSet("init", flag.ContinueOnError),
	}

	cmd.fs.StringVar(&cmd.name, "name", "", "the workspace name")
	cmd.fs.StringVar(&cmd.cluster, "cluster", "", "the cluster name")
	cmd.fs.StringVar(&cmd.region, "region", "us-east-2", "the aws region name")

	return cmd
}

// Name - return the command name
func (cmd *InitCommand) Name() string {
	return cmd.fs.Name()
}

// Init - parse the command and subcommands
func (cmd *InitCommand) Init(args []string) error {
	return cmd.fs.Parse(args)
}

// Run - execute the commands
func (cmd *InitCommand) Run() error {
	name := cmd.name
	cluster := cmd.cluster
	region := cmd.region

	if name == "" || cluster == "" || region == "" {
		return errors.New("Some argument is incorrect")
	}

	fmt.Printf("Creating workspace: %s\n", cmd.name)
	if err := config.CreateWorkspace(name, cluster, region); err != nil {
		return err
	}
	return nil
}
