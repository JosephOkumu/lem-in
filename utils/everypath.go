package lemin

import "container/list"

var Paths = [][]string{}

// GetAllPaths implements BFS to find all paths from start to end rooms
func (g *Graph) GetAllPaths(room string) {
	// BFS setup
	queue := list.New()              // Queue for BFS
	queue.PushBack([]string{room})   // Start with the start room
	visited := make(map[string]bool) // Visited rooms map
	visited[room] = true

	// Perform BFS
	for queue.Len() > 0 {
		// Get the first path from the queue
		currPath := queue.Remove(queue.Front()).([]string)
		lastRoom := currPath[len(currPath)-1] // Last room in the current path

		// If we've reached the end room, store the path
		if lastRoom == g.End {
			// Copy path and add to Paths
			pathCopy := append([]string(nil), currPath...)
			Paths = append(Paths, pathCopy)
		}

		// Explore neighbors (rooms connected to the current room)
		for _, neighbor := range g.Links[lastRoom] {
			if !visited[neighbor] {
				// Mark the room as visited and add the new path to the queue
				visited[neighbor] = true
				newPath := append(currPath, neighbor)
				queue.PushBack(newPath)
			}
		}
	}
}
