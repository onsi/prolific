package main

import (
	"fmt"
	"regexp"

	"strings"
)

var CSV_HEADERS = []string{
	"Title",
	"Type",
	"Description",
	"Labels",
}

type Story struct {
	Type        string
	Title       string
	Description string
	Labels      []string
}

func (s *Story) AppendLine(line string) error {
	var err error

	if s.Title == "" {
		if line == "" {
			return nil
		}
		s.Title, s.Type, err = s.ParseTitleLine(line)
		return err
	}

	if strings.HasPrefix(line, "L: ") {
		labels := strings.Split(strings.TrimPrefix(line, "L: "), ",")
		for i := range labels {
			labels[i] = strings.TrimSpace(labels[i])
		}
		s.Labels = labels
		return nil
	}

	s.Description += line + "\n"
	return nil
}

func (s *Story) ParseTitleLine(line string) (string, string, error) {
	if line[0] == '[' {
		storyType, err := s.ParseStoryType(line)
		if err != nil {
			return "", "", err
		}
		re := regexp.MustCompile(`\]`)
		indices := re.FindStringIndex(line)
		title := strings.TrimSpace(line[indices[1]:])
		return title, storyType, nil
	} else {
		return line, "feature", nil
	}
}

func (s *Story) ParseStoryType(line string) (string, error) {
	re := regexp.MustCompile(`\[(.*)\]`)
	segments := re.FindStringSubmatch(line)
	if len(segments) != 2 {
		return "", fmt.Errorf("Invalid story type for:\n\t%s\n", line)
	}
	storyType := strings.ToLower(segments[1])
	for _, validStoryType := range []string{"feature", "chore", "bug", "release"} {
		if strings.HasPrefix(validStoryType, storyType) {
			return validStoryType, nil
		}
	}
	return "", fmt.Errorf("Invalid story type:\n\t%s\nfor:\n\t%s\n", segments[1], line)
}

func (s Story) CSVRecords() []string {
	return []string{
		s.Title,
		s.Type,
		strings.TrimSpace(s.Description),
		strings.Join(s.Labels, ","),
	}
}
