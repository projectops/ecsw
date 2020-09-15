package main

import (
	"fmt"
	"os"

	"github.com/projectops/ecsw/cmd"
)

func main() {
	if err := cmd.Root(os.Args[1:]); err != nil {
		fmt.Println(err)
	}
}
