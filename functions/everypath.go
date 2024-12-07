package functions

import (
	"container/list"
)

// ?Function that will get us all the possible paths with the BFS Algorithm
func GetAllPaths(Room map[string]*ARoom, start, end string) [][]string {
	var paths [][]string
	queue := list.New()
	queue.PushBack([]string{start})

	//While we have paths that didn't reach end yet, we continue
	for queue.Len() > 0 {

		//Removing the first element of the queue and storing it on the path variable
		path := queue.Remove(queue.Front()).([]string)

		//Storing the name of the last room of the path
		lastRoom := path[len(path)-1]

		//Checking if the room is currently at the end
		if lastRoom == end {
			//If that's the case we append this path and go to the next one
			paths = append(paths, path)
			continue
		}

		//We explore all the links of the current room
		for _, nextRoom := range Room[lastRoom].Links {
			if !contains(path, nextRoom) {
				//Creating a new path that will have the same length as the current path
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, nextRoom)
				queue.PushBack(newPath)
			}
		}
	}
	return paths
}


func contains(slice []string, element string) bool {
	for _, current := range slice {
		if current == element {
			return true
		}
	}
	
	return false
}
