# Tetris

This is a Go-based Tetris project for macOS. The current focus is a playable core gameplay loop, clear module boundaries, and a codebase that can be extended and verified incrementally.

## What is included today

The project currently provides these core capabilities:

- A playable single-player Tetris game window
- Basic controls: move left, move right, rotate, and soft drop
- Game state transitions: pause, resume, restart, and game over
- Core rules: piece spawning, collision checks, locking, line clears, scoring, and level progression
- On-screen information: board, active piece, next piece preview, score, level, and cleared lines

These capabilities are based on the current codebase and OpenSpec documents. The goal is to build a solid and maintainable gameplay foundation before expanding into non-core features.

## Technical background

- Language: Go
- Graphics and input: Ebiten
- Target platform: macOS
- Project style: keep game rules, input handling, and rendering separated so the gameplay logic remains understandable and testable

## Download

- GitHub Releases: <https://github.com/force416/tetris/releases/latest>
- Each tagged release publishes a macOS `.dmg` asset that can be downloaded directly from the release page

## Run and verify

The project provides `Makefile` entry points:

```bash
make run
```

Starts the windowed Tetris application.

```bash
make test
```

Runs `go test ./...` to verify the current automated tests.

```bash
make build
```

Builds the executable to `dist/build/Tetris`.

```bash
make app
```

Packages a macOS `.app` bundle at `dist/Tetris.app`.

```bash
make dmg
```

Packages a universal macOS `.dmg` file at `dist/Tetris-darwin-universal.dmg`.

If you prefer to use Go commands directly, the main entry point is `cmd/tetris/main.go`.

## Controls

- `←` / `A`: move left
- `→` / `D`: move right
- `↑` / `W` / `X` / `Space`: rotate
- `↓` / `S`: soft drop
- `P`: pause or resume
- `R`: restart

## Project structure

- `cmd/tetris`
  Application entry point. Creates the game configuration and starts the app.
- `internal/app`
  Application coordinator. Owns the main loop, time accumulation, input-action to game-command translation, and gravity tick progression.
- `internal/game`
  Core rules and state model. Owns the board, piece spawning, collision checks, rotation, locking, line clears, scoring, leveling, and overall game state.
- `internal/input`
  Keyboard polling layer. Translates key presses into gameplay actions.
- `internal/render`
  Rendering layer. Draws the current game state as the board, sidebar, and pause or game-over overlays.
- `openspec`
  Requirements, design, and task-tracking documents for proposed and archived changes.

## Current limitations

- The project is currently focused on the single-player core gameplay loop and does not include leaderboards, multiplayer, accounts, or network features.
- The README only describes capabilities that are implemented today or supported by existing specs; it does not promise unfinished features.
- This is still an early-stage project, and both the documentation and feature set will continue evolving through OpenSpec-managed changes.

## Specs and development context

- The core gameplay specification is in `openspec/specs/playable-core-loop/spec.md`
- Project changes are managed through OpenSpec, moving from proposal to design to specs before implementation
