package sync

import (
	"context"
	"fmt"

	"mirrorbot/internal/github"
	"mirrorbot/internal/gitlab"
)

type PullRequestProcessor struct {
	gitlabClient *gitlab.Client
}

func NewPullRequestProcessor(gl *gitlab.Client) *PullRequestProcessor {
	return &PullRequestProcessor{gitlabClient: gl}
}

func (p *PullRequestProcessor) Process(ctx context.Context, rawPayload []byte) error {
	prInput, err := github.ParsePullRequestEvent(rawPayload)
	if err != nil {
		return fmt.Errorf("failed to parse pr payload: %w", err)
	}

	mr, err := p.gitlabClient.CreateMergeRequest(ctx, gitlab.CreateMergeRequestInput{
		Title:        prInput.Title,
		Description:  prInput.Description,
		SourceBranch: prInput.SourceBranch,
		TargetBranch: prInput.TargetBranch,
		Labels:       prInput.Labels,
	})
	if err != nil {
		return fmt.Errorf("failed to create gitlab merge request: %w", err)
	}

	_ = mr
	return nil
}

var _ EventProcessor = (*PullRequestProcessor)(nil)
