package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"vimhero/internal/ui"
)

func main() {
	p := tea.NewProgram(ui.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "vimhero:", err)
		os.Exit(1)
	}
}
