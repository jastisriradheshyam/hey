package main

import (
	"fmt"
	"hey/execution"
	"hey/management"
	"hey/misc"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	management.CheckAndInit()
	app := &cli.App{
		Name:    "hey",
		Usage:   "Easing CLI usage",
		Version: "0.1.0",
		Authors: []*cli.Author{
			{Name: "Jasti Sri Radhe Shyam", Email: "samabhasatejsrs@outlook.com"},
		},
		Commands: []*cli.Command{
			{
				Name:        "run",
				Aliases:     []string{"r"},
				Usage:       "use to run a set of command(s)",
				Description: "run command will run the named command(s) as configured",
				Action: func(c *cli.Context) error {
					var modulesPtr *execution.Modules
					modules := make(execution.Modules)
					modulesPtr = &modules

					args := c.Args().Get(0)
					parsedRunArgs := misc.ParseRunArguments(args)
					fmt.Println(len(parsedRunArgs), parsedRunArgs)
					if len(parsedRunArgs) <= 0 {
						log.Fatal("Run Arguments are not present")
						os.Exit(1)
					}
					module, taskName := misc.GetModuleAndCommandName(parsedRunArgs)
					modulesPtr.ProcessTask(module, taskName)
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

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
