/**
 * Menu package for user menu.
 * It runs the selected test from user, after done, show the menu again.
 *
 * It works well on Linux, but not on Windows.
 * Another issue is the unicode, it only show half of the symbol.
 * e.g. the pepper \U00002714.
 * So I use '->' instead.
 */

package menu

import (
	"strings"

	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
)

type choice struct {
	Target      string
	Description string
	Index       int
}

// PrintMenu prints the customized menu, return the index of user's choice.
func PrintMenu() int {
	choices := []choice{
		{Target: "Info", Description: "Show program and system information.", Index: 0},
		{Target: "Simple Test", Description: "Test logging and Fabonacci.", Index: 1},
		{Target: "Numeric Test", Description: "Test numeric functions such as random, etc.", Index: 2},
		{Target: "Datetime Test", Description: "Test date and time.", Index: 3},
		{Target: "Channel Test", Description: "Test channel feature.", Index: 4},
		{Target: "Tabasco", Description: "30000", Index: 5},
		{Target: "Malagueta", Description: "50000", Index: 6},
		{Target: "Habanero", Description: "100000", Index: 7},
		{Target: "Red Savina Habanero", Description: "350000", Index: 8},
		{Target: "Dragon’s Breath", Description: "855000", Index: 9},
		{Target: "Exit", Description: "Exit the program.", Index: 99},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "-> {{ .Target | yellow | bold }} ({{ .Description | magenta }})",
		Inactive: "  {{ .Target | cyan }} ({{ .Description | faint }})",
		Selected: "-> {{ .Target | green | bold }}",
		Details: `
-------------------   Annotation  --------------------
{{ "Target:" | green }}   {{ .Target }}
{{ "Description:" | green }}  {{ .Description }}
{{ "Index:" | green }}    {{ .Index }}
------------------------------------------------------`,
	}

	searcher := func(input string, idx int) bool {
		item := choices[idx]
		objective := strings.Replace(strings.ToLower(item.Target), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(objective, input)
	}

	prompt := promptui.Select{
		Label:     "Select the what you want to test",
		Items:     choices,
		Templates: templates,
		Size:      10, // how many items in a page
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		return -1
	}

	log.Infof("You chose %s with index: %d\n", choices[i].Target, choices[i].Index)
	return choices[i].Index
}
