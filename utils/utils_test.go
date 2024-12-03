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