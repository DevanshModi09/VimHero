package editor

import "strings"

// applyLinewise deletes/yanks/changes whole lines [row1,row2] (inclusive).
func (b *Buffer) applyLinewise(op byte, reg rune, row1, row2 int) {
	if row1 > row2 {
		row1, row2 = row2, row1
	}
	if row1 < 0 {
		row1 = 0
	}
	if row2 >= len(b.Lines) {
		row2 = len(b.Lines) - 1
	}
	text := strings.Join(b.Lines[row1:row2+1], "\n")

	switch op {
	case 'y':
		b.setRegister(reg, text, true)
		b.Cursor.Row = row1
		b.Cursor.Col = b.lineFirstNonBlank(row1)
		return
	case 'd', 'c':
		b.pushUndo()
		b.setRegister(reg, text, true)
		rest := append([]string(nil), b.Lines[:row1]...)
		rest = append(rest, b.Lines[row2+1:]...)
		if len(rest) == 0 {
			rest = []string{""}
		}
		b.Lines = rest
		if op == 'd' {
			b.Cursor.Row = row1
			if b.Cursor.Row >= len(b.Lines) {
				b.Cursor.Row = len(b.Lines) - 1
			}
			b.Cursor.Col = b.lineFirstNonBlank(b.Cursor.Row)
		} else {
			// 'c'c: leave a blank line to type into.
			newLines := append([]string(nil), b.Lines[:row1]...)
			newLines = append(newLines, "")
			newLines = append(newLines, b.Lines[row1:]...)
			b.Lines = newLines
			b.Cursor.Row = row1
			b.Cursor.Col = 0
			b.Mode = ModeInsert
		}
	}
}

// applyCharwise deletes/yanks/changes the span from..to on possibly
// multiple lines. `inclusive` means the character at `to` is included.
func (b *Buffer) applyCharwise(op byte, reg rune, from, to Pos, inclusive bool) {
	if to.Row < from.Row || (to.Row == from.Row && to.Col < from.Col) {
		from, to = to, from
	}
	if inclusive {
		to.Col++
	}
	if from.Row == to.Row {
		line := []rune(b.line(from.Row))
		fc := clampCol(from.Col, len(line))
		tc := clampCol(to.Col, len(line))
		if fc >= tc {
			return
		}
		text := string(line[fc:tc])
		switch op {
		case 'y':
			b.setRegister(reg, text, false)
			b.Cursor = Pos{from.Row, fc}
			return
		case 'd', 'c':
			b.pushUndo()
			b.setRegister(reg, text, false)
			newLine := deleteRange(line, fc, tc)
			b.Lines[from.Row] = string(newLine)
			b.Cursor = Pos{from.Row, fc}
			if op == 'c' {
				b.Mode = ModeInsert
			}
		}
		return
	}

	firstLine := []rune(b.line(from.Row))
	lastLine := []rune(b.line(to.Row))
	fc := clampCol(from.Col, len(firstLine))
	tc := clampCol(to.Col, len(lastLine))

	var textParts []string
	textParts = append(textParts, string(firstLine[fc:]))
	for r := from.Row + 1; r < to.Row; r++ {
		textParts = append(textParts, b.line(r))
	}
	textParts = append(textParts, string(lastLine[:tc]))
	text := strings.Join(textParts, "\n")

	switch op {
	case 'y':
		b.setRegister(reg, text, false)
		b.Cursor = Pos{from.Row, fc}
		return
	case 'd', 'c':
		b.pushUndo()
		b.setRegister(reg, text, false)
		merged := string(firstLine[:fc]) + string(lastLine[tc:])
		newLines := append([]string(nil), b.Lines[:from.Row]...)
		newLines = append(newLines, merged)
		newLines = append(newLines, b.Lines[to.Row+1:]...)
		b.Lines = newLines
		b.Cursor = Pos{from.Row, fc}
		if op == 'c' {
			b.Mode = ModeInsert
		}
	}
}

func (b *Buffer) doPaste(after bool, reg rune) {
	c := b.getRegister(reg)
	if c.text == "" {
		return
	}
	b.pushUndo()
	if c.linewise {
		parts := strings.Split(c.text, "\n")
		row := b.Cursor.Row
		insertAt := row
		if after {
			insertAt = row + 1
		}
		newLines := append([]string(nil), b.Lines[:insertAt]...)
		newLines = append(newLines, parts...)
		newLines = append(newLines, b.Lines[insertAt:]...)
		b.Lines = newLines
		b.Cursor.Row = insertAt
		b.Cursor.Col = b.lineFirstNonBlank(insertAt)
		return
	}
	line := []rune(b.line(b.Cursor.Row))
	col := b.Cursor.Col
	if after && len(line) > 0 {
		col++
	}
	if strings.Contains(c.text, "\n") {
		parts := strings.Split(c.text, "\n")
		before := string(line[:clampCol(col, len(line))])
		afterStr := string(line[clampCol(col, len(line)):])
		newLines := append([]string(nil), b.Lines[:b.Cursor.Row]...)
		newLines = append(newLines, before+parts[0])
		newLines = append(newLines, parts[1:len(parts)-1]...)
		newLines = append(newLines, parts[len(parts)-1]+afterStr)
		newLines = append(newLines, b.Lines[b.Cursor.Row+1:]...)
		b.Lines = newLines
		b.Cursor.Row += len(parts) - 1
		b.Cursor.Col = len([]rune(parts[len(parts)-1]))
		return
	}
	newLine := insertAt(line, col, c.text)
	b.Lines[b.Cursor.Row] = string(newLine)
	b.Cursor.Col = col + len([]rune(c.text)) - 1
}
