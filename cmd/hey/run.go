package main

import (
	"context"
	"hey/internal/execution"
	"hey/internal/management"

	cli "github.com/urfave/cli/v3"
)

func runCommand() *cli.Command {
	return &cli.Command{
		Name:        "run",
		Aliases:     []string{"r"},
		Usage:       "execute the task",
		UsageText:   generateUsageText("run moduleName.taskName"),
		Description: "run command will run the named command(s) as configured",
		Action:      runAction,
	}
}

func runAction(ctx context.Context, cmd *cli.Command) error {
	management.CheckAndInit()

	modules := new(execution.Modules)
	args := cmd.Args().Get(0)
	modules.ParseModuleAndProcessTask(args)
	return nil
}
