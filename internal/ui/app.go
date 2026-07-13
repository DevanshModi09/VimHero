package ui

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"vimhero/internal/progress"
	"vimhero/pkg/curriculum"
	"vimhero/pkg/editor"
)

type screen int

const (
	screenDayList screen = iota
	screenChallengeList
	screenPlay
	screenThemes
)

const dayPageSize = 15

type Model struct {
	prog *progress.State
	days []curriculum.Day

	screen screen
	width  int
	height int

	dayCursor       int
	challengeCursor int
	themeCursor     int

	dayIdx       int
	challengeIdx int

	buf       *editor.Buffer
	challenge curriculum.Challenge

	won      bool
	wonStars int
	wonKeys  int
	quitting bool
}

func NewModel() Model {
	prog := progress.Load()
	applyTheme(prog.Theme)
	return Model{
		prog:   prog,
		days:   curriculum.All(),
		screen: screenDayList,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case tea.KeyMsg:
		key := msg.String()
		if key == "ctrl+c" {
			m.quitting = true
			return m, tea.Quit
		}
		switch m.screen {
		case screenDayList:
			return m.updateDayList(key)
		case screenChallengeList:
			return m.updateChallengeList(key)
		case screenPlay:
			return m.updatePlay(key)
		case screenThemes:
			return m.updateThemes(key)
		}
	}
	return m, nil
}

func (m Model) updateDayList(key string) (tea.Model, tea.Cmd) {
	themeRow := len(m.days)
	switch key {
	case "q":
		m.quitting = true
		return m, tea.Quit
	case "t":
		m.themeCursor = currentThemeIdx
		m.screen = screenThemes
		return m, nil
	case "j", "down":
		if m.dayCursor < themeRow {
			m.dayCursor++
		}
	case "k", "up":
		if m.dayCursor > 0 {
			m.dayCursor--
		}
	case "left", "h":
		page := m.dayPageOf(m.dayCursor)
		if page > 0 {
			m.dayCursor = (page - 1) * dayPageSize
		}
	case "right":
		page := m.dayPageOf(m.dayCursor)
		if page < m.dayPageCount()-1 {
			m.dayCursor = (page + 1) * dayPageSize
			if m.dayCursor > themeRow {
				m.dayCursor = themeRow
			}
		}
	case "enter", "l":
		if m.dayCursor == themeRow {
			m.themeCursor = currentThemeIdx
			m.screen = screenThemes
			return m, nil
		}
		day := m.days[m.dayCursor]
		if m.prog.IsDayUnlocked(day.Number) {
			m.dayIdx = m.dayCursor
			m.challengeCursor = 0
			m.screen = screenChallengeList
		}
	}
	return m, nil
}

func (m Model) updateThemes(key string) (tea.Model, tea.Cmd) {
	switch key {
	case "esc", "q", "t":
		m.screen = screenDayList
	case "j", "down":
		if m.themeCursor < len(themes)-1 {
			m.themeCursor++
		}
	case "k", "up":
		if m.themeCursor > 0 {
			m.themeCursor--
		}
	case "enter":
		applyTheme(m.themeCursor)
		m.prog.Theme = m.themeCursor
		_ = m.prog.Save()
		m.screen = screenDayList
	}
	return m, nil
}

func (m Model) updateChallengeList(key string) (tea.Model, tea.Cmd) {
	day := m.days[m.dayIdx]
	switch key {
	case "esc", "q":
		m.screen = screenDayList
	case "j", "down":
		if m.challengeCursor < len(day.Challenges)-1 {
			m.challengeCursor++
		}
	case "k", "up":
		if m.challengeCursor > 0 {
			m.challengeCursor--
		}
	case "enter", "l":
		m = m.enterChallenge(m.dayIdx, m.challengeCursor)
	}
	return m, nil
}

func (m Model) updatePlay(key string) (tea.Model, tea.Cmd) {
	if m.won {
		switch key {
		case "enter":
			day := m.days[m.dayIdx]
			if m.challengeIdx+1 < len(day.Challenges) {
				return m.enterChallenge(m.dayIdx, m.challengeIdx+1), nil
			}
			m.screen = screenDayList
			m.won = false
			return m, nil
		case "r":
			return m.enterChallenge(m.dayIdx, m.challengeIdx), nil
		case "esc":
			m.screen = screenChallengeList
			m.won = false
			return m, nil
		}
		return m, nil
	}

	if m.buf.Mode == editor.ModeNormal {
		switch key {
		case "esc":
			m.screen = screenChallengeList
			return m, nil
		case "ctrl+r":
			return m.enterChallenge(m.dayIdx, m.challengeIdx), nil
		}
	}

	m.buf.Input(key)
	if m.checkWin() {
		m.won = true
		m.wonKeys = m.buf.Keystrokes
		m.wonStars = curriculum.Stars(m.wonKeys, m.challenge.Par)
		day := m.days[m.dayIdx]
		m.prog.RecordClear(day.Number, m.challengeIdx, m.wonKeys, m.wonStars, len(day.Challenges))
		_ = m.prog.Save()
	}
	return m, nil
}

func (m Model) contentWidth() int {
	w := m.width
	if w <= 0 {
		w = 80
	}
	if w > 100 {
		w = 100
	}
	return w - 4
}

func (m Model) fullWidth() int {
	if m.width <= 0 {
		return 80
	}
	return m.width
}

func (m Model) checkWin() bool {
	switch m.challenge.Kind {
	case curriculum.KindGoal:
		return m.buf.Cursor.Row == m.challenge.GoalPos.Row && m.buf.Cursor.Col == m.challenge.GoalPos.Col
	case curriculum.KindEdit:
		return m.buf.Text() == strings.Join(m.challenge.Target, "\n")
	}
	return false
}

func (m Model) enterChallenge(dayIdx, challengeIdx int) Model {
	day := m.days[dayIdx]
	ch := day.Challenges[challengeIdx]
	m.dayIdx = dayIdx
	m.challengeIdx = challengeIdx
	m.challenge = ch
	m.buf = editor.New(append([]string(nil), ch.Start...), editor.Pos{Row: ch.CursorStart.Row, Col: ch.CursorStart.Col})
	m.won = false
	m.screen = screenPlay
	return m
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}
	var out string
	switch m.screen {
	case screenDayList:
		out = m.viewDayList()
	case screenChallengeList:
		out = m.viewChallengeList()
	case screenPlay:
		out = m.viewPlay()
	case screenThemes:
		out = m.viewThemes()
	}
	return fitToHeight(out, m.height)
}

func fitToHeight(s string, height int) string {
	if height <= 0 {
		return s
	}
	lines := strings.Split(s, "\n")
	if len(lines) <= height {
		return s
	}
	collapsed := make([]string, 0, len(lines))
	prevBlank := false
	for _, l := range lines {
		blank := strings.TrimSpace(l) == ""
		if blank && prevBlank {
			continue
		}
		collapsed = append(collapsed, l)
		prevBlank = blank
	}
	if len(collapsed) <= height {
		return strings.Join(collapsed, "\n")
	}
	if height < 1 {
		height = 1
	}
	return strings.Join(collapsed[:height], "\n")
}

func (m Model) viewDayList() string {
	var b strings.Builder
	b.WriteString(renderHeader("VimHero — Zero to Hero in 45 Days", m.fullWidth()))
	b.WriteString("\n\n")

	page := m.dayPageOf(m.dayCursor)
	pageCount := m.dayPageCount()
	start := page * dayPageSize
	end := start + dayPageSize
	if end > len(m.days) {
		end = len(m.days)
	}

	for i := start; i < end; i++ {
		day := m.days[i]
		marker := "  "
		if i == m.dayCursor {
			marker = "▸ "
		}
		label := fmt.Sprintf("Day %-3d %s", day.Number, day.Title)
		avgStars, cleared, total := m.dayStars(day)
		locked := !m.prog.IsDayUnlocked(day.Number)
		status := ""
		switch {
		case locked:
			status = lockedStyle.Render(" locked")
			label = lockedStyle.Render(label)
		case cleared == total:
			label = dayLineCleared.Render(label)
			status = starStyle.Render(fmt.Sprintf(" ★ %d/3 avg", avgStars))
		default:
			status = dimStyle.Render(fmt.Sprintf(" %d/%d cleared", cleared, total))
		}
		if i == m.dayCursor && !locked {
			marker = markerSelectedStyle.Render(marker)
			label = dayLineSelected.Render(label)
		}
		b.WriteString(marker + label + status + "\n")
	}
	b.WriteString("\n")
	if pageCount > 1 {
		b.WriteString(dimStyle.Render(fmt.Sprintf("Page %d/%d  (days %d-%d of %d)", page+1, pageCount, start+1, end, len(m.days))))
		b.WriteString("\n\n")
	}
	if m.prog.Streak > 0 {
		b.WriteString(streakStyle.Render(fmt.Sprintf("🔥 %d day streak", m.prog.Streak)) + "\n")
	}
	if page == pageCount-1 {
		themeMarker := ""
		themeLabel := "🎨 Change Theme"
		if m.dayCursor == len(m.days) {
			themeMarker = markerSelectedStyle.Render("▸ ")
			themeLabel = dayLineSelected.Render(themeLabel)
		}
		b.WriteString(themeMarker + themeLabel + dimStyle.Render(" (press t)") + "\n\n")
	}
	if pageCount > 1 {
		b.WriteString(helpStyle.Render("j/k move · ←/→ page · enter select · t theme · q quit"))
	} else {
		b.WriteString(helpStyle.Render("j/k move · enter select · t theme · q quit"))
	}
	b.WriteString("\n\n")
	b.WriteString(dimStyle.Render("Created by Devansh"))
	return b.String()
}

func (m Model) dayPageCount() int {
	pages := (len(m.days) + dayPageSize - 1) / dayPageSize
	if pages == 0 {
		pages = 1
	}
	return pages
}

func (m Model) dayPageOf(cursor int) int {
	if cursor >= len(m.days) {
		return m.dayPageCount() - 1
	}
	return cursor / dayPageSize
}

func (m Model) dayStars(day curriculum.Day) (avgStars, cleared, total int) {
	total = len(day.Challenges)
	sum := 0
	for i := range day.Challenges {
		if r, ok := m.prog.ChallengeResult(day.Number, i); ok && r.Cleared {
			cleared++
			sum += r.Stars
		}
	}
	if cleared > 0 {
		avgStars = sum / cleared
	}
	return
}

func (m Model) viewThemes() string {
	var list strings.Builder
	for i, t := range themes {
		marker := "  "
		name := challengeTitleStyle.Render(t.Name)
		if i == m.themeCursor {
			marker = markerSelectedStyle.Render("▸ ")
			name = challengeSelectedStyle.Render(t.Name)
		}
		current := ""
		if i == currentThemeIdx {
			current = dimStyle.Render(" (current)")
		}
		list.WriteString(fmt.Sprintf("%s%s %s%s\n", marker, themeSwatch(t.Accent), name, current))
	}

	var b strings.Builder
	b.WriteString(renderHeader("Choose a Theme", m.fullWidth()))
	b.WriteString("\n\n")
	b.WriteString(boxStyle.Render(strings.TrimRight(list.String(), "\n")))
	b.WriteString("\n\n")
	b.WriteString(helpStyle.Render("j/k move · enter apply · esc back"))
	return b.String()
}

func (m Model) viewChallengeList() string {
	day := m.days[m.dayIdx]
	var list strings.Builder
	for i, ch := range day.Challenges {
		marker := "  "
		title := challengeTitleStyle.Render(ch.Title)
		if i == m.challengeCursor {
			marker = markerSelectedStyle.Render("▸ ")
			title = challengeSelectedStyle.Render(ch.Title)
		}
		stars := ""
		if r, ok := m.prog.ChallengeResult(day.Number, i); ok && r.Cleared {
			stars = " " + starStyle.Render(strings.Repeat("★", r.Stars)+strings.Repeat("☆", 3-r.Stars))
		}
		list.WriteString(fmt.Sprintf("%s%s%s\n", marker, title, stars))
	}

	var b strings.Builder
	b.WriteString(renderHeader(fmt.Sprintf("Day %d — %s", day.Number, day.Title), m.fullWidth()))
	b.WriteString("\n\n")
	b.WriteString(subtitleStyle.Width(m.contentWidth()).Render(day.Summary))
	b.WriteString("\n\n")
	b.WriteString(boxStyle.Render(strings.TrimRight(list.String(), "\n")))
	b.WriteString("\n\n")
	b.WriteString(helpStyle.Render("j/k move · enter play · esc back"))
	return b.String()
}

func (m Model) viewPlay() string {
	var b strings.Builder
	day := m.days[m.dayIdx]
	b.WriteString(renderHeader(fmt.Sprintf("Day %d — %s", day.Number, m.challenge.Title), m.fullWidth()))
	b.WriteString("\n\n")
	b.WriteString(subtitleStyle.Width(m.contentWidth()).Render(m.challenge.Instructions))
	b.WriteString("\n")
	if len(m.challenge.NewKeys) > 0 {
		chips := make([]string, len(m.challenge.NewKeys))
		for i, k := range m.challenge.NewKeys {
			chips[i] = newKeyStyle.Render(k)
		}
		b.WriteString(dimStyle.Render("New: ") + strings.Join(chips, " "))
		b.WriteString("\n")
	}
	if m.challenge.Tip != "" {
		b.WriteString(tipStyle.Width(m.contentWidth()).Render("💡 " + m.challenge.Tip))
		b.WriteString("\n")
	}
	b.WriteString("\n")

	b.WriteString(boxStyle.Render(m.renderBuffer()))
	b.WriteString("\n\n")

	if m.challenge.Kind == curriculum.KindEdit {
		b.WriteString(dimStyle.Render("Target:") + "\n")
		b.WriteString(dimStyle.Render(strings.Join(m.challenge.Target, "\n")))
		b.WriteString("\n\n")
	}

	b.WriteString(m.modeBadge())
	b.WriteString(fmt.Sprintf("  keystrokes: %d  par: %d", m.buf.Keystrokes, m.challenge.Par))
	if m.buf.Mode == editor.ModeCommand {
		b.WriteString("\n" + m.buf.CommandLine)
	}
	if m.buf.StatusMsg != "" {
		b.WriteString("\n" + dimStyle.Render(m.buf.StatusMsg))
	}
	b.WriteString("\n\n")

	if m.won {
		stars := strings.Repeat("★", m.wonStars) + strings.Repeat("☆", 3-m.wonStars)
		msg := goodStyle.Render(fmt.Sprintf("🎉 Solved!  %d keystrokes  ", m.wonKeys)) + starStyle.Render(stars)
		b.WriteString(winBoxStyle.Render(msg))
		b.WriteString("\n")
		b.WriteString(helpStyle.Render("enter: next · r: retry · esc: back to day"))
	} else {
		b.WriteString(helpStyle.Render("esc (normal mode): back to challenge list · ctrl+r: restart"))
	}
	return b.String()
}

func (m Model) modeBadge() string {
	switch m.buf.Mode {
	case editor.ModeInsert:
		return modeInsert.Render("INSERT")
	case editor.ModeVisual:
		return modeVisual.Render("VISUAL")
	case editor.ModeVisualLine:
		return modeVisual.Render("V-LINE")
	case editor.ModeCommand:
		return modeCmd.Render("COMMAND")
	default:
		return modeNormal.Render("NORMAL")
	}
}

func (m Model) renderBuffer() string {
	var b strings.Builder
	gutterW := len(strconv.Itoa(len(m.buf.Lines)))
	for row, line := range m.buf.Lines {
		if row > 0 {
			b.WriteString("\n")
		}
		num := fmt.Sprintf("%*d", gutterW, row+1)
		if row == m.buf.Cursor.Row {
			b.WriteString(gutterCurrentStyle.Render(num))
		} else {
			b.WriteString(gutterStyle.Render(num))
		}
		b.WriteString(dimStyle.Render(" │ "))
		runes := []rune(line)
		if len(runes) == 0 {
			if m.isGoalCell(row, 0) {
				b.WriteString(goalCellStyle.Render(" "))
			} else if m.buf.Cursor.Row == row {
				b.WriteString(cursorCellStyle.Render(" "))
			} else {
				b.WriteString(" ")
			}
			continue
		}
		for col, r := range runes {
			cell := string(r)
			switch {
			case m.buf.Cursor.Row == row && m.buf.Cursor.Col == col:
				b.WriteString(cursorCellStyle.Render(cell))
			case m.isGoalCell(row, col):
				b.WriteString(goalCellStyle.Render(cell))
			default:
				b.WriteString(cell)
			}
		}
	}
	return b.String()
}

func (m Model) isGoalCell(row, col int) bool {
	return m.challenge.Kind == curriculum.KindGoal && m.challenge.GoalPos.Row == row && m.challenge.GoalPos.Col == col
}
