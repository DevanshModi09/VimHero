package editor

import "strconv"

func (b *Buffer) countVal() int {
	if b.count == "" {
		return 1
	}
	n, err := strconv.Atoi(b.count)
	if err != nil || n < 1 {
		return 1
	}
	return n
}

func (b *Buffer) totalCount() int {
	op := b.pendingOpCount
	if op < 1 {
		op = 1
	}
	return op * b.countVal()
}

func (b *Buffer) resetPending() {
	b.count = ""
	b.pendingOp = 0
	b.pendingOpCount = 0
	b.pendingG = false
	b.pendingKind = 0
	b.awaitAround = false
}

func isDigit(key string) bool {
	return len(key) == 1 && key[0] >= '0' && key[0] <= '9'
}

func (b *Buffer) handleNormal(key string) {
	if b.pendingKind != 0 {
		b.handleAwaitedArg(key)
		return
	}

	if b.pendingG {
		b.pendingG = false
		switch key {
		case "g":
			row := b.countVal() - 1
			if b.count == "" {
				row = 0
			}
			b.gotoMotion(Pos{row, b.lineFirstNonBlank(row)}, true)
		default:
			b.resetPending()
		}
		return
	}

	if isDigit(key) && !(key == "0" && b.count == "") {
		b.count += key
		return
	}

	if key == "q" && b.recordingReg != 0 {
		b.stopRecording()
		b.resetPending()
		return
	}

	switch key {
	case "esc":
		b.resetPending()
		return
	}

	if b.pendingOp == 0 {
		switch key {
		case "d", "c", "y":
			b.pendingOp = key[0]
			b.pendingOpCount = b.countVal()
			b.count = ""
			return
		}
	} else {
		if key == string(b.pendingOp) {
			n := b.totalCount()
			op := b.pendingOp
			b.applyLinewise(op, 0, b.Cursor.Row, b.Cursor.Row+n-1)
			b.resetPending()
			return
		}
		switch key {
		case "i", "a":
			b.pendingKind = 'i'
			b.awaitAround = key == "a"
			return
		}
	}

	switch key {
	case "h", "left":
		b.gotoMotion(b.moveLeft(b.totalCount()), false)
	case "l", "right":
		b.gotoMotion(b.moveRight(b.totalCount(), b.pendingOp != 0), false)
	case "j", "down":
		n := b.totalCount()
		if b.pendingOp != 0 {
			b.applyLinewise(b.pendingOp, 0, b.Cursor.Row, b.Cursor.Row+n)
			b.resetPending()
			return
		}
		b.Cursor = b.moveDown(n)
		b.resetPending()
	case "k", "up":
		n := b.totalCount()
		if b.pendingOp != 0 {
			b.applyLinewise(b.pendingOp, 0, b.Cursor.Row, b.Cursor.Row-n)
			b.resetPending()
			return
		}
		b.Cursor = b.moveUp(n)
		b.resetPending()
	case "w":
		if b.pendingOp == 'c' && wordClass(runeAt(b.line(b.Cursor.Row), b.Cursor.Col)) != 0 {
			b.gotoMotion(b.wordEnd(b.totalCount(), false), true)
		} else {
			b.gotoMotion(b.wordForward(b.totalCount(), false), false)
		}
	case "W":
		if b.pendingOp == 'c' && wordClass(runeAt(b.line(b.Cursor.Row), b.Cursor.Col)) != 0 {
			b.gotoMotion(b.wordEnd(b.totalCount(), true), true)
		} else {
			b.gotoMotion(b.wordForward(b.totalCount(), true), false)
		}
	case "b":
		b.gotoMotion(b.wordBackward(b.totalCount(), false), false)
	case "B":
		b.gotoMotion(b.wordBackward(b.totalCount(), true), false)
	case "e":
		b.gotoMotion(b.wordEnd(b.totalCount(), false), true)
	case "E":
		b.gotoMotion(b.wordEnd(b.totalCount(), true), true)
	case "0":
		b.gotoMotion(Pos{b.Cursor.Row, 0}, false)
	case "^":
		b.gotoMotion(Pos{b.Cursor.Row, b.lineFirstNonBlank(b.Cursor.Row)}, false)
	case "$":
		n := b.totalCount()
		row := b.Cursor.Row + n - 1
		if row >= len(b.Lines) {
			row = len(b.Lines) - 1
		}
		b.gotoMotion(Pos{row, b.normalMaxCol(row)}, true)
	case "G":
		row := len(b.Lines) - 1
		if b.count != "" {
			row = b.countVal() - 1
		}
		b.gotoMotion(Pos{row, b.lineFirstNonBlank(row)}, true)
	case "g":
		b.pendingG = true
		return
	case "%":
		if p, ok := b.matchPercent(); ok {
			b.gotoMotion(p, true)
		} else {
			b.resetPending()
		}
	case "f", "t", "F", "T":
		b.pendingKind = key[0]
		return
	case ";":
		b.repeatFind(1)
	case ",":
		b.repeatFind(-1)
	}

	if b.pendingOp != 0 {
		switch key {
		case "h", "left", "l", "right", "w", "W", "b", "B", "e", "E", "0", "^", "$", "G", "%", "f", "t", "F", "T", ";", ",":
		default:
			b.resetPending()
		}
		return
	}

	switch key {
	case "x":
		n := b.totalCount()
		line := []rune(b.line(b.Cursor.Row))
		to := clampCol(b.Cursor.Col+n, len(line))
		if to > b.Cursor.Col {
			b.pushUndo()
			b.setRegister(0, string(line[b.Cursor.Col:to]), false)
			b.Lines[b.Cursor.Row] = string(deleteRange(line, b.Cursor.Col, to))
		}
		b.resetPending()
	case "X":
		n := b.totalCount()
		line := []rune(b.line(b.Cursor.Row))
		from := clampCol(b.Cursor.Col-n, len(line))
		if from < b.Cursor.Col {
			b.pushUndo()
			b.setRegister(0, string(line[from:b.Cursor.Col]), false)
			b.Lines[b.Cursor.Row] = string(deleteRange(line, from, b.Cursor.Col))
			b.Cursor.Col = from
		}
		b.resetPending()
	case "D":
		b.applyCharwise('d', 0, b.Cursor, Pos{b.Cursor.Row, b.lineLen(b.Cursor.Row)}, false)
		b.resetPending()
	case "C":
		b.applyCharwise('c', 0, b.Cursor, Pos{b.Cursor.Row, b.lineLen(b.Cursor.Row)}, false)
		b.resetPending()
	case "s":
		n := b.totalCount()
		line := []rune(b.line(b.Cursor.Row))
		to := clampCol(b.Cursor.Col+n, len(line))
		b.applyCharwise('c', 0, b.Cursor, Pos{b.Cursor.Row, to}, false)
		b.resetPending()
	case "S":
		n := b.totalCount()
		b.applyLinewise('c', 0, b.Cursor.Row, b.Cursor.Row+n-1)
		b.resetPending()
	case "Y":
		n := b.totalCount()
		b.applyLinewise('y', 0, b.Cursor.Row, b.Cursor.Row+n-1)
		b.resetPending()
	case "p":
		b.doPaste(true, 0)
		b.resetPending()
	case "P":
		b.doPaste(false, 0)
		b.resetPending()
	case "u":
		b.undo()
		b.resetPending()
	case "i":
		b.pushUndo()
		b.Mode = ModeInsert
		b.resetPending()
	case "a":
		b.pushUndo()
		if b.lineLen(b.Cursor.Row) > 0 {
			b.Cursor.Col++
		}
		b.Mode = ModeInsert
		b.resetPending()
	case "I":
		b.pushUndo()
		b.Cursor.Col = b.lineFirstNonBlank(b.Cursor.Row)
		b.Mode = ModeInsert
		b.resetPending()
	case "A":
		b.pushUndo()
		b.Cursor.Col = b.lineLen(b.Cursor.Row)
		b.Mode = ModeInsert
		b.resetPending()
	case "o":
		b.pushUndo()
		row := b.Cursor.Row
		newLines := append([]string(nil), b.Lines[:row+1]...)
		newLines = append(newLines, "")
		newLines = append(newLines, b.Lines[row+1:]...)
		b.Lines = newLines
		b.Cursor = Pos{row + 1, 0}
		b.Mode = ModeInsert
		b.resetPending()
	case "O":
		b.pushUndo()
		row := b.Cursor.Row
		newLines := append([]string(nil), b.Lines[:row]...)
		newLines = append(newLines, "")
		newLines = append(newLines, b.Lines[row:]...)
		b.Lines = newLines
		b.Cursor = Pos{row, 0}
		b.Mode = ModeInsert
		b.resetPending()
	case "v":
		b.Mode = ModeVisual
		b.visualStart = b.Cursor
		b.resetPending()
	case "V":
		b.Mode = ModeVisualLine
		b.visualStart = b.Cursor
		b.resetPending()
	case "r":
		b.pendingKind = 'r'
	case "m":
		b.pendingKind = 'm'
	case "'", "`":
		b.pendingKind = '\''
	case "~":
		line := []rune(b.line(b.Cursor.Row))
		if b.Cursor.Col < len(line) {
			b.pushUndo()
			r := line[b.Cursor.Col]
			line[b.Cursor.Col] = toggleCase(r)
			b.Lines[b.Cursor.Row] = string(line)
			if b.Cursor.Col < b.normalMaxCol(b.Cursor.Row)+1 {
				b.Cursor.Col++
			}
		}
		b.resetPending()
	case ":":
		b.Mode = ModeCommand
		b.commandKind = ':'
		b.CommandLine = ""
		b.resetPending()
	case "/":
		b.Mode = ModeCommand
		b.commandKind = '/'
		b.CommandLine = ""
		b.searchDir = 1
		b.resetPending()
	case "?":
		b.Mode = ModeCommand
		b.commandKind = '/'
		b.CommandLine = ""
		b.searchDir = -1
		b.resetPending()
	case "n":
		b.repeatSearch(1)
		b.resetPending()
	case "N":
		b.repeatSearch(-1)
		b.resetPending()
	case "*":
		b.searchWord(1)
		b.resetPending()
	case "#":
		b.searchWord(-1)
		b.resetPending()
	case "q":
		b.pendingKind = 'q'
	case "@":
		b.pendingKind = '@'
	default:
		b.resetPending()
	}
}

func toggleCase(r rune) rune {
	switch {
	case r >= 'a' && r <= 'z':
		return r - 32
	case r >= 'A' && r <= 'Z':
		return r + 32
	}
	return r
}

func (b *Buffer) gotoMotion(target Pos, inclusive bool) {
	if b.pendingOp == 0 {
		b.Cursor = target
		b.desiredCol = target.Col
		b.resetPending()
		return
	}
	op := b.pendingOp
	from := b.Cursor
	b.applyCharwise(op, 0, from, target, inclusive)
	b.resetPending()
}

func (b *Buffer) repeatFind(dir int) {
	if b.lastFind.char == 0 {
		return
	}
	kind := b.lastFind.kind
	if dir < 0 {
		switch kind {
		case 'f':
			kind = 'F'
		case 'F':
			kind = 'f'
		case 't':
			kind = 'T'
		case 'T':
			kind = 't'
		}
	}
	if p, ok := b.findChar(kind, b.lastFind.char, b.totalCount()); ok {
		inclusive := kind == 'f' || kind == 't'
		b.gotoMotion(p, inclusive)
	} else {
		b.resetPending()
	}
}

func (b *Buffer) handleAwaitedArg(key string) {
	kind := b.pendingKind
	r := []rune(key)
	var ch rune
	if len(r) > 0 {
		ch = r[0]
	}

	switch kind {
	case 'f', 't', 'F', 'T':
		if key == "esc" {
			b.resetPending()
			return
		}
		b.lastFind = pendingFind{kind: kind, char: ch}
		if p, ok := b.findChar(kind, ch, b.totalCount()); ok {
			inclusive := kind == 'f' || kind == 't'
			b.gotoMotion(p, inclusive)
		} else {
			b.resetPending()
		}
	case 'r':
		if key == "esc" {
			b.resetPending()
			return
		}
		n := b.totalCount()
		line := []rune(b.line(b.Cursor.Row))
		if b.Cursor.Col+n <= len(line) {
			b.pushUndo()
			for i := 0; i < n; i++ {
				line[b.Cursor.Col+i] = ch
			}
			b.Lines[b.Cursor.Row] = string(line)
			b.Cursor.Col += n - 1
		}
		b.resetPending()
	case 'm':
		b.marks[ch] = b.Cursor
		b.resetPending()
	case '\'':
		if p, ok := b.marks[ch]; ok {
			b.gotoMotion(p, false)
		} else {
			b.resetPending()
		}
	case 'i':
		around := b.awaitAround
		op := b.pendingOp
		from, to, ok := b.textObjectRange(around, ch)
		if !ok {
			b.resetPending()
			return
		}
		row := b.Cursor.Row
		if op != 0 {
			b.applyCharwise(op, 0, Pos{row, from}, Pos{row, to}, false)
		}
		b.resetPending()
	case 'q':
		b.recordingReg = ch
		b.macroBuf.Reset()
		b.resetPending()
	case '@':
		b.resetPending()
		if ch == '@' {
			ch = b.lastMacroReg
		}
		b.playMacro(ch)
	}
}

func (b *Buffer) stopRecording() {
	if b.recordingReg == 0 {
		return
	}
	b.macros[b.recordingReg] = b.macroBuf.String()
	b.recordingReg = 0
}

func (b *Buffer) playMacro(reg rune) {
	macro, ok := b.macros[reg]
	if !ok || b.replayDepth > 50 {
		return
	}
	b.lastMacroReg = reg
	b.replayDepth++
	for _, k := range splitKeys(macro) {
		b.Input(k)
	}
	b.replayDepth--
}

func splitKeys(s string) []string {
	if s == "" {
		return nil
	}
	out := []string{}
	cur := ""
	for _, r := range s {
		if r == 0 {
			out = append(out, cur)
			cur = ""
			continue
		}
		cur += string(r)
	}
	return out
}

func joinKeys(keys []string) string {
	out := ""
	for _, k := range keys {
		out += k + "\x00"
	}
	return out
}
