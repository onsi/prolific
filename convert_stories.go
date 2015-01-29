package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ConvertAndEmitStories(author string, file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Couldn't load file: %s\n%s", file, err.Error())
	}

	stories := ExtractStories(content, author)

	w := csv.NewWriter(os.Stdout)

	w.Write(CSV_HEADERS)

	for _, story := range stories {
		w.Write(story.CSVRecords())
	}

	w.Flush()
	return nil
}

var CSV_HEADERS = []string{
	"Requested By",
	"Title",
	"Description",
	"Labels",
}

type Story struct {
	Author  string
	Title   string
	Content string
	Labels  []string
}

func (s *Story) AppendLine(line string) {
	if s.Title == "" {
		if line == "" {
			return
		}
		s.Title = line
		return
	}

	if strings.HasPrefix(line, "L: ") {
		labels := strings.Split(strings.TrimPrefix(line, "L: "), ",")
		for i := range labels {
			labels[i] = strings.TrimSpace(labels[i])
		}
		s.Labels = labels
		return
	}

	s.Content += line + "\n"
}

func (s *Story) FinalizeContents() {
	s.Content = strings.TrimSpace(s.Content)
}

func (s Story) CSVRecords() []string {
	return []string{
		s.Author,
		s.Title,
		s.Content,
		strings.Join(s.Labels, ","),
	}
}

func ExtractStories(content []byte, author string) []Story {
	parts := bytes.Split(content, []byte("\n---\n\n"))
	stories := []Story{}
	for _, part := range parts {
		story, err := ExtractStory(part, author)
		if err != nil {
			continue
		}
		stories = append(stories, story)
	}

	return stories
}

func ExtractStory(part []byte, author string) (Story, error) {
	lines := bytes.Split(part, []byte("\n"))
	if len(lines) == 0 {
		return Story{}, errors.New("empty story")
	}

	story := &Story{
		Author: author,
	}

	for _, line := range lines {
		story.AppendLine(string(line))
	}

	story.FinalizeContents()

	return *story, nil
}
