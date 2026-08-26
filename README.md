# Shell Witness

> Your command line history testifies against your sloppy debugging.

A terminal tool reads your shell history and turns each command into a witness in a case against your debugging habits. It flags dangerous sudo, mystery rm, and repeated retries as hostile testimony. Export a verdict report that is equal parts performance review and comedy sketch.

## Features
- Parse local shell history into a timeline of commands, errors, and retries.
- Assign witness personalities to command families such as git, sudo, and docker.
- Score a chaos verdict based on risky patterns and repeated mistakes.
- Generate a Markdown trial transcript with quotes, damages, and recommended habits.

## Stack
- Go
- SQLite
- Cobra

## Getting started
```
go build -o shell-witness . && ./shell-witness report --history ~/.zsh_history --out verdict.md
```

---
*Farmed 🚜 by [Appshaker](https://github.com/buberlo) — shaken into existence.*
