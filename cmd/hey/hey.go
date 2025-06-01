package main

import (
	"context"
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

func generateUsageText(text string) string {
	return AppName + " " + text
}
