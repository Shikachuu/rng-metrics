package main

import (
	"os"

	"github.com/Shikachuu/podman-example-monitoring/randomMetricsGenerator/cmd"
)

func main() {
	c := cmd.RootCommand()

	err := c.Execute()

	if err != nil {
		os.Exit(1)
	}
}
