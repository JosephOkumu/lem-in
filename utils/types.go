package lemin

// Graph represents the ant farm with rooms and tunnels.
type Graph struct {
	Start, End string         
	Links      map[string][]string // Tunnels between rooms
}

// Ant represents an individual ant in the farm.
type Ant struct {
	Path        []string 
	CurrentRoom int      
	Previous    string   
	IsInactive  bool   
}

// Connect links two rooms with a bidirectional tunnel.
func (graph *Graph) Connect(room1, room2 string) {
	if graph.Links == nil {
		graph.Links = make(map[string][]string)
	}
	graph.Links[room1] = append(graph.Links[room1], room2)
	graph.Links[room2] = append(graph.Links[room2], room1)
}



// InputData stores the content of the input file.
var InputData string
