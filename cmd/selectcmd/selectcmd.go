package selectcmd

import (
	"flag"
	"fmt"

	"github.com/projectops/ecsw/pkg/config"
)

// SelectCommand - the flagset select
type SelectCommand struct {
	fs *flag.FlagSet

	workspace string
}

// NewSelectCmd - create the new select flagset
func NewSelectCmd() *SelectCommand {
	cmd := &SelectCommand{
		fs: flag.NewFlagSet("select", flag.ContinueOnError),
	}

	cmd.fs.StringVar(&cmd.workspace, "workspace", "", "workspace name")

	return cmd
}

// Name - return the command name
func (cmd *SelectCommand) Name() string {
	return cmd.fs.Name()
}

// Init - parse the command and subcommands
func (cmd *SelectCommand) Init(args []string) error {
	return cmd.fs.Parse(args)
}

// Run - execute the commands
func (cmd *SelectCommand) Run() error {
	fmt.Printf("Changing to workspace: %s\n", cmd.workspace)
	if err := config.ChangeWorkspace(cmd.workspace); err != nil {
		return err
	}
	return nil
}
