package search

import (
	"github.com/besser/canales-cli/internal/common"
	"github.com/charmbracelet/huh"
)

func newSearchResultsMultiSelect(results common.SearchResultMsg) *huh.Form {
	var options []huh.Option[string]
	for _, opt := range results.Results {
		options = append(options, huh.NewOption(opt, opt))
	}
	return huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Key("Canales").
				Options(options...).
				Title("Canales").
				Limit(20),
		),
	).WithShowHelp(false)
}
