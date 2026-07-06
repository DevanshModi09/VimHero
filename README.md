# VimHero

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)
![Curriculum](https://img.shields.io/badge/curriculum-8%2F45%20days-blueviolet)
![Engine](https://img.shields.io/badge/vim%20engine-hand--built-orange)

**A terminal game that teaches Vim from zero to hero, one day at a time.**

Every day unlocks a short lesson and a couple of hands-on challenges,
played on a real hand-built modal editor — not a wrapper around your
system's `vim`.

```
╭─ Day 6 — dw, cw, ciw ───────────────────────────────────────╮
│                                                              │
│  cw changes a word, but stops at the end of it instead of   │
│  also eating the space after — unlike dw. Fix "quikc dog"   │
│  into "quick dog" using cw.                                 │
│                                                              │
│  💡 this is why cw feels natural for fixing typos.          │
│                                                              │
│  ┌────────────────────────────────────────────────────────┐│
│  │ quikc dog                                               ││
│  └────────────────────────────────────────────────────────┘│
│                                                              │
│  NORMAL   keystrokes: 0   par: 8                            │
╰──────────────────────────────────────────────────────────────╯
```

## Play

```sh
go run .
```

Progress — unlocked days, best keystroke counts, star ratings, and your
daily streak — is saved to `~/.vimhero/progress.json`.

### Controls

| Screen                    | Keys                                                    |
| -------------------------- | -------------------------------------------------------- |
| Day list / challenge list | `j`/`k` or arrows to move · `enter` to select · `esc` back · `q` quit |
| Inside a challenge        | real Vim keys — each lesson tells you what's new         |
|                            | `esc` in Normal mode backs out to the challenge list      |

### How it teaches

Every challenge explains what the new command actually *does* — not just
what to press — before asking you to use it, and drops a 💡 tip on
screen with a related trick or a common gotcha (e.g. why `b` sometimes
"wastes" a keystroke snapping to the start of your current word instead
of jumping to the previous one).

Every few days, a **boss checkpoint** throws no new keys at you at all —
just a handful of tasks that only work if you combine everything you've
learned so far.

### Scoring

| Result                         | Stars |
| ------------------------------- | :---: |
| At or under par keystrokes      | ★★★   |
| Under double par                | ★★    |
| Anything else (still counts!)   | ★     |

Clearing every challenge in a day unlocks the next one.

## Curriculum roadmap

| Week | Days  | Theme                                             | Status                    |
| ---- | ----- | -------------------------------------------------- | -------------------------- |
| 1    | 1–5   | Basic movement & modes                             | ✅ done                    |
| 2    | 6–14  | More operators (`cw`, `ciw`, `daw`, `D`/`C`, ...)  | 🚧 in progress (6–8 done)  |
| 3    | 15–21 | Counts & text objects                              | ⏳ planned                 |
| 4    | 22–28 | Find & search                                      | ⏳ planned                 |
| 5    | 29–35 | Visual mode                                        | ⏳ planned                 |
| 6    | 36–42 | Marks, macros, substitution, global                | ⏳ planned                 |
| —    | 43–45 | Final boss challenges                              | ⏳ planned                 |

## How it's built

- **`internal/editor`** — the engine. A real modal text editor
  supporting Vim's core grammar: motions (`h j k l w b e 0 ^ $ gg G f t
  F T %`), operators (`d c y` composed with motions and text objects
  like `iw`, `i(`, `i"`), line ops (`dd yy p P`), insert mode,
  visual/visual-line mode, undo, marks, macros (`q`/`@`), search (`/`
  `n` `N`), and `:s` substitution. It's driven by a single
  `Buffer.Input(key string)` call per keystroke, which is what makes it
  easy to test and to embed in a TUI.
- **`internal/curriculum`** — the 45-day lesson plan as data: each `Day`
  has a summary and one or more `Challenge`s, either "reach this cursor
  position" or "transform the buffer into this target text."
- **`internal/progress`** — persists unlock state, best scores, and
  streaks.
- **`internal/ui`** — the Bubble Tea TUI tying it all together.

## Building a binary

```sh
go build -o vimhero .
```

To cross-compile for another OS (e.g. from macOS/Linux for Windows), set
`GOOS`/`GOARCH` — no extra toolchain needed:

```sh
GOOS=windows GOARCH=amd64 go build -o vimhero.exe .
```

## Tests

```sh
go test ./...
```

`internal/editor` has unit tests for the core engine. `internal/ui` has
a playthrough test that scripts the intended solution for *every*
authored challenge and asserts it actually wins — a guard against
authoring mistakes like wrong par counts, unreachable goals, or wrong
targets.
