/*
Copyright 2023 Jasti Sri Radhe Shyam

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"hey/execution"
	"hey/management"
	"log"
	"net/mail"
	"os"

	cli "github.com/urfave/cli/v3"
)

func main() {
	ctx := context.Background()
	management.CheckAndInit()
	appName := "hey"

	generateUsageText := func(text string) string {
		return appName + " " + text
	}

	// in the cli config, "Usage" is short description (IDK why it's named Usage)
	app := &cli.Command{
		Name:    appName,
		Usage:   "Easing CLI usage",
		Version: "0.1.0",
		Authors: []any{
			mail.Address{Name: "Jasti Sri Radhe Shyam", Address: "samabhasatejsrs@outlook.com"},
		},
		Commands: []*cli.Command{
			{
				Name:        "run",
				Aliases:     []string{"r"},
				Usage:       "execute the task",
				UsageText:   generateUsageText("run moduleName.taskName"),
				Description: "run command will run the named command(s) as configured",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					var modulesPtr *execution.Modules
					modules := make(execution.Modules)
					modulesPtr = &modules
					args := cmd.Args().Get(0)
					modulesPtr.ParseModuleAndProcessTask(args)
					return nil
				},
			},
			{
				Name:      "import",
				Usage:     "imports configuration from a filepath",
				UsageText: generateUsageText("import filepath"),
				Action: func(ctx context.Context, cmd *cli.Command) error {
					importPath := cmd.Args().Get(0)
					err := management.Import(importPath)
					if err != nil {
						log.Fatal("Error : ", err)
					}
					return nil
				},
			},
			{
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
				Action: func(ctx context.Context, cmd *cli.Command) error {
					exportPath := cmd.Args().Get(0)
					// flags := c.App.Flags
					// c.Command.Flags
					fmt.Println(cmd.Args())
					excludeModules := cmd.StringSlice("exclude")
					fmt.Println(excludeModules)
					err := management.Export(exportPath, excludeModules)
					if err != nil {
						log.Fatal("Error : ", err)
					}
					return nil
				},
			},
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
