package github

import (
	"encoding/json"
	"fmt"

	github "github.com/google/go-github/v62/github"
)

type IssueInput struct {
	Title       string
	Description string
	Labels      []string
}

func ParseIssuesEvent(rawPayload []byte) (*IssueInput, error) {
	var event github.IssuesEvent
	if err := json.Unmarshal(rawPayload, &event); err != nil {
		return nil, fmt.Errorf("failed to unmarshal github issues event: %w", err)
	}

	if event.GetAction() != "opened" {
		return nil, fmt.Errorf("unsupported action '%s': for now only 'opened' is handled", event.GetAction())
	}

	issue := event.GetIssue()
	if issue == nil {
		return nil, fmt.Errorf("event payload missing issue object")
	}

	var labels []string
	for _, label := range issue.Labels {
		if name := label.GetName(); name != "" {
			labels = append(labels, name)
		}
	}

	return &IssueInput{
		Title:       issue.GetTitle(),
		Description: issue.GetBody(),
		Labels:      labels,
	}, nil
}

type PullRequestInput struct {
	Title        string
	Description  string
	SourceBranch string
	TargetBranch string
	Labels       []string
}

func ParsePullRequestEvent(rawPayload []byte) (*PullRequestInput, error) {
	var event github.PullRequestEvent
	if err := json.Unmarshal(rawPayload, &event); err != nil {
		return nil, fmt.Errorf("failed to unmarshal github pull request event: %w", err)
	}

	if event.GetAction() != "opened" {
		return nil, fmt.Errorf("unsupported action %v, only 'opened' is supported", event.GetAction())
	}

	pr := event.GetPullRequest()
	if pr == nil {
		return nil, fmt.Errorf("event payload missing pull request object")
	}

	var labels []string
	for _, label := range pr.Labels {
		if name := label.GetName(); name != "" {
			labels = append(labels, name)
		}
	}

	return &PullRequestInput{
		Title:        pr.GetTitle(),
		Description:  pr.GetBody(),
		SourceBranch: pr.GetHead().GetRef(),
		TargetBranch: pr.GetBase().GetRef(),
		Labels:       labels,
	}, nil
}
