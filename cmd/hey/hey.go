package main

import (
	"context"
	"hey/internal/execution"
	"hey/internal/management"
	"log"
	"net/mail"
	"os"

	cli "github.com/urfave/cli/v3"
)

const (
	AppName        = "hey"
	AppVersion     = "0.1.0"
	AppDescription = "Easing CLI usage for executing tasks and managing configurations"
)

func main() {
	ctx := context.Background()

	app := &cli.Command{
		Name:    AppName,
		Usage:   AppDescription,
		Version: AppVersion,
		Authors: []any{
			mail.Address{Name: "Jasti Sri Radhe Shyam", Address: "samabhasatejsrs@outlook.com"},
		},
		Commands: []*cli.Command{
			runCommand(),
			importCommand(),
			exportCommand(),
		},
		// TODO: add remove edit, multiple, single, commands, import and export
		// import with overwrite whole dir, overwrite (conflicting files) and merge files
		// Configuration version should be maintained
		// multilevel config k8.po (under k8.yaml), default.ls (also can be done by ls)
		// System level config and user level config
		// implementing system level config will create an issue with merging system level config with user level config
		// ,this will create confusion to the user from which config it has been taken
	}

	if err := app.Run(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}

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

	var modulesPtr *execution.Modules
	modules := make(execution.Modules)
	modulesPtr = &modules
	args := cmd.Args().Get(0)
	modulesPtr.ParseModuleAndProcessTask(args)
	return nil
}

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

func generateUsageText(text string) string {
	return AppName + " " + text
}
