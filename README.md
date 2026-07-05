# VimHero

A terminal game that teaches Vim from zero to hero over 45 days. Every day
unlocks a short lesson and a couple of hands-on challenges, played on a
real (hand-built) modal editor — not a wrapper around your system's `vim`.

## Play

```
go run .
```

Progress (unlocked days, best keystroke counts, star ratings, daily streak)
is saved to `~/.vimhero/progress.json`.

### Controls

- Day list / challenge list: `j`/`k` or arrows to move, `enter` to select, `esc` to go back, `q` to quit.
- In a challenge: real Vim keys. Each day's lesson tells you which ones are new.
- `esc` while in Normal mode backs out to the challenge list.

### Learning aids

Every challenge explains what the new command actually does — not just
what to press — before asking you to use it, and shows a 💡 tip on
screen with a related trick or a common gotcha (e.g. why `b` sometimes
"wastes" a keystroke snapping to the start of your current word).

### Scoring

Each challenge has a par keystroke count. Clear it at or under par for
★★★, under double par for ★★, otherwise ★. Clearing every challenge in a
day unlocks the next one.

## How it's built

- `internal/editor` — the engine. A real modal text editor supporting
  Vim's core grammar: motions (`h j k l w b e 0 ^ $ gg G f t F T %`),
  operators (`d c y` composed with motions and text objects like `iw`,
  `i(`, `i"`), line ops (`dd yy p P`), insert mode, visual/visual-line
  mode, undo, marks, macros (`q`/`@`), search (`/` `n` `N`), and `:s`
  substitution. It's driven by a simple `Buffer.Input(key string)` call
  per keystroke, which is what makes it easy to test and to embed in a
  TUI.
- `internal/curriculum` — the 45-day lesson plan as data: each `Day` has
  a summary and one or more `Challenge`s, either "reach this cursor
  position" or "transform the buffer into this target text."
- `internal/progress` — persists unlock state, best scores, and streaks.
- `internal/ui` — the Bubble Tea TUI tying it all together.

## Building a binary

```
go build -o vimhero .
```

To cross-compile for another OS (e.g. from macOS/Linux for Windows),
set `GOOS`/`GOARCH` — no extra toolchain needed:

```
GOOS=windows GOARCH=amd64 go build -o vimhero.exe .
```

## Status

Days 1-5 (Week 1: basic movement & modes) are built. Days 6-45 are in
progress.

## Tests

```
go test ./...
```

`internal/editor` has unit tests for the core engine. `internal/ui` has
a playthrough test that scripts the intended solution for every
authored challenge and asserts it actually wins — a guard against
authoring mistakes (wrong par counts, unreachable goals, wrong targets).
