package main

import (
	"context"
	"log"
	"os"

	"github.com/talaniz/prime-mover-go/activities"
	"github.com/talaniz/prime-mover-go/workflows"
	"go.temporal.io/sdk/client"
)

func main() {
	clientOptions := client.Options{}
	if hostPort := os.Getenv("TEMPORAL_ADDRESS"); hostPort != "" {
		clientOptions.HostPort = hostPort
	}

	c, err := client.Dial(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "find-latest-codex-ready-doom-dashboard",
		TaskQueue: workflows.TaskQueue,
	}
	run, err := c.ExecuteWorkflow(
		context.Background(),
		options,
		workflows.FindLatestCodexReadyIssue,
		"doom-dashboard",
	)
	if err != nil {
		log.Fatal(err)
	}

	var result activities.LatestIssueResult
	if err := run.Get(context.Background(), &result); err != nil {
		log.Fatal(err)
	}

	log.Printf(
		"latest codex-ready issue: %s#%d %q %s",
		result.Repository,
		result.Number,
		result.Title,
		result.URL,
	)
}
