package ui

import "github.com/charmbracelet/lipgloss"

var (
	colorAccent  = lipgloss.Color("212")
	colorGood    = lipgloss.Color("42")
	colorWarn    = lipgloss.Color("214")
	colorMuted   = lipgloss.Color("240")
	colorLocked  = lipgloss.Color("237")
	colorCursor  = lipgloss.Color("235")
	colorGoal    = lipgloss.Color("214")
	colorCommand = lipgloss.Color("39")

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("255")).
			Background(colorAccent).
			Padding(0, 2)

	subtitleStyle = lipgloss.NewStyle().Foreground(colorMuted).Italic(true)

	dayLineStyle    = lipgloss.NewStyle().PaddingLeft(2)
	dayLineSelected = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255"))
	lockedStyle     = lipgloss.NewStyle().Foreground(colorLocked)

	starStyle  = lipgloss.NewStyle().Foreground(colorWarn)
	dimStyle   = lipgloss.NewStyle().Foreground(colorMuted)
	goodStyle  = lipgloss.NewStyle().Foreground(colorGood).Bold(true)
	boxStyle   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2).BorderForeground(colorAccent)
	helpStyle  = lipgloss.NewStyle().Foreground(colorMuted)
	modeNormal = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255")).Background(lipgloss.Color("62")).Padding(0, 1)
	modeInsert = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("0")).Background(colorGood).Padding(0, 1)
	modeVisual = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("0")).Background(colorWarn).Padding(0, 1)
	modeCmd    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255")).Background(colorCommand).Padding(0, 1)

	cursorCellStyle = lipgloss.NewStyle().Reverse(true)
	goalCellStyle   = lipgloss.NewStyle().Foreground(colorGoal).Bold(true)

	tipStyle = lipgloss.NewStyle().Foreground(colorCommand)
)
