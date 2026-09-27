package githubissues

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-github/v63/github"
)

type IssueSummary struct {
	Number int
	Title  string
	URL    string
}

type Finder struct {
	client *github.Client
}

func NewFinder(token string) (*Finder, error) {
	httpClient := http.DefaultClient
	if token != "" {
		httpClient = &http.Client{
			Transport: tokenTransport{
				token: token,
				base:  http.DefaultTransport,
			},
		}
	}
	return &Finder{client: github.NewClient(httpClient)}, nil
}

func NewFinderWithHTTPClient(httpClient *http.Client, baseURL string) (*Finder, error) {
	client := github.NewClient(httpClient)
	if baseURL != "" {
		parsed, err := url.Parse(baseURL)
		if err != nil {
			return nil, err
		}
		client.BaseURL = parsed
	}
	return &Finder{client: client}, nil
}

type tokenTransport struct {
	token string
	base  http.RoundTripper
}

func (t tokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.Header.Set("Authorization", "Bearer "+t.token)
	return t.base.RoundTrip(clone)
}

func (f *Finder) LatestCodexReadyIssue(ctx context.Context, owner, repo string) (*IssueSummary, error) {
	issues, _, err := f.client.Issues.ListByRepo(ctx, owner, repo, &github.IssueListByRepoOptions{
		State:     "open",
		Labels:    []string{"codex-ready"},
		Sort:      "created",
		Direction: "desc",
		ListOptions: github.ListOptions{
			PerPage: 1,
		},
	})
	if err != nil {
		return nil, err
	}
	if len(issues) == 0 {
		return nil, fmt.Errorf("no open codex-ready issues found in %s/%s", owner, repo)
	}

	issue := issues[0]
	return &IssueSummary{
		Number: issue.GetNumber(),
		Title:  issue.GetTitle(),
		URL:    issue.GetHTMLURL(),
	}, nil
}
