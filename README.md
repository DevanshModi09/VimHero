# VimHero

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)
![Curriculum](https://img.shields.io/badge/curriculum-14%2F45%20days-blueviolet)
![Engine](https://img.shields.io/badge/vim%20engine-hand--built-orange)

**A terminal game that teaches Vim from zero to hero, one day at a time.**

Every day unlocks a short lesson and five hands-on challenges, played on
a real hand-built modal editor — not a wrapper around your system's
`vim`.

```
╭────────────────────────────────────────────────────────────────╮
│ Day 14 — Final Boss (Everything From Days 1-13)                 │
│                                                                │
│ A whole messy document, twelve lines, one pass — every tool     │
│ from Days 1-13, back to back.                                  │
│                                                                │
│ 💡 Work top to bottom, one problem at a time — this one's      │
│ meant to take a few minutes, not a few seconds.                │
│                                                                │
│     XXXX this document has a messy header                     │
│                                                                │
│ NORMAL   keystrokes: 0   par: 100                              │
╰────────────────────────────────────────────────────────────────╯
```

## Contents

- [Play](#play)
- [How it teaches](#how-it-teaches)
- [Scoring](#scoring)
- [Curriculum roadmap](#curriculum-roadmap)
- [How it's built](#how-its-built)
- [Building a binary](#building-a-binary)
- [Tests](#tests)

## Play

Requires [Go 1.26+](https://go.dev/dl/).

```sh
git clone https://github.com/DevanshModi09/VimHero.git
cd VimHero
go run .
```

Progress — unlocked days, best keystroke counts, star ratings, and your
daily streak — is saved to `~/.vimhero/progress.json`.

### Controls

| Screen                    | Keys                                                                                                                |
| ------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| Day list / challenge list | `j`/`k` or arrows to move · `enter` to select · `esc` back · `q` quit                                               |
| Inside a challenge        | real Vim keys — each lesson tells you what's new · `esc` in Normal mode backs out · `ctrl+r` restarts the challenge |
| After solving a challenge | `enter` for the next one · `r` to retry for a better score · `esc` back to the day                                  |

### How it teaches

Every challenge explains what the new command actually _does_ — not just
what to press — before asking you to use it, and drops a 💡 tip on
screen with a related trick or a common gotcha (e.g. why `b` sometimes
"wastes" a keystroke snapping to the start of your current word instead
of jumping to the previous one).

Every few days, a **boss checkpoint** throws no new keys at you at all —
just a handful of tasks that only work if you combine everything you've
learned so far.

### Scoring

| Result                        | Stars |
| ----------------------------- | :---: |
| At or under par keystrokes    |  ★★★  |
| Under double par              |  ★★   |
| Anything else (still counts!) |   ★   |

Clearing every challenge in a day unlocks the next one.

## Curriculum roadmap

| Week | Days  | Theme                                             | Status                    |
| ---- | ----- | ------------------------------------------------- | ------------------------- |
| 1    | 1–5   | Basic movement & modes                            | ✅ done                   |
| 2    | 6–14  | More operators (`cw`, `ciw`, `daw`, `D`/`C`, `r`/`~`, `s`/`S`, `X`/`u`, `0`/`^`/`$`, ...) | ✅ done |
| 3    | 15–21 | Counts & text objects                             | ⏳ planned                |
| 4    | 22–28 | Find & search                                     | ⏳ planned                |
| 5    | 29–35 | Visual mode                                       | ⏳ planned                |
| 6    | 36–42 | Marks, macros, substitution, global               | ⏳ planned                |
| —    | 43–45 | Final boss challenges                             | ⏳ planned                |

Every authored day carries 5 hands-on challenges, each one verified by
scripting its solution against the real engine before being written
down — not hand-computed.

## How it's built

- **`pkg/editor`** — the engine. A real modal text editor
  supporting Vim's core grammar: motions (`h j k l w b e 0 ^ $ gg G f t
F T %`), operators (`d c y` composed with motions and text objects
  like `iw`, `i(`, `i"`), line ops (`dd yy p P`), replace/case (`r`,
  `~`), substitute (`s`, `S`), backward delete (`X`), undo (`u`), insert
  mode, visual/visual-line mode, marks, macros (`q`/`@`), search (`/`
  `n` `N`), and `:s` substitution. It's driven by a single
  `Buffer.Input(key string)` call per keystroke, which is what makes it
  easy to test and to embed in a TUI.
- **`pkg/curriculum`** — the 45-day lesson plan as data: each `Day`
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

`pkg/editor` has unit tests for the core engine. `internal/ui` has
a playthrough test that scripts the intended solution for _every_
authored challenge and asserts it actually wins — a guard against
authoring mistakes like wrong par counts, unreachable goals, or wrong
targets.
