package common

import "github.com/charmbracelet/lipgloss"

var (
	BgColor       = lipgloss.Color("#1E1E2E")
	FgColor       = lipgloss.Color("#CDD6F4")
	AccentColor   = lipgloss.Color("62")
	InactiveColor = lipgloss.Color("240")

	BgDarkBlue   = lipgloss.Color("#0F172A") // Azul oscuro profundo para el fondo
	BorderBlue   = lipgloss.Color("#3B82F6") // Azul brillante para líneas
	AccentOrange = lipgloss.Color("#F97316") // Naranja vibrante para activos
	TextWhite    = lipgloss.Color("#F8FAFC") // Blanco roto para leer bien
	InactiveText = lipgloss.Color("#64748B") // Gris/Azulado para lo inactivo
)

var ActiveTabStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), true, true, false, true).
	BorderForeground(AccentOrange).
	Padding(0, 2).
	Foreground(AccentOrange).
	Background(BgDarkBlue).
	Bold(true)

var InactiveTabStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), true, true, false, true).
	BorderForeground(BorderBlue).
	Padding(0, 2).
	Foreground(InactiveText).
	Background(BgDarkBlue)
