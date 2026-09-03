package main

import (
	"context"
	"os"

	"mirrorbot/internal/config"
	"mirrorbot/internal/gitlab"
	"mirrorbot/internal/sync"

	"github.com/sethvargo/go-githubactions"
)

var c config.Config

func main() {
	action := githubactions.New()

	c, err := config.Load(action)
	if err != nil {
		action.Fatalf("failed to load config: %v", err)
	}

	glClient, err := gitlab.NewClient(c.GitLabToken, c.GitLabProjectID)
	if err != nil {
		action.Fatalf("failed to create GitLab client: %v", err)
	}

	payload, err := os.ReadFile(c.EventPath)
	if err != nil {
		action.Fatalf("failed to read event file from %s: %v", c.EventPath, err)
	}

	ctx := context.Background()
	engine := sync.NewEngine(glClient)

	if err := engine.Dispatch(ctx, c.EventName, payload); err != nil {
		action.Fatalf("event processing failed: %v", err)
	}

	action.Infof("Successfully processed event %q and mirrored to GitLab!", c.EventName)
}
