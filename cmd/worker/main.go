package main

import (
	"log"
	"os"

	"github.com/talaniz/prime-mover-go/activities"
	"github.com/talaniz/prime-mover-go/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	options := client.Options{}
	if hostPort := os.Getenv("TEMPORAL_ADDRESS"); hostPort != "" {
		options.HostPort = hostPort
	}

	c, err := client.Dial(options)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	w := worker.New(c, workflows.TaskQueue, worker.Options{})
	w.RegisterWorkflow(workflows.FindLatestCodexReadyIssue)
	w.RegisterActivity(activities.Activities{})

	log.Printf("worker listening on task queue %q", workflows.TaskQueue)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatal(err)
	}
}
