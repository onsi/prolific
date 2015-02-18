package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func ConvertAndEmitStoriesFromFile(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Couldn't load file: %s\n%s", file, err.Error())
	}

	return ConvertAndEmitStories(content)
}

func ConvertAndEmitStories(content []byte) error {
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

func ExtractStories(content []byte) ([]Story, int, []error) {
	errors := []error{}

	parts := bytes.Split(content, []byte("\n---\n\n"))

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

func ExtractStory(part []byte) (Story, error) {
	lines := bytes.Split(part, []byte("\n"))
	if len(lines) == 0 {
		return Story{}, EmptyStoryError
	}

	story := &Story{}

	for _, line := range lines {
		err := story.AppendLine(string(line))
		if err != nil {
			return Story{}, err
		}
	}

	return *story, nil
}
