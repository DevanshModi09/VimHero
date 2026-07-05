package editor

// motionResult describes where a motion lands and how an operator should
// treat the span between the start position and here.
type motionResult struct {
	pos       Pos
	inclusive bool // true for e, f, t, %, $ (character is included)
	linewise  bool // true for gg, G, j, k, dd/yy-style
	valid     bool
}

func (b *Buffer) moveLeft(count int) Pos {
	p := b.Cursor
	p.Col -= count
	if p.Col < 0 {
		p.Col = 0
	}
	return p
}

func (b *Buffer) moveRight(count int, allowEnd bool) Pos {
	p := b.Cursor
	max := b.normalMaxCol(p.Row)
	if allowEnd {
		max = b.lineLen(p.Row)
	}
	p.Col += count
	if p.Col > max {
		p.Col = max
	}
	return p
}

func (b *Buffer) moveDown(count int) Pos {
	p := b.Cursor
	p.Row += count
	if p.Row >= len(b.Lines) {
		p.Row = len(b.Lines) - 1
	}
	p.Col = clampCol(b.desiredCol, b.normalMaxCol(p.Row))
	return p
}

func (b *Buffer) moveUp(count int) Pos {
	p := b.Cursor
	p.Row -= count
	if p.Row < 0 {
		p.Row = 0
	}
	p.Col = clampCol(b.desiredCol, b.normalMaxCol(p.Row))
	return p
}

func (b *Buffer) lineFirstNonBlank(row int) int {
	line := []rune(b.line(row))
	for i, r := range line {
		if r != ' ' && r != '\t' {
			return i
		}
	}
	return 0
}

// wordForward implements 'w'/'W'. bigWord=true means WORD (whitespace-
// delimited only, ignoring punctuation classes).
func (b *Buffer) wordForward(count int, bigWord bool) Pos {
	row, col := b.Cursor.Row, b.Cursor.Col
	for n := 0; n < count; n++ {
		line := []rune(b.line(row))
		if col >= len(line) {
			if row < len(b.Lines)-1 {
				row++
				col = 0
				line = []rune(b.line(row))
				for col < len(line) && wordClass(line[col]) == 0 {
					col++
				}
				continue
			}
			break
		}
		startClass := wordClass(runeAt(b.line(row), col))
		if bigWord && startClass != 0 {
			startClass = 1
		}
		for col < len(line) {
			cls := wordClass(line[col])
			if bigWord && cls != 0 {
				cls = 1
			}
			if cls != startClass {
				break
			}
			col++
		}
		for {
			line = []rune(b.line(row))
			for col < len(line) {
				cls := wordClass(line[col])
				if cls != 0 {
					break
				}
				col++
			}
			if col < len(line) || row >= len(b.Lines)-1 {
				break
			}
			row++
			col = 0
			if len(b.line(row)) == 0 {
				break
			}
		}
	}
	return Pos{row, col}
}

func (b *Buffer) wordBackward(count int, bigWord bool) Pos {
	row, col := b.Cursor.Row, b.Cursor.Col
	for n := 0; n < count; n++ {
		for {
			if col == 0 {
				if row == 0 {
					return Pos{0, 0}
				}
				row--
				col = b.lineLen(row)
				break
			}
			col--
			break
		}
		line := []rune(b.line(row))
		for col > 0 && col <= len(line) && (col == len(line) || wordClass(line[col]) == 0) {
			if col == 0 {
				break
			}
			col--
		}
		if len(line) == 0 {
			continue
		}
		if col >= len(line) {
			continue
		}
		cls := wordClass(line[col])
		if bigWord && cls != 0 {
			cls = 1
		}
		for col > 0 {
			pcls := wordClass(line[col-1])
			if bigWord && pcls != 0 {
				pcls = 1
			}
			if pcls != cls {
				break
			}
			col--
		}
	}
	return Pos{row, col}
}

func (b *Buffer) wordEnd(count int, bigWord bool) Pos {
	row, col := b.Cursor.Row, b.Cursor.Col
	for n := 0; n < count; n++ {
		col++
		for {
			line := []rune(b.line(row))
			if col >= len(line) {
				if row >= len(b.Lines)-1 {
					col = len(line) - 1
					if col < 0 {
						col = 0
					}
					goto done
				}
				row++
				col = 0
				continue
			}
			if wordClass(line[col]) != 0 {
				break
			}
			col++
		}
		{
			line := []rune(b.line(row))
			cls := wordClass(line[col])
			if bigWord && cls != 0 {
				cls = 1
			}
			for col+1 < len(line) {
				ncls := wordClass(line[col+1])
				if bigWord && ncls != 0 {
					ncls = 1
				}
				if ncls != cls {
					break
				}
				col++
			}
		}
	}
done:
	return Pos{row, col}
}

// findChar implements f/t/F/T. Returns ok=false if char not found.
func (b *Buffer) findChar(kind byte, ch rune, count int) (Pos, bool) {
	line := []rune(b.line(b.Cursor.Row))
	col := b.Cursor.Col
	found := col
	remaining := count
	switch kind {
	case 'f':
		c := col
		for remaining > 0 {
			c++
			for c < len(line) && line[c] != ch {
				c++
			}
			if c >= len(line) {
				return Pos{}, false
			}
			remaining--
		}
		found = c
	case 't':
		c := col
		for remaining > 0 {
			start := c + 1
			if remaining < count {
				start = c + 2
			}
			c2 := start
			for c2 < len(line) && line[c2] != ch {
				c2++
			}
			if c2 >= len(line) {
				return Pos{}, false
			}
			c = c2 - 1
			remaining--
		}
		found = c
	case 'F':
		c := col
		for remaining > 0 {
			c--
			for c >= 0 && line[c] != ch {
				c--
			}
			if c < 0 {
				return Pos{}, false
			}
			remaining--
		}
		found = c
	case 'T':
		c := col
		for remaining > 0 {
			start := c - 1
			if remaining < count {
				start = c - 2
			}
			c2 := start
			for c2 >= 0 && line[c2] != ch {
				c2--
			}
			if c2 < 0 {
				return Pos{}, false
			}
			c = c2 + 1
			remaining--
		}
		found = c
	}
	return Pos{b.Cursor.Row, found}, true
}

var bracketPairs = map[rune]rune{
	'(': ')', '[': ']', '{': '}',
}
var bracketOpen = map[rune]rune{
	')': '(', ']': '[', '}': '{',
}

// matchPercent implements '%': jump to the matching bracket on the current line/buffer.
func (b *Buffer) matchPercent() (Pos, bool) {
	row, col := b.Cursor.Row, b.Cursor.Col
	line := []rune(b.line(row))
	for col < len(line) {
		r := line[col]
		if close, ok := bracketPairs[r]; ok {
			return b.scanMatch(row, col, r, close, 1)
		}
		if open, ok := bracketOpen[r]; ok {
			return b.scanMatch(row, col, open, r, -1)
		}
		col++
	}
	return Pos{}, false
}

func (b *Buffer) scanMatch(row, col int, open, close rune, dir int) (Pos, bool) {
	depth := 0
	r, c := row, col
	for {
		line := []rune(b.line(r))
		if c >= 0 && c < len(line) {
			ch := line[c]
			if dir > 0 {
				if ch == open {
					depth++
				} else if ch == close {
					depth--
					if depth == 0 {
						return Pos{r, c}, true
					}
				}
			} else {
				if ch == close {
					depth++
				} else if ch == open {
					depth--
					if depth == 0 {
						return Pos{r, c}, true
					}
				}
			}
		}
		c += dir
		if c < 0 {
			r--
			if r < 0 {
				return Pos{}, false
			}
			c = b.lineLen(r) - 1
		} else if c >= b.lineLen(r) && dir > 0 {
			r++
			if r >= len(b.Lines) {
				return Pos{}, false
			}
			c = 0
		}
	}
}

// textObjectRange resolves things like iw, aw, i(, a(, i", a' into a
// [from,to) rune range on a single line, plus whether it's found.
func (b *Buffer) textObjectRange(around bool, kind rune) (from, to int, ok bool) {
	row := b.Cursor.Row
	line := []rune(b.line(row))
	col := b.Cursor.Col
	switch kind {
	case 'w':
		if len(line) == 0 {
			return 0, 0, false
		}
		if col >= len(line) {
			col = len(line) - 1
		}
		cls := wordClass(line[col])
		start, end := col, col
		for start > 0 && wordClass(line[start-1]) == cls {
			start--
		}
		for end < len(line) && wordClass(line[end]) == cls {
			end++
		}
		if around {
			trimmed := end
			for trimmed < len(line) && line[trimmed] == ' ' {
				trimmed++
			}
			if trimmed > end {
				end = trimmed
			} else {
				for start > 0 && line[start-1] == ' ' {
					start--
				}
			}
		}
		return start, end, true
	case '(', ')', '[', ']', '{', '}':
		open, close := kind, kind
		switch kind {
		case '(', ')':
			open, close = '(', ')'
		case '[', ']':
			open, close = '[', ']'
		case '{', '}':
			open, close = '{', '}'
		}
		oi := lastIndexBefore(line, col, open, close)
		if oi < 0 {
			return 0, 0, false
		}
		ci := indexClose(line, oi, open, close)
		if ci < 0 {
			return 0, 0, false
		}
		if around {
			return oi, ci + 1, true
		}
		return oi + 1, ci, true
	case '"', '\'':
		q := kind
		count := 0
		first, second := -1, -1
		for i, r := range line {
			if r == q {
				count++
				if count == 1 {
					first = i
				} else if count == 2 {
					second = i
					break
				}
			}
		}
		if first < 0 || second < 0 || col < first || col > second {
			return 0, 0, false
		}
		if around {
			return first, second + 1, true
		}
		return first + 1, second, true
	}
	return 0, 0, false
}

func lastIndexBefore(line []rune, col int, open, close rune) int {
	depth := 0
	if col < len(line) && line[col] == open {
		return col
	}
	for i := col; i >= 0; i-- {
		if i < len(line) {
			if line[i] == close && i != col {
				depth++
			} else if line[i] == open {
				if depth == 0 {
					return i
				}
				depth--
			}
		}
	}
	return -1
}

func indexClose(line []rune, from int, open, close rune) int {
	depth := 0
	for i := from; i < len(line); i++ {
		if line[i] == open {
			depth++
		} else if line[i] == close {
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	return -1
}
