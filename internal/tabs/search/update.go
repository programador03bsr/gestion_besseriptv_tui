package search

import (
	"github.com/besser/canales-cli/internal/common"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m SearchTabModel) Update(msg tea.Msg) (SearchTabModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case spinner.TickMsg:
		if m.currentStep == stepProcessing {
			var cmd tea.Cmd
			m.loader, cmd = m.loader.Update(msg)
			return m, cmd
		}
	case common.ProcessFinishedMsg:
		m.searchForm = createSearchForm()
		m.currentStep = stepTyping
		return m, m.searchForm.Init()

	case common.StreamLogMsg:
		return m, WaitForLogCmd(m.logChan)

	}

	if searchMsg, ok := msg.(common.SearchResultMsg); ok {
		if searchMsg.Err == nil {
			m.resultsForm = newSearchResultsMultiSelect(searchMsg)
			m.currentStep = stepSelecting
			return m, m.resultsForm.Init()
		}

		return m, nil
	}

	switch m.currentStep {
	case stepTyping:
		if m.searchForm != nil {
			form, cmd := m.searchForm.Update(msg)
			if f, ok := form.(*huh.Form); ok {
				m.searchForm = f
				cmds = append(cmds, cmd)
				if m.searchForm.State == huh.StateCompleted {
					query := m.searchForm.GetString("busqueda")
					cmds = append(cmds, RunSearchCMD(query))
				}
				// return m, tea.Batch(cmds...)
			}
		}

	case stepSelecting:
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			if keyMsg.String() == "esc" {
				m.currentStep = stepTyping
				cmds = append(cmds, m.searchForm.Init())
				return m, tea.Batch(cmds...)
			}
		}

		if m.resultsForm != nil {
			results, cmd := m.resultsForm.Update(msg)
			if r, ok := results.(*huh.Form); ok {
				m.resultsForm = r
				cmds = append(cmds, cmd)

				if m.resultsForm.State == huh.StateCompleted {
					selectedChannels := m.resultsForm.Get("Canales")
					if selectedChannels, ok := selectedChannels.([]string); ok {
						m.currentStep = stepProcessing
						m.logChan = make(chan string)
						cmds = append(cmds,
							m.loader.Tick,
							RunSelectedChannelsCMD(selectedChannels, m.logChan),
							WaitForLogCmd(m.logChan),
						)
					}
				}

			}
		}
	case stepDone:
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
			m.currentStep = stepTyping
		}
	}
	return m, tea.Batch(cmds...)
}
