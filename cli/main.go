package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
