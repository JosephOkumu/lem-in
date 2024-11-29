package main

import (
	"fmt"
	lemin "lemin/utils"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		filename := os.Args[1]
		antNum, graph := lemin.ReadFile(filename)
		group := graph.GetBestPath(antNum)
		graph.DeployAnts(group, antNum)
	} else if len(os.Args) < 2 {
		fmt.Println("Invalid number of arguments.")
	}
}
