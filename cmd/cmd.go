package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/projectops/ecsw/cmd/scalecmd"

	"github.com/projectops/ecsw/cmd/listcmd"

	"github.com/projectops/ecsw/cmd/initcmd"
	"github.com/projectops/ecsw/cmd/selectcmd"
	"github.com/projectops/ecsw/cmd/showcmd"
)

// Runner -
type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

// Root -
func Root(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass a subcommand ()")
	}

	subcommand := os.Args[1]

	switch command := args[0]; command {
	case "init":
		cmds := []Runner{
			initcmd.NewInitCmd(),
		}

		for _, cmd := range cmds {
			if cmd.Name() == subcommand {
				cmd.Init(os.Args[2:])
				return cmd.Run()
			}
		}
	case "select":
		cmds := []Runner{
			selectcmd.NewSelectCmd(),
		}

		for _, cmd := range cmds {
			if cmd.Name() == subcommand {
				cmd.Init(os.Args[2:])
				return cmd.Run()
			}
		}
	case "show":
		cmds := []Runner{
			showcmd.NewShowCmd(),
		}

		for _, cmd := range cmds {
			if cmd.Name() == subcommand {
				cmd.Init(os.Args[2:])
				return cmd.Run()
			}
		}
	case "list":
		cmds := []Runner{
			listcmd.NewListCMD(),
		}

		for _, cmd := range cmds {
			if cmd.Name() == subcommand {
				cmd.Init(os.Args[2:])
				return cmd.Run()
			}
		}
	case "scale":
		cmds := []Runner{
			scalecmd.NewScaleCMD(),
		}

		for _, cmd := range cmds {
			if cmd.Name() == subcommand {
				cmd.Init(os.Args[2:])
				return cmd.Run()
			}
		}
	default:
		fmt.Println("Nothing!")
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}
