package lemin

import "testing"

func TestGraphConnect(t *testing.T) {
	graph := &Graph{}
	graph.Connect("room1", "room2")

	// Check bidirectional connection
	if len(graph.Links["room1"]) == 0 || graph.Links["room1"][0] != "room2" {
		t.Errorf("room1 should be connected to room2")
	}
	if len(graph.Links["room2"]) == 0 || graph.Links["room2"][0] != "room1" {
		t.Errorf("room2 should be connected to room1")
	}
}

func TestGraphInitialization(t *testing.T) {
	graph := &Graph{}
	graph.Connect("room1", "room2")

	// Verify Links map is initialized
	if graph.Links == nil {
		t.Errorf("Links map should be initialized after Connect")
	}
}

func TestAntStruct(t *testing.T) {
	ant := Ant{
		Path:         []string{"start", "middle", "end"},
		CurrentRoom:  1,
		Previous:     "start",
		IsInactive:   false,
	}

	if len(ant.Path) != 3 {
		t.Errorf("Ant path should have 3 rooms")
	}
	if ant.CurrentRoom != 1 {
		t.Errorf("Incorrect current room")
	}
	if ant.IsInactive {
		t.Errorf("Ant should not be inactive")
	}
}

func TestGetAllPaths(t *testing.T) {
	// Create a test graph
	graph := Graph{
		Start: "A",
		End:   "D",
		Links: map[string][]string{
			"A": {"B", "C"},
			"B": {"A", "D"},
			"C": {"A", "D"},
			"D": {"B", "C"},
		},
	}

	// Call GetAllPaths
	Paths = [][]string{} // Reset global Paths
	graph.GetAllPaths(graph.Start)

	// Define the expected paths
	expectedPaths := [][]string{
		{"A", "B", "D"},
	}

	// Check the number of paths found
	if len(Paths) != len(expectedPaths) {
		t.Fatalf("Expected %d paths, but got %d", len(expectedPaths), len(Paths))
	}

	// Verify the content of the paths
	for _, expectedPath := range expectedPaths {
		found := false
		for _, actualPath := range Paths {
			if equalPaths(expectedPath, actualPath) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected path %v not found in actual paths %v", expectedPath, Paths)
		}
	}
}

// Helper function to compare two paths
func equalPaths(path1, path2 []string) bool {
	if len(path1) != len(path2) {
		return false
	}
	for i := range path1 {
		if path1[i] != path2[i] {
			return false
		}
	}
	return true
}