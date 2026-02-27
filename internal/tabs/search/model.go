// Package search contains the Model of the TUI
package search

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type SearchTabModel struct {
	currentStep step
	searchForm  *huh.Form
	resultsForm *huh.Form

	loader        spinner.Model
	finalMessages []string
	logChan       chan string
}

func New() SearchTabModel {
	form := createSearchForm()
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#F97316"))
	return SearchTabModel{
		currentStep: stepTyping,
		searchForm:  form,
		loader:      s,
	}
}

func (m SearchTabModel) View() string {
	var views []string
	if m.searchForm != nil {
		views = append(views, m.searchForm.View())
	}
	switch m.currentStep {
	// case stepTyping:
	// 	if m.searchForm != nil {
	// 		var resultsView string = "No hay resultados aún..."
	// 		if m.resultsForm != nil {
	// 			resultsView = m.resultsForm.View()
	// 		}
	// 		return lipgloss.JoinVertical(
	// 			lipgloss.Left,
	// 			"\n",
	// 			lipgloss.NewStyle().Foreground(lipgloss.Color("#34D399")).Render("Escribe tu consulta y presiona Enter:"),
	// 			m.searchForm.View(),
	// 			"\n",
	// 			lipgloss.NewStyle().Foreground(lipgloss.Color("#34D399")).Render("Resultados:"),
	// 			"\n",
	// 			resultsView,
	// 		)
	// 	}
	case stepSelecting:
		if m.resultsForm != nil {
			views = append(views, "\n"+m.resultsForm.View())
		} else {
			views = append(views, "\n"+lipgloss.NewStyle().Foreground(lipgloss.Color("#F97316")).Render("No hay resultados para mostrar."))
		}
	case stepProcessing:
		views = append(views, "\n"+m.loader.View()+" Buscando canales...")

	case stepDone:
		var rendered []string
		for _, msg := range m.finalMessages {
			rendered = append(rendered, lipgloss.NewStyle().Foreground(lipgloss.Color("#34D399")).Render(msg))
		}
		views = append(views, "\n"+lipgloss.JoinVertical(lipgloss.Left, rendered...))
	}

	return lipgloss.JoinVertical(lipgloss.Left, views...)
}

func (m SearchTabModel) Init() tea.Cmd {
	if m.searchForm != nil {
		return m.searchForm.Init()
	}
	// if m.resultsForm != nil {
	// 	return m.resultsForm.Init()
	// }
	return nil
}
