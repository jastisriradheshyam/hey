package main

import (
	"context"
	"hey/internal/management"

	cli "github.com/urfave/cli/v3"
)

func importCommand() *cli.Command {
	return &cli.Command{
		Name:      "import",
		Usage:     "imports configuration from a filepath",
		UsageText: generateUsageText("import filepath"),
		Action:    importAction,
	}
}

func importAction(ctx context.Context, cmd *cli.Command) error {
	management.CheckAndInit()

	importPath := cmd.Args().Get(0)
	return management.Import(importPath)
}
