package util

import (
	"fmt"
	"path"
	"sort"
	"time"

	"github.com/manifoldco/promptui"
)

type HierarchyBuilder func(user string, date time.Time) string

var hierarchyOptions = map[string]HierarchyBuilder{
	"User > Year > Month > Day > Time": func(user string, date time.Time) string {
		return path.Join(
			user,
			date.Format("2006"),
			date.Format("01"),
			date.Format("02"),
			date.Format("15-04-05"),
		)
	},
	"Year > Month > Day > Time&User": func(user string, date time.Time) string {
		return path.Join(
			date.Format("2006"),
			date.Format("01"),
			date.Format("02"),
			fmt.Sprintf("%s_%s", date.Format("15-04-05"), user),
		)
	},
}

var hierarchyLabels []string

var promptHeirarchy = promptui.Select{
	Label: "How should the images be organized?",
	Items: hierarchyLabels,
}

func init() {
	for key := range hierarchyOptions {
		hierarchyLabels = append(hierarchyLabels, key)
	}
	sort.Strings(hierarchyLabels)
	promptHeirarchy.Items = hierarchyLabels
}

func PromptHierarchy() (HierarchyBuilder, error) {
	_, hierarchy, err := promptHeirarchy.Run()
	if err != nil {
		return nil, err
	}
	return hierarchyOptions[hierarchy], nil

}
