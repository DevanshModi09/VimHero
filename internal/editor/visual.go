package editor

func (b *Buffer) handleVisual(key string) {
	if b.pendingKind != 0 {
		b.handleVisualAwaitedArg(key)
		return
	}
	if b.pendingG {
		b.pendingG = false
		if key == "g" {
			b.Cursor = Pos{0, 0}
		}
		return
	}
	if isDigit(key) && !(key == "0" && b.count == "") {
		b.count += key
		return
	}

	switch key {
	case "esc":
		b.Mode = ModeNormal
		b.resetPending()
		return
	case "v":
		if b.Mode == ModeVisual {
			b.Mode = ModeNormal
		} else {
			b.Mode = ModeVisual
		}
		b.resetPending()
		return
	case "V":
		if b.Mode == ModeVisualLine {
			b.Mode = ModeNormal
		} else {
			b.Mode = ModeVisualLine
		}
		b.resetPending()
		return
	case "o":
		b.Cursor, b.visualStart = b.visualStart, b.Cursor
		return
	}

	n := b.countVal()
	switch key {
	case "h", "left":
		b.Cursor = b.moveLeft(n)
		b.count = ""
		return
	case "l", "right":
		b.Cursor = b.moveRight(n, true)
		b.count = ""
		return
	case "j", "down":
		b.Cursor = b.moveDown(n)
		b.count = ""
		return
	case "k", "up":
		b.Cursor = b.moveUp(n)
		b.count = ""
		return
	case "w":
		b.Cursor = b.wordForward(n, false)
		b.count = ""
		return
	case "W":
		b.Cursor = b.wordForward(n, true)
		b.count = ""
		return
	case "b":
		b.Cursor = b.wordBackward(n, false)
		b.count = ""
		return
	case "B":
		b.Cursor = b.wordBackward(n, true)
		b.count = ""
		return
	case "e":
		b.Cursor = b.wordEnd(n, false)
		b.count = ""
		return
	case "E":
		b.Cursor = b.wordEnd(n, true)
		b.count = ""
		return
	case "0":
		b.Cursor.Col = 0
		return
	case "^":
		b.Cursor.Col = b.lineFirstNonBlank(b.Cursor.Row)
		return
	case "$":
		b.Cursor.Col = b.normalMaxCol(b.Cursor.Row)
		b.count = ""
		return
	case "g":
		b.pendingG = true
		return
	case "G":
		row := len(b.Lines) - 1
		if b.count != "" {
			row = b.countVal() - 1
		}
		b.Cursor = Pos{row, b.lineFirstNonBlank(row)}
		b.count = ""
		return
	case "%":
		if p, ok := b.matchPercent(); ok {
			b.Cursor = p
		}
		return
	case "f", "t", "F", "T":
		b.pendingKind = key[0]
		return
	case "i", "a":
		b.pendingKind = 'i'
		b.awaitAround = key == "a"
		return
	}

	switch key {
	case "d", "x":
		b.applyVisualOp('d')
	case "y":
		b.applyVisualOp('y')
	case "c":
		b.applyVisualOp('c')
	case "~":
		b.toggleCaseRange()
	}
}

func (b *Buffer) handleVisualAwaitedArg(key string) {
	kind := b.pendingKind
	ch := rune(0)
	if r := []rune(key); len(r) > 0 {
		ch = r[0]
	}
	switch kind {
	case 'f', 't', 'F', 'T':
		b.pendingKind = 0
		if key == "esc" {
			return
		}
		b.lastFind = pendingFind{kind: kind, char: ch}
		if p, ok := b.findChar(kind, ch, b.countVal()); ok {
			b.Cursor = p
		}
		b.count = ""
	case 'i':
		around := b.awaitAround
		b.pendingKind = 0
		from, to, ok := b.textObjectRange(around, ch)
		if !ok {
			return
		}
		row := b.Cursor.Row
		b.visualStart = Pos{row, from}
		b.Cursor = Pos{row, to - 1}
		if b.Cursor.Col < 0 {
			b.Cursor.Col = 0
		}
	}
}

func (b *Buffer) visualRange() (from, to Pos) {
	from, to = b.visualStart, b.Cursor
	if to.Row < from.Row || (to.Row == from.Row && to.Col < from.Col) {
		from, to = to, from
	}
	return
}

func (b *Buffer) applyVisualOp(op byte) {
	from, to := b.visualRange()
	linewise := b.Mode == ModeVisualLine
	b.Mode = ModeNormal
	if linewise {
		b.applyLinewise(op, 0, from.Row, to.Row)
	} else {
		b.applyCharwise(op, 0, from, to, true)
	}
	b.resetPending()
}

func (b *Buffer) toggleCaseRange() {
	from, to := b.visualRange()
	linewise := b.Mode == ModeVisualLine
	b.Mode = ModeNormal
	b.pushUndo()
	if linewise {
		for r := from.Row; r <= to.Row && r < len(b.Lines); r++ {
			line := []rune(b.Lines[r])
			for i := range line {
				line[i] = toggleCase(line[i])
			}
			b.Lines[r] = string(line)
		}
	} else if from.Row == to.Row {
		line := []rune(b.line(from.Row))
		end := clampCol(to.Col+1, len(line))
		for i := from.Col; i < end; i++ {
			line[i] = toggleCase(line[i])
		}
		b.Lines[from.Row] = string(line)
	}
	b.Cursor = from
	b.resetPending()
}
