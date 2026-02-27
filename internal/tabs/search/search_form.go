package search

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func createSearchForm() *huh.Form {
	accentOrange := lipgloss.Color("#F97316")
	borderBlue := lipgloss.Color("#3B82F6")

	theme := huh.ThemeBase()

	theme.Focused.Title = theme.Focused.Title.Foreground(accentOrange).Bold(true)
	theme.Focused.TextInput.Prompt = theme.Focused.TextInput.Prompt.Foreground(accentOrange).Bold(true)
	theme.Focused.TextInput.Cursor = theme.Focused.TextInput.Cursor.Foreground(borderBlue)

	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("busqueda").
				Title("Búsqueda de canales:").
				Placeholder("ej: ESPN").
				Prompt("? "),
		),
	).
		WithShowHelp(false).
		WithTheme(theme)
}
