package githubissues

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
)

func TestLatestCodexReadyIssueUsesGitHubQuery(t *testing.T) {
	transport := fakeTransport{
		t: t,
		body: `[{
			"number": 42,
			"title": "Tighten dashboard button spacing",
			"html_url": "https://github.com/talaniz/doom-control/issues/42"
		}]`,
	}
	finder, err := NewFinderWithHTTPClient(&http.Client{Transport: transport}, "https://api.github.test/")
	if err != nil {
		t.Fatal(err)
	}

	issue, err := finder.LatestCodexReadyIssue(context.Background(), "talaniz", "doom-control")
	if err != nil {
		t.Fatal(err)
	}
	if issue.Number != 42 || issue.Title != "Tighten dashboard button spacing" {
		t.Fatalf("unexpected issue: %#v", issue)
	}
}

func TestLatestCodexReadyIssueReportsEmptyResult(t *testing.T) {
	finder, err := NewFinderWithHTTPClient(
		&http.Client{Transport: fakeTransport{t: t, body: "[]"}},
		"https://api.github.test/",
	)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := finder.LatestCodexReadyIssue(context.Background(), "talaniz", "doom-control"); err == nil {
		t.Fatal("expected empty result error")
	}
}

type fakeTransport struct {
	t    *testing.T
	body string
}

func (f fakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	f.t.Helper()
	if req.URL.Host != "api.github.test" {
		f.t.Fatalf("unexpected host %s", req.URL.Host)
	}
	if req.URL.Path != "/repos/talaniz/doom-control/issues" {
		f.t.Fatalf("unexpected path %s", req.URL.Path)
	}
	query := req.URL.Query()
	if query.Get("state") != "open" {
		f.t.Fatalf("state query = %q", query.Get("state"))
	}
	if query.Get("labels") != "codex-ready" {
		f.t.Fatalf("labels query = %q", query.Get("labels"))
	}
	if query.Get("sort") != "created" || query.Get("direction") != "desc" {
		f.t.Fatalf("sort query = %q direction = %q", query.Get("sort"), query.Get("direction"))
	}
	if query.Get("per_page") != "1" {
		f.t.Fatalf("per_page query = %q", query.Get("per_page"))
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(bytes.NewBufferString(f.body)),
		Request:    req,
	}, nil
}
