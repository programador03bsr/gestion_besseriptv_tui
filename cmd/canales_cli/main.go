package main

import (
	"fmt"
	"os"

	"github.com/besser/canales-cli/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.GetInitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error al inicializar la aplicación: %v", err)
		os.Exit(1)
	}
}
