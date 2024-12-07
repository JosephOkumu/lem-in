package main

import (
	"fmt"
	"lemin/functions"
	"os"
)

func main() {
	//Checking if we have the good lenght of argument
	if len(os.Args) != 2 {
		fmt.Println("Invalid Input, Usage: go run . [filename]")
		return
	}

	//Storing the filename into a variable
	filename := os.Args[1]
	Graph := functions.Graph{}

	//Checking if the file is valid or not
	file := Graph.ReadFile(filename)
	if !file {
		fmt.Println("Error: problem reading file")
		return
	}

	//Storing the Room info on a struct with their Name as key to that value
	RoomMap := make(map[string]*functions.ARoom)
	for i := 0; i < len(Graph.Rooms); i++ {
		RoomMap[Graph.Rooms[i].Name] = &Graph.Rooms[i]
	}

	//Getting all the possible path using the BFS Algorithm
	Paths := functions.GetAllPaths(RoomMap, Graph.StartRoom.Name, Graph.EndRoom.Name)

	//Prints all the possible paths
	// fmt.Println(Paths)

	//Filtering the solution to get the most path possibles that doesn't use similar Room
	BestPath := functions.FilterPath(Paths, Graph.StartRoom.Name, Graph.EndRoom.Name)

	// fmt.Println("Best paths: ", BestPath)

	//Distribution of the Ants into the room they are going to use
	ant := functions.PlaceAnts(BestPath, Graph.AntCount)

	//Moving the Ants on the path and printing their movements
	functions.SimulateAntMovement(BestPath, ant)
	
}
