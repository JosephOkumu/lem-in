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
	// Create a mock graph structure
	graph := &Graph{
		Start: "A",
		End:   "D",
		Links: map[string][]string{
			"A": {"B", "C"},
			"B": {"D"},
			"C": {"D"},
			"D": {},
		},
	}

	// Expected paths
	expectedPaths := [][]string{
		{"A", "B", "D"},
		{"A", "C", "D"},
	}

	// Clear the global Paths slice before testing
	Paths = [][]string{}

	// Call the function
	graph.GetAllPaths(graph.Start)

	// Verify the number of paths
	if len(Paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(Paths))
	}

	// Verify the contents of each path
	for _, expectedPath := range expectedPaths {
		found := false
		for _, actualPath := range Paths {
			if equalSlices(expectedPath, actualPath) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected path %v not found in actual paths %v", expectedPath, Paths)
		}
	}
}

// Helper function to compare two slices for equality
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}