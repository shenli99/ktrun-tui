package main

import (
	"fmt"
	"os"

	"github.com/shenli99/ktrun-tui/view"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(view.InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
