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
func (g *Graph) Connect(room1, room2 string) {
	if g.Links == nil {
		g.Links = make(map[string][]string)
	}
	g.Links[room1] = append(g.Links[room1], room2)
	g.Links[room2] = append(g.Links[room2], room1)
}



// InputData stores the content of the input file.
var InputData string
