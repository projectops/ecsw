package scalecmd

import (
	"errors"
	"flag"

	"github.com/projectops/ecsw/pkg/ecs"

	"github.com/projectops/ecsw/pkg/config"
)

// ScaletCommand - the flagset scale
type ScaletCommand struct {
	fs *flag.FlagSet

	service string
	tasks   int64
}

// NewScaleCMD - create the new scale flagset
func NewScaleCMD() *ScaletCommand {
	cmd := &ScaletCommand{
		fs: flag.NewFlagSet("scale", flag.ContinueOnError),
	}

	cmd.fs.StringVar(&cmd.service, "service", "", "service name.")
	cmd.fs.Int64Var(&cmd.tasks, "tasks", 1, "number of tasks to scale.")

	return cmd
}

// Name - return the command name
func (cmd *ScaletCommand) Name() string {
	return cmd.fs.Name()
}

// Init - parse the command and subcommands
func (cmd *ScaletCommand) Init(args []string) error {
	return cmd.fs.Parse(args)
}

// Run - execute the commands
func (cmd *ScaletCommand) Run() error {
	workspace := config.NewConfig()

	cluster := workspace.CurrentWorkspace.Cluster
	region := workspace.CurrentWorkspace.Region

	if cmd.service == "" {
		return errors.New("service arguments is required")
	}

	if err := ecs.ScaleTask(cluster, cmd.service, region, cmd.tasks); err != nil {
		return err
	}

	return nil
}
