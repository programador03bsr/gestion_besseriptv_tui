package tui

import (
	"github.com/besser/canales-cli/internal/common"
	"github.com/besser/canales-cli/internal/tabs/search"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RootModel struct {
	width, height int
	tabs          []string
	activeTab     int
	searchTab     search.SearchTabModel

	systemLogs []string
}

func GetInitialModel() RootModel {
	return RootModel{
		tabs:      []string{"Gestión de Canales IPTV Besser"},
		activeTab: 0,
		searchTab: search.New(),
	}
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		m.searchTab, _ = m.searchTab.Update(msg)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "right":
			// Lógica para cambiar m.activeTab...
			return m, nil
		}
	case common.StreamLogMsg:
		m.systemLogs = append(m.systemLogs, string(msg))
		if len(m.systemLogs) > 50 {
			m.systemLogs = m.systemLogs[len(m.systemLogs)-50:]
		}
	}

	if m.activeTab == 0 {
		var childCmd tea.Cmd
		m.searchTab, childCmd = m.searchTab.Update(msg)
		cmd = childCmd
	}

	return m, cmd
}

func (m RootModel) View() string {
	if m.width == 0 {
		return "Inicializando..."
	}

	const logPanelHeight = 8
	mainWindowHeight := m.height - logPanelHeight - 5
	logStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(common.InactiveColor).
		Foreground(common.TextWhite).
		Width(m.width-2).
		Height(logPanelHeight-2).
		Padding(0, 1)

	var renderedTabs []string
	for i, t := range m.tabs {
		if i == m.activeTab {
			renderedTabs = append(renderedTabs, common.ActiveTabStyle.Render(t))
		} else {
			renderedTabs = append(renderedTabs, common.InactiveTabStyle.Render(t))
		}
	}
	tabBar := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	var content string

	if m.activeTab == 0 {
		content = m.searchTab.View()
	}

	mainWindow := m.BuildWindowStyle().
		Width(m.width - 2).
		Height(mainWindowHeight).
		Render(content)

	var logContent string
	if len(m.systemLogs) == 0 {
		logContent = "Sistema listo. Esperando tareas..."
	} else {
		maxVisibleLogs := logPanelHeight - 2
		startIdx := 0
		if len(m.systemLogs) > maxVisibleLogs {
			startIdx = len(m.systemLogs) - maxVisibleLogs
		}

		for _, log := range m.systemLogs[startIdx:] {
			logContent += log + "\n"
		}
	}
	logPanel := logStyle.Render(logContent)

	layout := lipgloss.JoinVertical(lipgloss.Left,
		tabBar,
		mainWindow,
		logPanel,
	)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Left, lipgloss.Bottom,
		layout,
		lipgloss.WithWhitespaceBackground(common.BgDarkBlue),
		lipgloss.WithWhitespaceChars(" "),
	)
}

func (m RootModel) Init() tea.Cmd {
	return m.searchTab.Init()
}

func (m RootModel) BuildWindowStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, true, true).
		BorderForeground(common.BorderBlue).
		Foreground(common.TextWhite).
		Background(common.BgDarkBlue).
		Padding(1, 2).
		Width(m.width - 2).
		Height(m.height - 6)
}
