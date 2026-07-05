// Package ui implements the VimHero Bubble Tea application: a day/
// challenge menu and a play screen that wraps the editor engine.
package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"vimhero/internal/curriculum"
	"vimhero/internal/editor"
	"vimhero/internal/progress"
)

type screen int

const (
	screenDayList screen = iota
	screenChallengeList
	screenPlay
)

type Model struct {
	prog *progress.State
	days []curriculum.Day

	screen screen

	dayCursor       int
	challengeCursor int

	dayIdx       int // index into m.days of the day being played/browsed
	challengeIdx int

	buf       *editor.Buffer
	challenge curriculum.Challenge

	won      bool
	wonStars int
	wonKeys  int
	quitting bool
}

func NewModel() Model {
	return Model{
		prog:   progress.Load(),
		days:   curriculum.All(),
		screen: screenDayList,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
		}
	}
	return m, nil
}

func (m Model) updateDayList(key string) (tea.Model, tea.Cmd) {
	switch key {
	case "q":
		m.quitting = true
		return m, tea.Quit
	case "j", "down":
		if m.dayCursor < len(m.days)-1 {
			m.dayCursor++
		}
	case "k", "up":
		if m.dayCursor > 0 {
			m.dayCursor--
		}
	case "enter", "l":
		day := m.days[m.dayCursor]
		if m.prog.IsDayUnlocked(day.Number) {
			m.dayIdx = m.dayCursor
			m.challengeCursor = 0
			m.screen = screenChallengeList
		}
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
		case "esc":
			m.screen = screenChallengeList
			m.won = false
			return m, nil
		}
		return m, nil
	}

	if key == "esc" && m.buf.Mode == editor.ModeNormal {
		m.screen = screenChallengeList
		return m, nil
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
	switch m.screen {
	case screenDayList:
		return m.viewDayList()
	case screenChallengeList:
		return m.viewChallengeList()
	case screenPlay:
		return m.viewPlay()
	}
	return ""
}

func (m Model) viewDayList() string {
	var b strings.Builder
	b.WriteString(titleStyle.Render("VimHero — Zero to Hero in 45 Days"))
	b.WriteString("\n\n")
	for i, day := range m.days {
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
			status = lockedStyle.Render(" (locked)")
			label = lockedStyle.Render(label)
		case cleared == total:
			status = starStyle.Render(fmt.Sprintf(" ★ %d/3 avg", avgStars))
		default:
			status = dimStyle.Render(fmt.Sprintf(" (%d/%d cleared)", cleared, total))
		}
		if i == m.dayCursor && !locked {
			label = dayLineSelected.Render(label)
		}
		b.WriteString(marker + label + status + "\n")
	}
	b.WriteString("\n")
	if m.prog.Streak > 0 {
		b.WriteString(dimStyle.Render(fmt.Sprintf("🔥 %d day streak", m.prog.Streak)) + "\n")
	}
	b.WriteString(helpStyle.Render("j/k move · enter select · q quit"))
	return b.String()
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

func (m Model) viewChallengeList() string {
	day := m.days[m.dayIdx]
	var b strings.Builder
	b.WriteString(titleStyle.Render(fmt.Sprintf("Day %d — %s", day.Number, day.Title)))
	b.WriteString("\n\n")
	b.WriteString(subtitleStyle.Render(day.Summary))
	b.WriteString("\n\n")
	for i, ch := range day.Challenges {
		marker := "  "
		if i == m.challengeCursor {
			marker = "▸ "
		}
		stars := ""
		if r, ok := m.prog.ChallengeResult(day.Number, i); ok && r.Cleared {
			stars = " " + starStyle.Render(strings.Repeat("★", r.Stars)+strings.Repeat("☆", 3-r.Stars))
		}
		b.WriteString(fmt.Sprintf("%s%s%s\n", marker, ch.Title, stars))
	}
	b.WriteString("\n")
	b.WriteString(helpStyle.Render("j/k move · enter play · esc back"))
	return b.String()
}

func (m Model) viewPlay() string {
	var b strings.Builder
	day := m.days[m.dayIdx]
	b.WriteString(titleStyle.Render(fmt.Sprintf("Day %d — %s", day.Number, m.challenge.Title)))
	b.WriteString("\n\n")
	b.WriteString(subtitleStyle.Render(m.challenge.Instructions))
	b.WriteString("\n")
	if len(m.challenge.NewKeys) > 0 {
		b.WriteString(dimStyle.Render("New: " + strings.Join(m.challenge.NewKeys, "  ")))
		b.WriteString("\n")
	}
	if m.challenge.Tip != "" {
		b.WriteString(tipStyle.Render("💡 " + m.challenge.Tip))
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
		b.WriteString(goodStyle.Render(fmt.Sprintf("Solved! %d keystrokes %s", m.wonKeys, strings.Repeat("★", m.wonStars)+strings.Repeat("☆", 3-m.wonStars))))
		b.WriteString("\n")
		b.WriteString(helpStyle.Render("enter: next · esc: back to day"))
	} else {
		b.WriteString(helpStyle.Render("esc (normal mode): back to challenge list"))
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
	for row, line := range m.buf.Lines {
		if row > 0 {
			b.WriteString("\n")
		}
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
