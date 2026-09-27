package projects

import "fmt"

type Project struct {
	ID    string
	Owner string
	Repo  string
}

func Resolve(id string) (Project, error) {
	switch id {
	case "doom-dashboard":
		return Project{
			ID:    "doom-dashboard",
			Owner: "talaniz",
			Repo:  "doom-control",
		}, nil
	default:
		return Project{}, fmt.Errorf("unknown project %q", id)
	}
}
