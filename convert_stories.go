package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func ConvertAndEmitStoriesFromFile(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Couldn't load file: %s\n%s", file, err.Error())
	}

	return ConvertAndEmitStories(string(content))
}

func ConvertAndEmitStories(content string) error {
	stories, numTasks, errors := ExtractStories(content)

	if len(errors) > 0 {
		fmt.Fprintln(os.Stderr, "There were errors parsing your file:")
		for _, err := range errors {
			fmt.Fprintln(os.Stderr, "- "+err.Error())
		}
	}

	w := csv.NewWriter(os.Stdout)

	headers := BASE_CSV_HEADERS
	for i := 0; i < numTasks; i++ {
		headers = append(headers, "Task")
	}

	w.Write(headers)

	for _, story := range stories {
		w.Write(story.CSVRecords(numTasks))
	}

	w.Flush()

	return nil
}

var EmptyStoryError = errors.New("You have an empty story.")

func ExtractStories(content string) ([]Story, int, []error) {
	errors := []error{}

	story_separator := regexp.MustCompile(`(?m)(\n\n|\A)----\s*\n\s*`)

	parts := story_separator.Split(content, -1)

	numTasks := 0

	stories := []Story{}
	for _, part := range parts {
		story, err := ExtractStory(part)
		if err != nil {
			if err != EmptyStoryError {
				errors = append(errors, err)
			}
			continue
		}
		if len(story.Tasks) > numTasks {
			numTasks = len(story.Tasks)
		}
		stories = append(stories, story)
	}

	return stories, numTasks, errors
}

func isStoryEmpty(content string) bool {
	match, _ := regexp.MatchString(`(?m)\A\s*\z`, content)
	return match
}

func ExtractStory(part string) (Story, error) {
	if isStoryEmpty(part) {
		return Story{}, EmptyStoryError
	}

	lines := strings.Split(part, "\n")
	story := &Story{}

	for _, line := range lines {
		err := story.AppendLine(string(line))
		if err != nil {
			return Story{}, err
		}
	}

	return *story, nil
}
