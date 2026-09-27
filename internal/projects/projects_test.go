package projects

import "testing"

func TestResolveDoomDashboard(t *testing.T) {
	project, err := Resolve("doom-dashboard")
	if err != nil {
		t.Fatal(err)
	}
	if project.Owner != "talaniz" || project.Repo != "doom-control" {
		t.Fatalf("resolved wrong repository: %#v", project)
	}
}

func TestResolveUnknownProject(t *testing.T) {
	if _, err := Resolve("missing"); err == nil {
		t.Fatal("expected unknown project error")
	}
}
