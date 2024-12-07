package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function that gives ants names
func (graph *Graph) NameAnt() {
	antsNum := graph.AntCount
	for i := 1; i <= antsNum; i++ {
		graph.AntNames = append(graph.AntNames, "L"+strconv.Itoa(i))
	}
}

// This function reads and prints the original file contents before simulating ant movement
func (graph *Graph) PrintFileContents(filename string) {
	file, err := os.Open("./examples/" + filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())		
	}
	fmt.Println()
}

// This function moves ants
func SimulateAntMovement(paths [][]string, antDistribution [][]int) {
	type AntPosition struct {
		ant  int
		path int
		step int
	}
	count := 0
	var antPositions []AntPosition
	for pathIndex, ants := range antDistribution {
		for _, ant := range ants {
			antPositions = append(antPositions, AntPosition{ant, pathIndex, 0})
		}
	}
	for len(antPositions) > 0 {
		var moves []string
		var newPositions []AntPosition
		usedLinks := make(map[string]bool)
		for _, pos := range antPositions {
			if pos.step < len(paths[pos.path])-1 {
				currentRoom := paths[pos.path][pos.step]
				nextRoom := paths[pos.path][pos.step+1]
				link := currentRoom + "-" + nextRoom
				if !usedLinks[link] {
					moves = append(moves, fmt.Sprintf("L%d-%s", pos.ant, nextRoom))
					newPositions = append(newPositions, AntPosition{pos.ant, pos.path, pos.step + 1})
					usedLinks[link] = true
				} else {
					newPositions = append(newPositions, pos)
				}
			}
		}
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
		antPositions = newPositions
		count++
	}
}