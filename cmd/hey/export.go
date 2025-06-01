package main

import (
	"context"
	"hey/internal/management"

	cli "github.com/urfave/cli/v3"
)

func exportCommand() *cli.Command {
	return &cli.Command{
		Name:      "export",
		Usage:     "exports configuration (in tar file with gz compression) to current directory or to specified path",
		UsageText: generateUsageText("export export_path"),
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "exclude",
				Usage: "exclude modules from the exporter archive",
				Value: []string{},
			},
		},
		Action: exportAction,
	}
}

func exportAction(ctx context.Context, cmd *cli.Command) error {
	management.CheckAndInit()

	exportPath := cmd.Args().Get(0)
	excludeModules := cmd.StringSlice("exclude")
	return management.Export(exportPath, excludeModules)
}
