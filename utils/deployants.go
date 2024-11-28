package lemin

import "fmt"

// SendAnts prints ants according to their paths
func (g *Graph) SendAnts(group [][]string, n int) {
	levels := n / len(group)
	if n%len(group) != 0 {
		levels++
	}

	var ants = make([]Ant, n+1)
	// IsInactive zero because ants start from 1
	ants[0].IsInactive = true
	id := 0

	for j := 0; j < levels; j++ {
		for _, path := range group {
			id++
			ants[id].Path = path
			ants[id].CurrentRoom = 0
			ants[id].IsInactive = false
			if id == n {
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
			if id == n {
				fmt.Println()
				if room == g.End {
					exit = true
				}
			}
			ants[id].CurrentRoom++
			taken[ants[id].Previous] = false
			if room != g.End {
				taken[room] = true
				ants[id].Previous = room
			}
			if room == g.End {
				ants[id].IsInactive = true
			}
		}
	}
}
