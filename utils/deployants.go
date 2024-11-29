package lemin

import "fmt"

// Distributes ants to the respective paths
func (graph *Graph) DeployAnts(group [][]string, num int) {
	levels := num / len(group)
	if num%len(group) != 0 {
		levels++
	}

	var ants = make([]Ant, num+1)
	// IsInactive zero because ants start from 1
	ants[0].IsInactive = true
	id := 0

	for j := 0; j < levels; j++ {
		for _, path := range group {
			id++
			ants[id].Path = path
			ants[id].CurrentRoom = 0
			ants[id].IsInactive = false
			if id == num {
				break
			}
		}
	}

	fmt.Println(InputData)
	fmt.Println()

	exit := false
	var taken = make(map[string]bool)
	for !exit {
		for id, ant := range ants {
			if ant.IsInactive {
				continue
			}
			room := ant.Path[ant.CurrentRoom]
			if taken[room] {
				fmt.Println()
				break
			}
			fmt.Print("L", id, "-", room, " ")
			if id == num {
				fmt.Println()
				if room == graph.End {
					exit = true
				}
			}
			ants[id].CurrentRoom++
			taken[ants[id].Previous] = false
			if room != graph.End {
				taken[room] = true
				ants[id].Previous = room
			}
			if room == graph.End {
				ants[id].IsInactive = true
			}
		}
	}
}
