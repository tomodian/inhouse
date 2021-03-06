package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tomodian/inhouse"
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
				Aliases: []string{"f", "functions", "funcs"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    dirFlag,
						Usage:   fmt.Sprintf("target `%s`", strings.ToUpper(dirFlag)),
						Aliases: []string{"d"},
					},
					&cli.StringFlag{
						Name: formatFlag,
						Usage: fmt.Sprintf(
							"select output `FORMAT` (%s, %s, %s, %s)",
							inhouse.ColonFormat.String(),
							inhouse.CSVFormat.String(),
							inhouse.TSVFormat.String(),
							inhouse.JSONFormat.String(),
						),
						DefaultText: inhouse.ColonFormat.String(),
						Aliases:     []string{"f"},
					},
					&cli.BoolFlag{
						Name:    exitFlag,
						Usage:   "terminate with exit code 1 on match",
						Aliases: []string{"e"},
					},
					&cli.BoolFlag{
						Name:    listFlag,
						Usage:   "list all functions and quit",
						Aliases: []string{"l"},
					},
					&cli.BoolFlag{
						Name:  flatFlag,
						Usage: "search files in flat directory, defaults to recursive",
						Value: false,
					},
					&cli.BoolFlag{
						Name:    testFlag,
						Usage:   "include test files to search target",
						Value:   false,
						Aliases: []string{"t"},
					},
				},
				Action: func(c *cli.Context) error {
					dir := c.String(dirFlag)
					format := c.String(formatFlag)
					name := c.Args().Get(0)
					flat := !c.Bool(flatFlag)

					var check *inhouse.Check

					switch {
					case c.Bool(testFlag):
						got, err := inhouse.Contains(dir, name, flat)

						if err != nil {
							return err
						}

						check = got

					default:
						got, err := inhouse.SourcesContains(dir, name, flat)

						if err != nil {
							return err
						}

						check = got
					}

					if c.Bool(listFlag) {
						for _, c := range check.Combine() {
							fmt.Println(c.Format(inhouse.CodeFormat(format)))
						}

						return nil
					}

					for _, c := range check.Matches {
						fmt.Println(c.Format(inhouse.CodeFormat(format)))
					}

					if check.Contained && c.Bool(exitFlag) {
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
