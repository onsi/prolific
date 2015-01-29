package main

import (
	"fmt"
	"os"
)

func PrintUsageAndExit() {
	fmt.Println(`prolific v2.0

Usage:
    prolific FILE
        converts the passed in prolific file to a CSV printed to stdout
        the CSV can be imported manually into Pivotal Tracker

        prolific stories.prolific > stories.csv

        is a useful one-liner

    prolific template
        generates a sample stories.prolific

    prolific help
        you're looking at it!
	
Syntax:
    Stories are separated by the delimiter:

    ---

    Each story is a block made up of:

    [STORY_TYPE] TITLE

    DESCRIPTION
    DESCRIPTION
    DESCRIPTION

    L: LABEL 1, LABEL 2

    Of these, only TITLE is required and must be on a single line.

    [STORY_TYPE] may be omitted.  If present it must be one of:
      [FEATURE]
      [BUG]
      [CHORE]
      [RELEASE]`)
	os.Exit(1)
}
