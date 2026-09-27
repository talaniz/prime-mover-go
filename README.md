# Prime Mover Go

Temporal-first Go prototype for Prime Mover workflows.

This repository is a clean-room spike for rebuilding Prime Mover around Temporal
workflows. The first target is intentionally small: run one workflow for the
`doom-dashboard` project, query `talaniz/doom-control`, and print the latest open
GitHub issue labeled `codex-ready`.

## Development VM

The included `Vagrantfile` starts a Debian 12 VM with Docker, Go, Temporal Server,
Postgres, and Temporal UI.

```sh
vagrant up
```

Default endpoints:

- Temporal frontend: `192.168.56.50:7233`
- Temporal UI: `http://192.168.56.50:8233`

Useful commands:

```sh
vagrant ssh
temporal-compose ps
temporal-compose logs -f temporal
docker run --rm --network host temporalio/admin-tools:1.25.2 temporal operator cluster health
```

You can override the VM and Temporal versions:

```sh
PRIME_MOVER_VM_IP=192.168.56.51 vagrant up
TEMPORAL_VERSION=1.25.2 TEMPORAL_UI_VERSION=2.31.2 vagrant provision
```

## First Workflow

Current first slice:

1. Start a local Temporal worker.
2. Run a workflow for `doom-dashboard`.
3. Resolve that project to `talaniz/doom-control`.
4. Fetch the latest open `codex-ready` GitHub issue.
5. Log and return the issue number, title, and URL.

GitHub access should use an explicit token such as `GITHUB_TOKEN`; workflow code
should stay deterministic, with network calls isolated in activities.

Run tests:

```sh
go test ./...
```

With the Temporal VM running, start a worker in one shell:

```sh
export TEMPORAL_ADDRESS=127.0.0.1:7233
export GITHUB_TOKEN=...
go run ./cmd/worker
```

Then start the workflow in another shell:

```sh
export TEMPORAL_ADDRESS=127.0.0.1:7233
go run ./cmd/run-workflow
```

The runner starts `FindLatestCodexReadyIssue` on the `prime-mover` task queue and
prints the newest open `codex-ready` issue in `talaniz/doom-control`.
