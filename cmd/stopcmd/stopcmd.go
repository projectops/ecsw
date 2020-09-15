package stopcmd

import (
	"errors"
	"flag"

	"github.com/projectops/ecsw/pkg/ecs"

	"github.com/projectops/ecsw/pkg/config"
)

// StopCommand - the flagset stop
type StopCommand struct {
	fs *flag.FlagSet

	service string
}

// NewStopCMD - create the new stop flagset
func NewStopCMD() *StopCommand {
	cmd := &StopCommand{
		fs: flag.NewFlagSet("stop", flag.ContinueOnError),
	}

	cmd.fs.StringVar(&cmd.service, "service", "", "the service name.")

	return cmd
}

// Name - return the command name
func (cmd *StopCommand) Name() string {
	return cmd.fs.Name()
}

// Init - parse the command and subcommands
func (cmd *StopCommand) Init(args []string) error {
	return cmd.fs.Parse(args)
}

// Run - execute the commands
func (cmd *StopCommand) Run() error {
	workspace := config.NewConfig()

	cluster := workspace.CurrentWorkspace.Cluster
	region := workspace.CurrentWorkspace.Region

	if cmd.service == "" {
		return errors.New("service arguments is required")
	}

	if err := ecs.StopTask(cluster, cmd.service, region); err != nil {
		return err
	}

	return nil
}
