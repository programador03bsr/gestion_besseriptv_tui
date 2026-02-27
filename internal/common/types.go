package common

type SearchResultMsg struct {
	Results []string
	Err     error
}

type ExecutionResultMsg struct {
	Logs []string
}

type StreamLogMsg string

type ProcessFinishedMsg struct{}
