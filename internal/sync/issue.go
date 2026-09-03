package sync

import (
	"context"
	"fmt"

	"mirrorbot/internal/github"
	"mirrorbot/internal/gitlab"
)

type IssueProcessor struct {
	gitlabClient *gitlab.Client
}

func NewIssueProcessor(gl *gitlab.Client) *IssueProcessor {
	return &IssueProcessor{gitlabClient: gl}
}

func (p *IssueProcessor) Process(ctx context.Context, rawPayload []byte) error {

	ghIssue, err := github.ParseIssuesEvent(rawPayload)
	if err != nil {
		return fmt.Errorf("issue processor failed to parse payload: %w", err)
	}

	createdIssue, err := p.gitlabClient.CreateIssue(ctx, gitlab.CreateIssueInput{
		Title:       ghIssue.Title,
		Description: ghIssue.Description,
		Labels:      ghIssue.Labels,
	})
	if err != nil {
		return fmt.Errorf("issue processor failed to create a gitlab issue: %w", err)
	}

	_ = createdIssue
	return nil
}

var _ EventProcessor = (*IssueProcessor)(nil)
