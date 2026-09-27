package activities

import (
	"context"
	"os"

	"github.com/talaniz/prime-mover-go/internal/githubissues"
	"github.com/talaniz/prime-mover-go/internal/projects"
)

type LatestIssueRequest struct {
	ProjectID string
}

type LatestIssueResult struct {
	ProjectID  string
	Repository string
	Number     int
	Title      string
	URL        string
}

type Activities struct{}

func (Activities) FindLatestCodexReadyIssue(ctx context.Context, request LatestIssueRequest) (*LatestIssueResult, error) {
	project, err := projects.Resolve(request.ProjectID)
	if err != nil {
		return nil, err
	}

	finder, err := githubissues.NewFinder(os.Getenv("GITHUB_TOKEN"))
	if err != nil {
		return nil, err
	}

	issue, err := finder.LatestCodexReadyIssue(ctx, project.Owner, project.Repo)
	if err != nil {
		return nil, err
	}

	return &LatestIssueResult{
		ProjectID:  project.ID,
		Repository: project.Owner + "/" + project.Repo,
		Number:     issue.Number,
		Title:      issue.Title,
		URL:        issue.URL,
	}, nil
}
