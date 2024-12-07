package main

import (
	"fmt"
	"lemin/functions"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: invalid data format, Usage: go run . [filename.txt]")
		return
	}
	// Storing the filename into a variable
	filename := os.Args[1]
	Graph := functions.Graph{}
	// Checking if the file is valid or not
	file := Graph.ReadFile(filename)
	if !file {
		return
	}

	// Print file contents before simulation
	Graph.PrintFileContents(filename)

	// Storing the Room info on a struct with their Name as key to that value
	RoomMap := make(map[string]*functions.ARoom)
	for i := 0; i < len(Graph.Rooms); i++ {
		RoomMap[Graph.Rooms[i].Name] = &Graph.Rooms[i]
	}
	// Getting all the possible path using the BFS Algorithm
	Paths := functions.GetAllPaths(RoomMap, Graph.StartRoom.Name, Graph.EndRoom.Name)
	// Filtering the solution to get the most path possibles that doesn't use similar Room
	BestPath := functions.FilterPath(Paths, Graph.StartRoom.Name, Graph.EndRoom.Name)
	// Distribution of the Ants into the room they are going to use
	ant := functions.PlaceAnts(BestPath, Graph.AntCount)
	// Moving the Ants on the path and printing their movements
	functions.SimulateAntMovement(BestPath, ant)
}