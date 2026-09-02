package gitlab

import (
	"context"
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

type CreateIssueInput struct {
	Title       string
	Description string
	Labels      []string
}

func (c *Client) CreateIssue(ctx context.Context, input CreateIssueInput) (*gitlab.Issue, error) {
	opts := &gitlab.CreateIssueOptions{
		Title:       gitlab.Ptr(input.Title),
		Description: gitlab.Ptr(input.Description),
		Labels:      gitlab.Ptr(gitlab.LabelOptions(input.Labels)),
	}

	issue, _, err := c.sdk.Issues.CreateIssue(c.projectID, opts, gitlab.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to create a gitlab issue: %w", err)
	}

	return issue, nil
}
