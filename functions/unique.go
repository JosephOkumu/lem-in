package functions

import (
	"fmt"
)

// CheckRoomUniquenessAndCoordinates checks if room names are unique and if Room have unique coordinates
func CheckRoomUniqueness(RoomArray *Graph) bool {
	// Checking if room names are unique
	for i := 0; i < len(RoomArray.Rooms); i++ {
		for j := i + 1; j < len(RoomArray.Rooms); j++ {
			if RoomArray.Rooms[i].Name == RoomArray.Rooms[j].Name {
				fmt.Println("------------------------------")
				fmt.Println(" Error: Room are not unique")
				fmt.Println("------------------------------")
				return false
			}
		}
	}

	// Checking for Room with the same coordinates
	for i := 0; i < len(RoomArray.Rooms); i++ {
		for j := i + 1; j < len(RoomArray.Rooms); j++ {
			if RoomArray.Rooms[i].XCoordinate == RoomArray.Rooms[j].XCoordinate && RoomArray.Rooms[i].YCoordinate == RoomArray.Rooms[j].YCoordinate {
				fmt.Println("Error: Multiple Room with the same coordinates")
				return false
			}
		}
	}

	// If no issues are found, return true
	return true
}
