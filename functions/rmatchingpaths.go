package functions

import "fmt"

// RemoveSimilarPaths checks for duplicate room names and coordinates
func (RoomArray *RoomStruct) RemoveSimilarPaths() bool {
	// Checking for room name uniqueness
	for i := 0; i < len(RoomArray.AllRooms); i++ {
		for j := i + 1; j < len(RoomArray.AllRooms); j++ {
			if RoomArray.AllRooms[i].Name == RoomArray.AllRooms[j].Name {
				fmt.Println("------------------------------")
				fmt.Println(" Error: Rooms are not unique")
				fmt.Println("------------------------------")
				return false
			}
		}
	}

	// Checking for duplicate room coordinates
	for i := 0; i < len(RoomArray.AllRooms); i++ {
		for j := i + 1; j < len(RoomArray.AllRooms); j++ {
			if RoomArray.AllRooms[i].X_value == RoomArray.AllRooms[j].X_value && RoomArray.AllRooms[i].Y_value == RoomArray.AllRooms[j].Y_value {
				fmt.Println("-------------------------------------------")
				fmt.Println("Error: Multiple rooms with same coordinates")
				fmt.Println("-------------------------------------------")
				return false
			}
		}
	}

	// Check if the links are valid and if a room is linked to itself
	for i := 0; i < len(RoomArray.AllRooms); i++ {
		for j := 0; j < len(RoomArray.AllRooms[i].Links); j++ {
			if RoomArray.AllRooms[i].Links[j] == RoomArray.AllRooms[i].Name {
				fmt.Println("------------------------------------")
				fmt.Println("Error: Cannot link a room to itself")
				fmt.Println("------------------------------------")
				return false
			}
		}
	}

	return true
}
