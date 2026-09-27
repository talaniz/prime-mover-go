package workflows

import (
	"time"

	"github.com/talaniz/prime-mover-go/activities"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "prime-mover"

func FindLatestCodexReadyIssue(ctx workflow.Context, projectID string) (*activities.LatestIssueResult, error) {
	logger := workflow.GetLogger(ctx)
	options := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var result activities.LatestIssueResult
	err := workflow.ExecuteActivity(
		ctx,
		activities.Activities.FindLatestCodexReadyIssue,
		activities.LatestIssueRequest{ProjectID: projectID},
	).Get(ctx, &result)
	if err != nil {
		return nil, err
	}

	logger.Info(
		"latest codex-ready issue",
		"project", result.ProjectID,
		"repository", result.Repository,
		"number", result.Number,
		"title", result.Title,
		"url", result.URL,
	)

	return &result, nil
}
