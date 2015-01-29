package main

/*
prolific: generate many tracker stories
*/

import (
	"os"

	"fmt"
)

func main() {
	if len(os.Args) == 1 {
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

	if len(os.Args) != 3 {
		PrintUsageAndExit()
	}

	author := os.Args[1]
	storyFile := os.Args[2]

	fmt.Fprintf(os.Stderr, "Converting %s\n", storyFile)
	fmt.Fprintf(os.Stderr, "Author will be %s\n", author)

	err := ConvertAndEmitStories(author, storyFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed:", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
