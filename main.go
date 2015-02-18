package main

/*
prolific: generate many tracker stories
*/

import (
	"os"
	"io/ioutil"

	"fmt"
)

func main() {
	content := readStdin()
	if len(os.Args) == 1 && content != nil {
		fmt.Fprintf(os.Stderr, "Converting STDIN\n")
		err := ConvertAndEmitStories(content)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed:", err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

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
	err := ConvertAndEmitStoriesFromFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed:", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func readStdin() []byte {
	stat, err := os.Stdin.Stat()
	if err != nil || (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil
	}

	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil
	}
	return stdin
}
