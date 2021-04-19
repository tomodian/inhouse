package main

import (
	"fmt"
	"os"

	"github.com/tomodian/inhouse"
)

func main() {
	os.Setenv(inhouse.CLIENV, "yes")

	if err := run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
