# The 45-Day Vim Curriculum

VimHero teaches Vim as a sequence of 45 days, grouped into 6 weeks plus a
closing run of final-boss days. Each day has a `Week`, a `Title`, a one
-paragraph `Summary` of the idea it teaches, and a handful of hands-on
`Challenge`s — either "reach this cursor position" (`KindGoal`) or
"transform the buffer into this target text" (`KindEdit`). Clearing every
challenge in a day unlocks the next one. The data lives in
[`pkg/curriculum/days.go`](pkg/curriculum/days.go); this file is the
human-readable map of it.

Legend: ✅ authored and complete · 🚧 authored but partial · ⏳ not yet
written (theme only, from the roadmap).

## Week 1: Basic Movement & Modes — Days 1-5 ✅

| Day | Title | Challenges |
| --- | ----- | ---------- |
| 1 | hjkl — Move Without Arrow Keys | 5 |
| 2 | Word Motions — w, e, b | 5 |
| 3 | Insert Mode — i, A, o, O | 5 |
| 4 | gg, G, and x | 5 |
| 5 | dd, yy, p, P — Delete, Yank, Paste | 5 |

Gets your fingers off the arrow keys and onto hjkl, teaches word-wise
movement, the four ways into insert mode, jumping to the top/bottom of a
buffer, and the core delete/yank/paste triad.

## Week 2: More Operators — Days 6-14 ✅

| Day | Title | Challenges |
| --- | ----- | ---------- |
| 6 | dw, cw, ciw, diw — Operators Meet Word Motions | 5 |
| 7 | daw, D, C — Around Word & End-of-Line Shortcuts | 5 |
| 8 | Boss Challenge — Everything From Days 1-7 | 5 |
| 9 | r and ~ — Replace One Character & Toggle Case | 5 |
| 10 | s and S — Substitute Character & Substitute Line | 5 |
| 11 | Boss | 5 |
| 12 | X and u — Delete Backward & Undo | 5 |
| 13 | 0, ^, and $ — Line Start, First Non-Blank, and Line End | 5 |
| 14 | Final Boss — Everything From Days 1-13 | 5 |

Introduces operator + motion composition (`d`/`c` paired with word
motions and text objects), single-character replace/case toggling,
substitute commands, backward delete, undo, and the three line-position
motions — with checkpoint "boss" days folded in every few lessons to
force recall instead of rote repetition.

## Week 3: Counts & Text Objects — Days 15-21 ✅

| Day | Title | Challenges |
| --- | ----- | ---------- |
| 15 | Counts — Repeat Any Motion or Operator in One Shot | 5 |
| 16 | i( and i" — Text Objects For Brackets & Quotes | 5 |
| 17 | i[ — Text Objects For Square Brackets | 1 🚧 |
| 18 | i{ and i' — Finishing The Text-Object Family | 5 |
| 19 | yiw, yi(, yi" — Yank Text Objects | 5 |
| 20 | Y — Yank Whole Line, The Fast Way | 5 |
| 21 | Full Revision — The Ultimate Boss (Everything From Days 1-20) | 5 |

Numeric counts (`3dw`, `5j`, ...) stacked on top of every operator and
motion from weeks 1-2, then the full `i`/`a` text-object family
(`(` `"` `[` `{` `'`) for surgical edits inside delimiters, capped by a
full-syllabus revision boss. Day 17 has just 1 of its planned challenges
— the rest of that lesson is still open.

## Week 4: Find & Search — Days 22-28 🚧

| Day | Title | Challenges |
| --- | ----- | ---------- |
| 22 | f, F, t, T — Jump Straight To Any Character | 5 |
| 23 | / and ? — Search The Whole Buffer | 5 |
| 24 | % — Jump To The Matching Bracket | 5 |
| 25 | * and # — Search The Word Under The Cursor | 5 |
| 26 | d/, c/, y/ — Operators Meet Search | 5 |
| 27 | Search Patterns — \d, ^, and Character Classes | 3 🚧 |
| 28 | *(week 4 boss — not yet written)* | ⏳ |

Same-line character jumps (`f`/`F`/`t`/`T` plus `;`/`,`), whole-buffer
search (`/`, `?`, `n`, `N`), bracket matching (`%`), search-the-word-
under-cursor (`*`/`#`), then search used as an operator's motion
(`d/`, `c/`, `y/`). Day 27 currently covers `\d`, `^`, and `[set]` as
three standalone goal-reach challenges; still missing an operator+regex
challenge and a multi-fix finale (see below). Day 28 is expected to be
this week's boss day, following the pattern set by days 8, 11, 14, and
21, but hasn't been authored yet.

## Week 5: Visual Mode, Marks, Macros, Substitution & Global — Days 29-34 ⏳

All remaining new teaching content is compressed into these 6 days, so
that every day from 35 onward is pure recap — no new keys introduced
past day 34.

| Day | Title | Status |
| --- | ----- | ------ |
| 29 | `v` and `V` — Visual & Visual-Line Mode | ⏳ |
| 30 | Extending A Visual Selection (motions + `o` to swap ends) | ⏳ |
| 31 | Text Objects & Case Toggle In Visual Mode (`vi(`, `vi"`, `viw`, `~`) | ⏳ |
| 32 | `m` and `'` — Marks | ⏳ |
| 33 | `q`, `@`, `@@` — Macros | ⏳ |
| 34 | `:s` and `:g` — Substitution & The Global Command | ⏳ |

The engine (`pkg/editor/visual.go`, `normal.go`, `command.go`) already
implements all of this — `v`/`V`/`o`, motions and text objects inside
visual mode, `d`/`y`/`c`/`~` over a selection, marks (`m`/`'`), macros
(`q`/`@`), `:s` substitution, and `:g` — so this week is buildable
against existing functionality with no engine changes needed.

## Days 35-45: Boss Gauntlet ⏳

No new keys from here on — every day is a cumulative recap, increasing
in scope until day 45 pulls in the entire curriculum.

| Day | Recap scope |
| --- | ----------- |
| 35 | Movement & modes (days 1-5) |
| 36 | Operators (days 6-14) |
| 37 | Counts & text objects (days 15-21) |
| 38 | Find & search (days 22-28) |
| 39 | Visual mode (days 29-31) |
| 40 | Marks & macros (days 32-33) |
| 41 | Substitution & global (day 34) |
| 42 | Search + visual combined |
| 43 | Operators + marks + macros combined |
| 44 | Big mixed recap — nearly everything |
| 45 | The Ultimate Boss — all 45 days |

## Where the engine outpaces the curriculum

`pkg/editor` already implements marks, macros, `:s`, and `:g` — features
slated for days 32-34 — as well as full regex search (Go's `regexp`, not
just literal substrings) that day 27 only partially exercises. When
authoring new days, check `pkg/editor/*.go` first: the capability is
often already there and just needs a lesson built around it.

## How days get written

Per the project's own convention (see recent commit history), each
challenge's keystroke sequence and expected outcome — cursor position or
resulting buffer text — is verified against the real engine (`Buffer.Input`
fed one key at a time, as in `pkg/editor/editor_test.go`) before its `Par`
value is written into `days.go`. Nothing here is hand-computed.
