package editor

import (
	"regexp"
	"strconv"
	"strings"
)

func (b *Buffer) handleCommandMode(key string) {
	switch key {
	case "esc":
		b.Mode = ModeNormal
		b.CommandLine = ""
		b.resetPending()
		return
	case "enter":
		line := b.CommandLine
		kind := b.commandKind
		b.Mode = ModeNormal
		b.CommandLine = ""
		b.resetPending()
		if kind == ':' {
			b.executeExCommand(line)
		} else {
			b.searchPattern = line
			b.repeatSearch(1)
		}
		return
	case "backspace":
		if len(b.CommandLine) > 0 {
			r := []rune(b.CommandLine)
			b.CommandLine = string(r[:len(r)-1])
		} else {
			b.Mode = ModeNormal
		}
		return
	}
	r := []rune(key)
	if len(r) == 1 {
		b.CommandLine += key
	}
}

func (b *Buffer) executeExCommand(line string) {
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}
	if n, err := strconv.Atoi(line); err == nil {
		row := n - 1
		if row < 0 {
			row = 0
		}
		if row >= len(b.Lines) {
			row = len(b.Lines) - 1
		}
		b.Cursor = Pos{row, b.lineFirstNonBlank(row)}
		return
	}

	rangeAll := false
	rest := line
	if strings.HasPrefix(rest, "%") {
		rangeAll = true
		rest = rest[1:]
	}
	if strings.HasPrefix(rest, "s") {
		rest = rest[1:]
		b.execSubstitute(rest, rangeAll)
		return
	}
	if strings.HasPrefix(rest, "g/") || strings.HasPrefix(line, "g/") {
		b.execGlobal(strings.TrimPrefix(line, "g"))
		return
	}
	b.StatusMsg = "unknown command: :" + line
}

func (b *Buffer) execSubstitute(rest string, all bool) {
	if len(rest) == 0 || rest[0] != '/' {
		b.StatusMsg = "invalid :s syntax, use /pattern/replacement/"
		return
	}
	parts := splitUnescaped(rest[1:], '/')
	if len(parts) < 2 {
		b.StatusMsg = "invalid :s syntax, use /pattern/replacement/"
		return
	}
	pattern, repl := parts[0], parts[1]
	flags := ""
	if len(parts) > 2 {
		flags = parts[2]
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		b.StatusMsg = "bad pattern: " + err.Error()
		return
	}
	global := strings.Contains(flags, "g")

	b.pushUndo()
	row1, row2 := b.Cursor.Row, b.Cursor.Row
	if all {
		row1, row2 = 0, len(b.Lines)-1
	}
	changed := false
	for r := row1; r <= row2 && r < len(b.Lines); r++ {
		l := b.Lines[r]
		var newLine string
		if global {
			newLine = re.ReplaceAllString(l, repl)
		} else {
			found := false
			newLine = re.ReplaceAllStringFunc(l, func(m string) string {
				if found {
					return m
				}
				found = true
				sub := re.FindStringSubmatchIndex(l)
				return string(re.ExpandString(nil, repl, l, sub))
			})
		}
		if newLine != l {
			changed = true
			b.Lines[r] = newLine
		}
	}
	if !changed {
		b.StatusMsg = "pattern not found: " + pattern
	}
}

func (b *Buffer) execGlobal(rest string) {
	parts := splitUnescaped(strings.TrimPrefix(rest, "/"), '/')
	if len(parts) < 1 {
		return
	}
	pattern := parts[0]
	cmd := "p"
	if len(parts) > 1 {
		cmd = strings.TrimSpace(parts[1])
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		b.StatusMsg = "bad pattern: " + err.Error()
		return
	}
	if cmd != "d" {
		return
	}
	b.pushUndo()
	kept := b.Lines[:0:0]
	for _, l := range b.Lines {
		if !re.MatchString(l) {
			kept = append(kept, l)
		}
	}
	if len(kept) == 0 {
		kept = []string{""}
	}
	b.Lines = kept
	b.clampCursor()
}

func splitUnescaped(s string, sep byte) []string {
	var parts []string
	cur := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+1 < len(s) {
			cur.WriteByte(s[i])
			cur.WriteByte(s[i+1])
			i++
			continue
		}
		if s[i] == sep {
			parts = append(parts, cur.String())
			cur.Reset()
			continue
		}
		cur.WriteByte(s[i])
	}
	parts = append(parts, cur.String())
	return parts
}

func (b *Buffer) searchWord(dir int) {
	line := []rune(b.line(b.Cursor.Row))
	col := b.Cursor.Col
	if len(line) == 0 || col >= len(line) || wordClass(line[col]) != 1 {
		b.StatusMsg = "no word under cursor"
		return
	}
	start, end := col, col
	for start > 0 && wordClass(line[start-1]) == 1 {
		start--
	}
	for end < len(line) && wordClass(line[end]) == 1 {
		end++
	}
	b.searchPattern = `\b` + regexp.QuoteMeta(string(line[start:end])) + `\b`
	b.searchDir = dir
	b.Cursor.Col = start
	b.repeatSearch(1)
}

func (b *Buffer) repeatSearch(dir int) {
	if b.searchPattern == "" {
		return
	}
	re, err := regexp.Compile(b.searchPattern)
	if err != nil {
		b.StatusMsg = "bad pattern: " + err.Error()
		return
	}
	effDir := dir * b.searchDir
	if effDir == 0 {
		effDir = 1
	}

	var matches []Pos
	for row, l := range b.Lines {
		for _, loc := range re.FindAllStringIndex(l, -1) {
			matches = append(matches, Pos{row, byteToRuneCol(l, loc[0])})
		}
	}
	if len(matches) == 0 {
		b.StatusMsg = "pattern not found: " + b.searchPattern
		return
	}

	cur := b.Cursor
	if effDir > 0 {
		for _, m := range matches {
			if m.Row > cur.Row || (m.Row == cur.Row && m.Col > cur.Col) {
				b.Cursor = m
				return
			}
		}
		b.Cursor = matches[0]
		return
	}
	for i := len(matches) - 1; i >= 0; i-- {
		m := matches[i]
		if m.Row < cur.Row || (m.Row == cur.Row && m.Col < cur.Col) {
			b.Cursor = m
			return
		}
	}
	b.Cursor = matches[len(matches)-1]
}

func byteToRuneCol(s string, byteIdx int) int {
	if byteIdx <= 0 {
		return 0
	}
	if byteIdx >= len(s) {
		return len([]rune(s))
	}
	return len([]rune(s[:byteIdx]))
}
