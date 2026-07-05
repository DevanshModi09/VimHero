package editor

func (b *Buffer) handleInsert(key string) {
	switch key {
	case "esc":
		b.Mode = ModeNormal
		if b.Cursor.Col > 0 {
			b.Cursor.Col--
		}
		return
	case "enter":
		line := []rune(b.line(b.Cursor.Row))
		col := clampCol(b.Cursor.Col, len(line))
		before := string(line[:col])
		after := string(line[col:])
		newLines := append([]string(nil), b.Lines[:b.Cursor.Row]...)
		newLines = append(newLines, before, after)
		newLines = append(newLines, b.Lines[b.Cursor.Row+1:]...)
		b.Lines = newLines
		b.Cursor.Row++
		b.Cursor.Col = 0
		return
	case "backspace":
		if b.Cursor.Col > 0 {
			line := []rune(b.line(b.Cursor.Row))
			b.Lines[b.Cursor.Row] = string(deleteRange(line, b.Cursor.Col-1, b.Cursor.Col))
			b.Cursor.Col--
		} else if b.Cursor.Row > 0 {
			prevLen := b.lineLen(b.Cursor.Row - 1)
			b.Lines[b.Cursor.Row-1] += b.Lines[b.Cursor.Row]
			b.Lines = append(b.Lines[:b.Cursor.Row], b.Lines[b.Cursor.Row+1:]...)
			b.Cursor.Row--
			b.Cursor.Col = prevLen
		}
		return
	case "tab":
		key = "\t"
	}
	r := []rune(key)
	if len(r) != 1 {
		return
	}
	line := []rune(b.line(b.Cursor.Row))
	col := clampCol(b.Cursor.Col, len(line))
	b.Lines[b.Cursor.Row] = string(insertAt(line, col, key))
	b.Cursor.Col = col + 1
}
