package lemin

import (
	"fmt"
	"os"
)

// GetBestPath gets one optimal path
func (graph *Graph) GetBestPath(ants int) [][]string {
	graph.GetAllPaths(graph.Start)

	for _, path := range Paths {
		if len(path) == 2 {
			fmt.Println(InputData)
			fmt.Println()
			for i := 1; i <= ants; i++ {
				fmt.Print("L", i, "-", graph.End, " ")
			}
			fmt.Println("")
			os.Exit(0)
		}
	}

	graph.GetIndependentPaths()
	graph.RmatchingPaths()

	var min int
	var heights = make(map[int][][]string)
	for _, group := range PathGroups {
		paths := len(group)
		max := len(group[len(group)-1])
		var fill = 0
		for _, path := range group {
			fill += max - len(path)
		}
		left := ants - fill
		if left == 0 {
			continue
		}
		height := max + left/paths
		if left%paths != 0 {
			height++
		}

		heights[height] = group
		// fmt.Println("group", group, "height:", height)
		min = height
	}

	for h := range heights {
		if h < min {
			min = h
		}
	}

	return heights[min]
}
