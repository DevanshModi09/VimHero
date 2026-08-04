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

## Week 4: Find & Search — Days 22-28 ✅

| Day | Title | Challenges |
| --- | ----- | ---------- |
| 22 | f, F, t, T — Jump Straight To Any Character | 5 |
| 23 | / and ? — Search The Whole Buffer | 5 |
| 24 | % — Jump To The Matching Bracket | 5 |
| 25 | * and # — Search The Word Under The Cursor | 5 |
| 26 | d/, c/, y/ — Operators Meet Search | 5 |
| 27 | Search Patterns — \d, ^, and Character Classes | 5 |
| 28 | Boss Challenge — Everything From Days 22-27 | 5 |

Same-line character jumps (`f`/`F`/`t`/`T` plus `;`/`,`), whole-buffer
search (`/`, `?`, `n`, `N`), bracket matching (`%`), search-the-word-
under-cursor (`*`/`#`), then search used as an operator's motion
(`d/`, `c/`, `y/`). Day 27 covers `\d`, `^`, and `[set]` as three
goal-reach challenges, then ties them back into day 26's `d/`/`c/`
operators — first with `d/\d` alone, then a four-line finale combining
all three patterns with `d/` and `c/`. Day 28 closes out the week as a
checkpoint boss, following the pattern set by days 8, 11, 14, and 21 —
no new keys, just `f`/`;`, `/`/`n`, `%`, `*`, and pattern-powered `d/`/`c/`
recombined, including `d%` (an operator+bracket-motion pairing not
spelled out on day 24 itself, in keeping with how day 14 first paired
`d`/`c` with day 13's `0`/`^`/`$` without it being taught explicitly).

## Week 5: Visual Mode, Marks, Macros, Substitution & Global — Days 29-34 ✅

All remaining new teaching content is compressed into these 6 days, so
that every day from 35 onward is pure recap — no new keys introduced
past day 34.

| Day | Title | Status |
| --- | ----- | ------ |
| 29 | `v` and `V` — Visual & Visual-Line Mode | ✅ |
| 30 | `o` — Swap Ends Of A Visual Selection | ✅ |
| 31 | Text Objects & Case Toggle In Visual Mode (`vi(`, `vi"`, `viw`, `~`) | ✅ |
| 32 | `m` and `'` — Marks | ✅ |
| 33 | `q`, `@`, `@@` — Macros | ✅ |
| 34 | `:s` and `:g` — Substitution & The Global Command | ✅ |

Day 29 introduces `v`/`V` on their own — selection grown with plain
motions (`l`, `e`, `$`, `j`), then acted on with `d`, `y`, `c`, or `~` —
deliberately holding back `o` (day 30) and visual text objects (day 31)
so each stays a clean, single new idea.

The engine (`pkg/editor/visual.go`, `normal.go`, `command.go`) already
implements all of this — `v`/`V`/`o`, motions and text objects inside
visual mode, `d`/`y`/`c`/`~` over a selection, marks (`m`/`'`), macros
(`q`/`@`), `:s` substitution, and `:g` — so this week is buildable
against existing functionality with no engine changes needed.

## Days 35-45: Boss Gauntlet ✅

No new keys from here on — every day is a cumulative recap, increasing
in scope until day 45 pulls in the entire curriculum.

| Day | Recap scope | Status |
| --- | ----------- | ------ |
| 35 | Movement & modes (days 1-5) | ✅ |
| 36 | Operators (days 6-14) | ✅ |
| 37 | Counts & text objects (days 15-21) | ✅ |
| 38 | Find & search (days 22-28) | ✅ |
| 39 | Visual mode (days 29-31) | ✅ |
| 40 | Marks & macros (days 32-33) | ✅ |
| 41 | Substitution & global (day 34) | ✅ |
| 42 | Search + visual combined | ✅ |
| 43 | Operators + marks + macros combined | ✅ |
| 44 | Big mixed recap — nearly everything | ✅ |
| 45 | The Ultimate Boss — all 45 days | ✅ |

Day 35 follows the checkpoint-boss pattern of days 8, 11, 14, 21, and 28
— no new keys, five tasks recombining `hjkl`, `w`/`e`, `gg`/`G`,
`i`/`A`/`o`/`O`, `x`, and `dd`/`yy`/`p`/`P`, with the route described
rather than dictated key by key. Its first challenge is a `KindGoal`
reach, matching how week 1 itself opened on pure movement.

Day 36 recaps the operators week (days 6-14): `dw`/`cw`/`ciw`/`diw`,
`daw`, `D`/`C`, `r`/`~`, `s`/`S`, `X`/`u`, and `0`/`^`/`$`, ending in a
six-line boss that touches all of them. Day 37 recaps counts and text
objects (days 15-21): a count prefix stacked on any operator or motion
(`3dw`, `2dd`), the full `i`/`a` family across `(`, `"`, `[`, `{`, and
`'`, yanking with a text object (`yiw`, `yi(`) instead of a motion, and
`Y` as the fast whole-line yank — its boss combines a counted
word-delete, a counted line-delete, one fix per bracket/quote
delimiter, and a final `Y`/`p` duplicate. Day 38 recaps find & search
(days 22-28): `f`/`F`/`t`/`T` chained with `;`/`,`, whole-buffer `/`/`?`
with `n`/`N` (including overshooting and correcting with `N`), `%`
across a nested bracket pair, `*`/`#` to rename by word instead of
retyping a search, and `d/`/`c/` paired with the `\d`/`^`/`[set]`
patterns from day 27 — its eight-line boss touches all six tools in one
pass. Day 39 recaps the visual-mode trio from days 29-31: `v`/`V` to
start and grow a selection, `o` to swap which end is movable, and the
visual text objects (`viw`, `vi(`, `vi"`) that aim a selection the same
way `iw`/`i(`/`i"` aim an operator — its six-line boss chains a visual
yank-and-paste, a `V`/`c` whole-line rewrite, `vi"`/`c`, `viw`/`d`, and
an `o`-widened delete into one pass. Day 40 recaps the other Week 5 pair:
`m`/`'` to bookmark a spot and jump straight back to it, and `q`/`@`/`@@`
to record a keystroke sequence once and replay it as many times as there
are lines left — its six-line boss bookmarks a paste destination, folds
a fix-and-advance edit into a macro, replays it across two more lines,
then yanks a value and jumps back to the bookmark to paste it before a
final cleanup delete. Day 41 recaps Day 34 alone: `:s/old/new/` on the
current line, the `g` flag for every match on that line, a `%` prefix
for every line in the buffer, and `:g/pattern/d` to delete every
matching line — its boss reruns Day 34's own capstone shape, clearing
junk lines with `:g` before rewriting what's left with `%s`/`g`. Day 42
is the first cross-week recap: search motions (`f`/`F`/`t`/`T`, `/`,
`%`, `*`) from Days 22-27 paired with visual mode (`v`, `V`, text
objects) from Days 29-31, built on the fact that `f`/`F`/`t`/`T` and `%`
work as ordinary growing motions inside an active `v`/`V` selection —
its boss chains a `*`-then-`viw`-`~` fix, an `f(`-then-`v%d` bracket
delete, a `/pattern`-then-`Vd` line delete, and a closing `o` insert.
All seven days are `KindEdit` challenges (no `KindGoal` reach), matching
the shape of days 8, 11, 14, 21, and 28 from the weeks they recap.

Day 43 is the first recap built specifically around Days 32-34's marks and
macros paired with the operators week (Days 6-14): `m`/`'` to bookmark a
spot and jump back after editing elsewhere, and a macro built from `r`,
`daw`, or `X`/`s` to replay one operator edit down a whole list — its boss
chains a mark, a macro built from `daw`, and a `D`/`C` fix reached by
jumping back to the bookmark. Day 44 widens the net further, pulling from
nearly every week at once: a counted `2dw` paired with a `/` search and
`daw`, `:g` cleanup paired with a visual case toggle, a macro recording an
entire `:s///g` command, and marks paired with `viw`/`~` twice — its boss
strings all four together across a six-line buffer. Day 45 closes out the
curriculum: each of its first four tasks spotlights one broad slice (Week
1 movement and Insert mode, Week 2-3 operators/text-objects/counts, Week
4-5 search/visual/marks, and Day 32-34 macros/substitution/global), and
the fifth is the grand finale — a mark, a `:g` cleanup, a macro built from
`ciw` and `di(`, a search-driven visual case toggle, a fresh line opened
with `o`, and a final jump back to the bookmark to retitle it, all in one
pass across the longest buffer in the game.

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
