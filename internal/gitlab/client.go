package gitlab

import (
	"fmt"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type Client struct {
	sdk       *gitlab.Client
	projectID string
}

func NewClient(token, projectID string) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("gitlab token cannot be empty")
	}
	if projectID == "" {
		return nil, fmt.Errorf("gitlab project ID cannot be empty")
	}

	git, err := gitlab.NewClient(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create gitlab client: %w", err)
	}

	return &Client{
		sdk:       git,
		projectID: projectID,
	}, nil
}
