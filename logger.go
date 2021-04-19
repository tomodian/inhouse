package inhouse

import (
	"fmt"
	"os"
)

const (
	debug = "DEBUG"
)

// Log conditionally STDOUTs the given message, depending on DEBUG flag.
func Log(msg ...string) {
	if os.Getenv(debug) == "" {
		return
	}

	fmt.Print(debug, ": ")

	for i, m := range msg {
		fmt.Print(m)

		if i+1 != len(msg) {
			fmt.Printf(" ")
		}
	}

	fmt.Println("")
}
