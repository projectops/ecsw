package listcmd

import (
	"flag"
	"fmt"

	"github.com/projectops/ecsw/pkg/config"
	"github.com/projectops/ecsw/pkg/ecs"
)

// ListCommand - the flagset init
type ListCommand struct {
	fs *flag.FlagSet

	all bool
}

// NewListCMD - create the new list flagset
func NewListCMD() *ListCommand {
	cmd := &ListCommand{
		fs: flag.NewFlagSet("list", flag.ContinueOnError),
	}

	cmd.fs.BoolVar(&cmd.all, "all", false, "return more informations about services")

	return cmd
}

// Name - return the command name
func (cmd *ListCommand) Name() string {
	return cmd.fs.Name()
}

// Init - parse the command and subcommands
func (cmd *ListCommand) Init(args []string) error {
	return cmd.fs.Parse(args)
}

// Run - execute the commands
func (cmd *ListCommand) Run() error {
	workspace := config.NewConfig()

	cluster := workspace.CurrentWorkspace.Cluster
	region := workspace.CurrentWorkspace.Region

	services := ecs.GetServices(cluster, region)

	for _, item := range services {
		if cmd.all {
			fmt.Printf("Service: %s - ARN: %s - (%d/%d)\n", item.Name, item.ARN, item.RunningTasks, item.DesiredTasks)
		} else {
			fmt.Printf("Service: %s - (%d/%d)\n", item.Name, item.RunningTasks, item.DesiredTasks)
		}
	}

	return nil
}
