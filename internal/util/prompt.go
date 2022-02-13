package util

import (
	"context"
	"fmt"

	"github.com/densestvoid/groupme"
	"github.com/manifoldco/promptui"
)

func PromptGroup(client *groupme.Client) (groupId string, err error) {
	ctx := context.Background()

	groups, err := client.IndexGroups(ctx, &groupme.GroupsQuery{PerPage: 100})
	if err != nil {
		return
	}

	names := make([]string, len(groups))
	for i, g := range groups {
		names[i] = fmt.Sprintf("%s (id: %s)", g.Name, g.ID)
	}

	p := promptui.Select{
		Label: "Pick a target group",
		Items: names,
	}

	i, _, err := p.Run()
	if err != nil {
		return
	}

	groupId = string(groups[i].ID)
	return
}
