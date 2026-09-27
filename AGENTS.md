# Prime Mover Go repository instructions

Prime Mover Go is a Temporal-first prototype. Keep work focused on completing the
ticket-to-PR workflow in small, reviewable steps. Do not broaden the system into a
production platform before the current workflow slice needs it.

## Review workflow

- For every code-changing commit, push the commit remotely and request an independent
  code review for that exact commit.
- A code review may be skipped only when the change is documentation-only or otherwise
  does not touch active working code.
- The code reviewer should comment on the relevant ticket or review thread. Record the
  reviewed commit and result in the working conversation.
- Keep reviewer findings in scope. If a reviewer recommends an out-of-scope change,
  negotiate to cancel or defer that finding rather than silently expanding the work.
- After all commits for a change are in and code review findings are resolved or
  explicitly deferred, request a separate end-to-end review over the full change.
- Create or mark the PR ready only after the code reviewer, the main implementer, and
  the end-to-end reviewer all sign off.

## Scope discipline

- Prefer the smallest implementation that advances the current ticket-to-PR workflow.
- Keep Temporal workflow orchestration deterministic; put network, filesystem, GitHub,
  and other side effects in activities.
- Do not add durable storage, broad configuration systems, review automation, deployment
  machinery, or multi-project generalization until a ticket explicitly requires it.
- Favor readable Go code, small packages, and tests around observable workflow behavior.
