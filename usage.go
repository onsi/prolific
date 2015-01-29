package main

import (
	"fmt"
	"os"
)

func PrintUsageAndExit() {
	fmt.Println("Usage:")
	fmt.Println("prolific template\n  generates a sample stories.prolific")
	fmt.Println(`prolific "Author Name" "path/to/stories.prolific"\n  converts the passed in prolific file to a CSV printed to stdout`)
	os.Exit(1)
}
