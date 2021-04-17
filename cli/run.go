package main

import (
	"fmt"
	"os"
	"strings"

	"inhouse"

	"github.com/urfave/cli/v2"
)

func run(args []string) error {
	app := &cli.App{
		Name:  "inhouse",
		Usage: "Go code regulation/convention checker 👀",
		Commands: []*cli.Command{
			{
				Name:    "function",
				Usage:   "search for function name",
				Aliases: []string{"f"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    dirFlag,
						Usage:   fmt.Sprintf("target `%s`", strings.ToUpper(dirFlag)),
						Aliases: []string{"d"},
					},
					&cli.BoolFlag{
						Name:    exitFlag,
						Usage:   "terminate with exit code 2 on match",
						Aliases: []string{"e"},
					},
					&cli.BoolFlag{
						Name:    listFlag,
						Usage:   "list all functions and quit",
						Aliases: []string{"l"},
					},
				},
				Action: func(c *cli.Context) error {
					dir := c.String(dirFlag)
					name := c.Args().Get(0)

					got, err := inhouse.SourcesContains(dir, name, true)

					if err != nil {
						return err
					}

					if c.Bool(listFlag) {
						for _, c := range got.Combine() {
							fmt.Printf("%s %s\n", c.Filepath, c.Function)
						}

						return nil
					}

					for _, c := range got.Matches {
						fmt.Printf("%s:%d\n", c.Filepath, c.Line)
					}

					if got.Contained && c.Bool(exitFlag) {
						os.Exit(1)
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(args); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		return err
	}

	return nil
}
