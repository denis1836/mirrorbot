package sync

import (
	"context"
	"fmt"

	"mirrorbot/internal/gitlab"
)

type Engine struct {
	gitlabClient *gitlab.Client
}

func NewEngine(gl *gitlab.Client) *Engine {
	return &Engine{gitlabClient: gl}
}

func (e *Engine) Dispatch(ctx context.Context, eventName string, payload []byte) error {
	var processor EventProcessor

	switch eventName {
	case "issues":
		processor = NewIssueProcessor(e.gitlabClient)
	default:
		return fmt.Errorf("unsupported event type: %v", eventName)
	}

	return processor.Process(ctx, payload)
}
