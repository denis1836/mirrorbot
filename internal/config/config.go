package config

import (
	"fmt"
	"os"

	"github.com/sethvargo/go-githubactions"
)

type Config struct {
	GitLabToken     string
	GitLabProjectID string
	EventName       string
	EventPath       string
}

func Load(action *githubactions.Action) (*Config, error) {
	token := action.GetInput("gitlab_token")
	if token == "" {
		token = os.Getenv("GITLAB_TOKEN")
	}

	projectID := action.GetInput("gitlab_project_id")
	if projectID == "" {
		projectID = os.Getenv("GITLAB_PROJECT_ID")
	}

	eventName := os.Getenv("GITHUB_EVENT_NAME")
	eventPath := os.Getenv("GITHUB_EVENT_PATH")

	if token == "" {
		return nil, fmt.Errorf("missing required input: gitlab_token")
	}
	if projectID == "" {
		return nil, fmt.Errorf("missing required input: gitlab_project_id")
	}
	if eventPath == "" {
		return nil, fmt.Errorf("missing environment variable: GITHUB_EVENT_PATH")
	}

	return &Config{
		GitLabToken:     token,
		GitLabProjectID: projectID,
		EventName:       eventName,
		EventPath:       eventPath,
	}, nil
}
