// Package editor implements a small but real modal text editor with
// Vim-compatible keybindings. It is the engine that every VimHero
// challenge runs on: navigation puzzles and text-transform puzzles are
// both just "reach a state" checks against this Buffer.
package editor

import "strings"

type Mode int

const (
	ModeNormal Mode = iota
	ModeInsert
	ModeVisual
	ModeVisualLine
	ModeCommand // ':' or '/' prompt at the bottom
)

type Pos struct {
	Row, Col int
}

type pendingFind struct {
	kind byte // 'f','t','F','T'
	char rune
}

// Buffer is the whole editor state machine. One Buffer == one challenge
// attempt. Callers feed it raw key tokens via Input and read back Lines/
// Cursor/Mode to render and to check win conditions.
type Buffer struct {
	Lines  []string
	Cursor Pos
	Mode   Mode

	desiredCol int // "sticky" column for j/k over short lines

	count       string // accumulating digit-prefix, e.g. "3" before "3w"
	pendingOp   byte   // 'd','c','y' waiting on a motion
	pendingG    bool   // saw 'g', waiting on second key (gg)
	pendingKind byte   // 0, or one of 'f','t','F','T','r','m','\'','q','@','i','a' awaiting an argument char
	awaitAround bool   // for pendingKind 'i'/'a': true means "a" (around), false means "i" (inner)

	pendingOpCount int // count captured before the operator key, multiplied with the motion's count

	visualStart Pos

	registers     map[rune]regContent
	marks         map[rune]Pos
	lastFind      pendingFind
	searchPattern string
	searchDir     int // +1 forward, -1 backward
	recordingReg  rune
	macroBuf      strings.Builder
	macros        map[rune]string
	replayDepth   int
	lastMacroReg  rune

	CommandLine string // text typed after ':' or '/'
	commandKind byte   // ':' or '/'
	StatusMsg   string

	undoStack []snapshot

	Keystrokes int
	KeyLog     []string
}

type snapshot struct {
	lines  []string
	cursor Pos
}

type regContent struct {
	text     string
	linewise bool
}

func New(lines []string, cursor Pos) *Buffer {
	if len(lines) == 0 {
		lines = []string{""}
	}
	b := &Buffer{
		Lines:      append([]string(nil), lines...),
		Cursor:     cursor,
		Mode:       ModeNormal,
		registers:  map[rune]regContent{},
		marks:      map[rune]Pos{},
		macros:     map[rune]string{},
		desiredCol: cursor.Col,
	}
	return b
}

func (b *Buffer) Text() string {
	return strings.Join(b.Lines, "\n")
}

func (b *Buffer) line(row int) string {
	if row < 0 || row >= len(b.Lines) {
		return ""
	}
	return b.Lines[row]
}

func (b *Buffer) lineLen(row int) int {
	return len([]rune(b.line(row)))
}

// clampCol clamps col into [0, len] where len allows one-past-the-end
// only in insert/visual contexts; normal mode clamps to len-1 (or 0).
func clampCol(col, max int) int {
	if max < 0 {
		max = 0
	}
	if col > max {
		col = max
	}
	if col < 0 {
		col = 0
	}
	return col
}

func (b *Buffer) normalMaxCol(row int) int {
	n := b.lineLen(row)
	if n == 0 {
		return 0
	}
	return n - 1
}

func (b *Buffer) clampCursor() {
	if b.Cursor.Row < 0 {
		b.Cursor.Row = 0
	}
	if b.Cursor.Row >= len(b.Lines) {
		b.Cursor.Row = len(b.Lines) - 1
	}
	max := b.normalMaxCol(b.Cursor.Row)
	if b.Mode == ModeInsert {
		max = b.lineLen(b.Cursor.Row)
	}
	b.Cursor.Col = clampCol(b.Cursor.Col, max)
}

func runeAt(s string, col int) rune {
	r := []rune(s)
	if col < 0 || col >= len(r) {
		return 0
	}
	return r[col]
}

func wordClass(r rune) int {
	switch {
	case r == 0 || r == ' ' || r == '\t':
		return 0
	case r == '_' || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'):
		return 1
	default:
		return 2
	}
}

func (b *Buffer) pushUndo() {
	linesCopy := append([]string(nil), b.Lines...)
	b.undoStack = append(b.undoStack, snapshot{lines: linesCopy, cursor: b.Cursor})
	if len(b.undoStack) > 200 {
		b.undoStack = b.undoStack[len(b.undoStack)-200:]
	}
}

func (b *Buffer) undo() {
	if len(b.undoStack) == 0 {
		b.StatusMsg = "already at oldest change"
		return
	}
	snap := b.undoStack[len(b.undoStack)-1]
	b.undoStack = b.undoStack[:len(b.undoStack)-1]
	b.Lines = snap.lines
	b.Cursor = snap.cursor
	b.clampCursor()
}

func (b *Buffer) setRegister(reg rune, text string, linewise bool) {
	if reg == 0 {
		reg = '"'
	}
	c := regContent{text: text, linewise: linewise}
	b.registers[reg] = c
	b.registers['"'] = c
}

func (b *Buffer) getRegister(reg rune) regContent {
	if reg == 0 {
		reg = '"'
	}
	return b.registers[reg]
}

func insertAt(s []rune, col int, add string) []rune {
	if col > len(s) {
		col = len(s)
	}
	out := make([]rune, 0, len(s)+len([]rune(add)))
	out = append(out, s[:col]...)
	out = append(out, []rune(add)...)
	out = append(out, s[col:]...)
	return out
}

func deleteRange(s []rune, from, to int) []rune {
	if from < 0 {
		from = 0
	}
	if to > len(s) {
		to = len(s)
	}
	if from >= to {
		return s
	}
	out := make([]rune, 0, len(s)-(to-from))
	out = append(out, s[:from]...)
	out = append(out, s[to:]...)
	return out
}
