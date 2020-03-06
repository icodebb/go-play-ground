/**
 * Menu package for user menu.
 * It runs the selected test from user, after done, show the menu again.
 *
 *
 */

package menu

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

type choice struct {
	Objective   string
	Description string
	Index       int
}

// PrintMenu prints the customized menu.
func PrintMenu() {
	peppers := []choice{
		{Objective: "Bell Pepper", Description: "0", Index: 0},
		{Objective: "Banana Pepper", Description: "100", Index: 1},
		{Objective: "Poblano", Description: "1000", Index: 2},
		{Objective: "Jalapeño", Description: "3500", Index: 3},
		{Objective: "Aleppo", Description: "10000", Index: 4},
		{Objective: "Tabasco", Description: "30000", Index: 5},
		{Objective: "Malagueta", Description: "50000", Index: 6},
		{Objective: "Habanero", Description: "100000", Index: 7},
		{Objective: "Red Savina Habanero", Description: "350000", Index: 8},
		{Objective: "Dragon’s Breath", Description: "855000", Index: 9},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Objective | cyan }} ({{ .Description | red }})",
		Inactive: "  {{ .Objective | cyan }} ({{ .Description | red }})",
		Selected: "\U0001F336 {{ .Objective | red | cyan }}",
		Details: `
--------- Pepper ----------
{{ "Objective:" | faint }}   {{ .Objective }}
{{ "Heat Unit:" | faint }}  {{ .Description }}
{{ "Index:" | faint }}    {{ .Index }}`,
	}

	searcher := func(input string, idx int) bool {
		pepper := peppers[idx]
		objective := strings.Replace(strings.ToLower(pepper.Objective), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(objective, input)
	}

	prompt := promptui.Select{
		Label:     "Spicy Level",
		Items:     peppers,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose number %d: %s\n", i+1, peppers[i].Objective)

}
