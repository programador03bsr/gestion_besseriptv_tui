package search

import (
	"fmt"
	"sync"
	"time"

	"github.com/besser/canales-cli/internal/common"
	tea "github.com/charmbracelet/bubbletea"
)

func RunSearchCMD(query string) tea.Cmd {
	return func() tea.Msg {
		resultadosSimulados := []string{
			"resultado_1_para_" + query,
			"resultado_2_para_" + query,
			"resultado_3_para_" + query,
			"resultado_4_para_" + query,
			"resultado_5_para_" + query,
			"resultado_6_para_" + query,
			"resultado_7_para_" + query,
			"resultado_8_para_" + query,
			"resultado_9_para_" + query,
			"resultado_10_para_" + query,
			"resultado_11_para_" + query,
			"resultado_12_para_" + query,
			"resultado_13_para_" + query,
			"resultado_14_para_" + query,
			"resultado_15_para_" + query,
			"resultado_16_para_" + query,
			"resultado_17_para_" + query,
			"resultado_18_para_" + query,
			"resultado_19_para_" + query,
			"resultado_20_para_" + query,
			"resultado_21_para_" + query,
			"resultado_22_para_" + query,
			"resultado_23_para_" + query,
			"resultado_24_para_" + query,
			"resultado_25_para_" + query,
			"resultado_26_para_" + query,
			"resultado_27_para_" + query,
			"resultado_28_para_" + query,
			"resultado_29_para_" + query,
			"resultado_30_para_" + query,
		}
		return common.SearchResultMsg{
			Results: resultadosSimulados,
			Err:     nil,
		}
	}
}

func RunSelectedChannelsCMD(channels []string, logsChan chan string) tea.Cmd {
	return func() tea.Msg {
		wg := sync.WaitGroup{}

		for i, channel := range channels {
			wg.Add(1)
			go func(ch string, idx int) {
				defer wg.Done()
				time.Sleep(time.Duration(idx+1) * time.Second)
				logsChan <- fmt.Sprintf("Canal procesado: %s", ch)
			}(channel, i)
		}

		wg.Wait()
		close(logsChan)

		return nil
	}
}

func WaitForLogCmd(logChan chan string) tea.Cmd {
	return func() tea.Msg {
		log, ok := <-logChan

		if !ok {
			return common.ProcessFinishedMsg{}
		}

		return common.StreamLogMsg(log)
	}
}
