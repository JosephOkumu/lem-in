package main

import (
	"fmt"
	lemin "lemin/utils"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		filename := os.Args[1]
		antNum, g := lemin.ReadFile(filename)
		group := g.GetBestPath(antNum)
		g.SendAnts(group, antNum)
	} else if len(os.Args) < 2 {
		fmt.Println("Invalid number of arguments.")
	}
}
