package showcmd

import (
	"flag"
	"fmt"

	"github.com/projectops/ecsw/pkg/config"
)

// ShowCommand - the show flagset
type ShowCommand struct {
	fs *flag.FlagSet

	name string
}

// NewShowCmd - create the new show flagset
func NewShowCmd() *ShowCommand {
	cmd := &ShowCommand{
		fs: flag.NewFlagSet("show", flag.ContinueOnError),
	}

	cmd.fs.StringVar(&cmd.name, "name", "", "the workspace name")
	return cmd
}

// Name - return the command name
func (cmd *ShowCommand) Name() string {
	return cmd.fs.Name()
}

// Init - parse the command and subcommands
func (cmd *ShowCommand) Init(args []string) error {
	return cmd.fs.Parse(args)
}

// Run - execute the commands
func (cmd *ShowCommand) Run() error {
	workspace := config.NewConfig()

	fmt.Println("Current Workspace")
	fmt.Println("===================")
	fmt.Printf("Workspace: %s\n", workspace.CurrentWorkspaceName)
	fmt.Printf("Cluster: %s\n", workspace.CurrentWorkspace.Cluster)
	fmt.Printf("Region: %s\n", workspace.CurrentWorkspace.Region)

	return nil
}
