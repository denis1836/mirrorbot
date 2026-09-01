package sync

import (
	"context"

	"mirrorbot/internal/gitlab"
)

type IssueProcessor struct {
	gitlabClient *gitlab.Client
}

func NewIssueProcessor(gl *gitlab.Client) *IssueProcessor {
	return &IssueProcessor{gitlabClient: gl}
}

func (p *IssueProcessor) Process(ctx context.Context, rawPayload []byte) error {

	//TODO: json issue parsing
	//TODO: sending request to GitLab API

	return nil
}

var _ EventProcessor = (*IssueProcessor)(nil)
