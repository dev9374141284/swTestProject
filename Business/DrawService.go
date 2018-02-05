package Business

import (
	"TestProject/Models"
	"fmt"
	"math/rand"
)

func GetGroups(teams []Models.Team, minCount int)([]Models.Group){
	var group Models.Group
	if minCount == 0{
		minCount = 1
	}
	var groupsCount = int(len(teams) / minCount)
	if groupsCount == 0{
		groupsCount = 1
	}
	groups := make([]Models.Group, groupsCount)

	for index,_ := range groups{
		groups[index].Name = fmt.Sprintf("Group %d", index + 1)
	}

	shuffledTeams := make([]Models.Team, len(teams))
	perm := rand.Perm(len(teams))
	for i, v := range perm {
		shuffledTeams[v] = teams[i]
	}

	if groupsCount < 2 {
		group.Teams = shuffledTeams
		group.Name = "Group 1"
		groups[0] = group
	} else{
		var groupIndex int = 0
		for _,element := range shuffledTeams{
			if groupIndex == groupsCount{
				groupIndex = 0
			}
			groups[groupIndex].Teams = append(groups[groupIndex].Teams, element)
			groupIndex++
		}
	}

	return groups
}
