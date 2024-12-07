package functions

import (
	"fmt"
)

// CheckRoomUniquenessAndCoordinates checks if room names are unique and if rooms have unique coordinates
func CheckRoomUniquenessAndCoordinates(RoomArray *RoomStruct) bool {
	// Checking if room names are unique
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

	// Checking for rooms with the same coordinates
	for i := 0; i < len(RoomArray.AllRooms); i++ {
		for j := i + 1; j < len(RoomArray.AllRooms); j++ {
			if RoomArray.AllRooms[i].X_value == RoomArray.AllRooms[j].X_value && RoomArray.AllRooms[i].Y_value == RoomArray.AllRooms[j].Y_value {
				fmt.Println("Error: Multiple rooms with the same coordinates")
				return false
			}
		}
	}

	// If no issues are found, return true
	return true
}
