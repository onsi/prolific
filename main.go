package main

/*
prolific: generate many tracker stories
*/

import (
	"os"

	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		PrintUsageAndExit()
	}

	if os.Args[1] == "help" {
		PrintUsageAndExit()
	}

	if os.Args[1] == "template" {
		fmt.Println("Writing template to stories.prolific")
		err := GenerateTemplate()
		if err != nil {
			fmt.Println("Failed:", err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	fmt.Fprintf(os.Stderr, "Converting %s\n", os.Args[1])
	err := ConvertAndEmitStories(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed:", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
