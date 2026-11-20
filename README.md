# Learning Go — templ

A companion course to [Learning-Go-Htmx](../Learning-Go-Htmx) that teaches
[templ](https://templ.guide) — a server-side HTML templating language for Go
where components are type-checked Go functions.

While `Learning-Go-Htmx` uses Go's standard `html/template`, this project uses
`templ` so you can compare the two approaches. The two topics are orthogonal:
you can use `templ` with or without HTMX.

## Lessons

| Lesson | What you learn |
|--------|----------------|
| 1 | Your first component — a Go function returning HTML |
| 2 | Props — component arguments |
| 3 | Conditionals — `if` / `else` |
| 4 | Loops — `for` over a slice |
| 5 | Composition — nesting components, passing components as children |
| 6 | Serving components over `net/http` |
| 7 | Forms — POST and re-render |
| 8 | templ + HTMX — returning HTML fragments |

## Running

```bash
go generate ./...   # regenerate the .templ.go files (requires the templ CLI)
go run .
```

Then open <http://localhost:8080>.

The generated `_templ.go` files are committed so the project builds without
installing the `templ` CLI.

## Why a separate project?

templ and HTMX are different concerns:

- **templ** = how you *generate* HTML on the server (components, props, loops, conditionals).
- **HTMX** = how you *update* the page on the client (swaps, triggers, targets).

Keeping them separate means each course stays focused and beginner-friendly.
Lesson 8 in this project shows how the two work together as a bridge.

## Project layout

```
templ/            templ source files (.templ) + generated .templ.go files
model.go          shared data types (Lesson)
main.go           HTTP handlers + course routes
static/style.css  shared stylesheet
go.mod            single dependency: github.com/a-h/templ
```

## Is this sufficient to be productive?

Yes for learning templ and building simple server-rendered Go UIs — all
8 lessons run (`go run .`) and the generated `_templ.go` files are
committed, so you don't need the `templ` CLI to build.

Not yet for shipping production code out of the box. Things you'd want
to add later when you've finished all the lessons:

- A `go:generate` directive or Makefile so `templ generate ./...` runs
  automatically before builds.
- Dev hot-reload (`templ watch` + `air`/`fresh`) so template edits
  rebuild without a manual server restart.
- Tests that render components and assert HTML output.
- 404/500 HTML error pages, logging middleware, and recovery.
- Production config (Dockerfile, env-based settings, structured logging).

Add those incrementally once you're comfortable writing templ components.
