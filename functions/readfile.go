package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadFile reads the file, validates the contents and stores Room
func (RoomArray *Graph) ReadFile(filename string) bool {
	// Opening the file
	file, err := os.Open("./examples/" + filename)
	if err != nil {
		return false
	}
	defer file.Close()

	// Variables to check if the Room that are linked exist
	startingroom := false
	endingroom := false

	// Variables to check how many times we encounter ##start and ##end
	startcount := 0
	endcount := 0

	// Variable to check only the very first line (number of Ants)
	first := true

	// Going through the file line by line
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()

		// Handle the first line for Ants count
		if first {
			numberofants, err := strconv.Atoi(line)
			if err != nil || numberofants < 1 {
				fmt.Println("Error: Invalid number of Ants")
				return false
			} else {
				RoomArray.AntCount = numberofants
			}
			first = false
		}

		// Splitting room data and link data
		roomarray := strings.Split(line, " ")
		linkarray := strings.Split(line, "-")

		// Handle Room and links
		if len(roomarray) == 3 && startingroom {
			coordinate_x, err1 := strconv.Atoi(roomarray[1])
			coordinate_y, err2 := strconv.Atoi(roomarray[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Error: Invalid coordinates (No ints)")
				return false
			} else if roomarray[0][0] == '#' || roomarray[0][0] == 'L' {
				fmt.Println("Error: Invalid room name (L or #)")
				return false
			} else {
				// Storing the room values into a struct
				SingleRoom := ARoom{
					Name:        roomarray[0],
					XCoordinate: coordinate_x,
					YCoordinate: coordinate_y,
				}
				RoomArray.StartRoom = SingleRoom
				RoomArray.Rooms = append(RoomArray.Rooms, SingleRoom)
				startingroom = false
			}
		} else if len(roomarray) == 3 && endingroom {
			coordinate_x, err1 := strconv.Atoi(roomarray[1])
			coordinate_y, err2 := strconv.Atoi(roomarray[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Error: Invalid coordinates (No ints)")
				return false
			} else if roomarray[0][0] == '#' || roomarray[0][0] == 'L' {
				fmt.Println("Error: Invalid room name (L or #)")
				return false
			} else {
				// Storing the room values into a struct
				SingleRoom := ARoom{
					Name:        roomarray[0],
					XCoordinate: coordinate_x,
					YCoordinate: coordinate_y,
				}
				RoomArray.EndRoom = SingleRoom
				RoomArray.Rooms = append(RoomArray.Rooms, SingleRoom)
				endingroom = false
			}
		} else if len(roomarray) == 3 {
			coordinate_x, err1 := strconv.Atoi(roomarray[1])
			coordinate_y, err2 := strconv.Atoi(roomarray[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Error: Invalid coordinates (No ints)")
				return false
			} else if roomarray[0][0] == '#' || roomarray[0][0] == 'L' {
				fmt.Println("Error: Invalid room name (L or #)")
				return false
			} else {
				SingleRoom := ARoom{
					Name:        roomarray[0],
					XCoordinate: coordinate_x,
					YCoordinate: coordinate_y,
				}
				RoomArray.Rooms = append(RoomArray.Rooms, SingleRoom)
			}
		} else if len(linkarray) == 2 {
			// Checking if a room is linked to itself
			if linkarray[0] == linkarray[1] {
				fmt.Println("Error: Cannot link a room to itself")
				return false
			}

			exist1, exist2 := false, false
			index1, index2 := 0, 0
			// Iterating over all room names to check if the link is valid
			for i := 0; i < len(RoomArray.Rooms); i++ {
				if RoomArray.Rooms[i].Name == linkarray[0] {
					exist1 = true
					index1 = i
				} else if RoomArray.Rooms[i].Name == linkarray[1] {
					exist2 = true
					index2 = i
				}
			}

			if !exist1 || !exist2 {
				fmt.Println("Error: Invalid Room Name (Doesn't exist)")
				return false
			} else {
				RoomArray.Rooms[index1].Links = append(RoomArray.Rooms[index1].Links, linkarray[1])
				RoomArray.Rooms[index2].Links = append(RoomArray.Rooms[index2].Links, linkarray[0])
			}
		}

		// Checking for start and end Room
		if line == "##start" {
			startcount++
			startingroom = true
		} else if line == "##end" {
			endcount++
			endingroom = true
		}
	}

	// Check if there is exactly one start and one end
	if startcount != 1 {
		fmt.Println("Error: No starting point")
		return false
	} else if endcount != 1 {
		fmt.Println("Error: No ending point")
		return false
	}

	// Now check for uniqueness and coordinates in a separate function
	if !CheckRoomUniqueness(RoomArray) {
		return false
	}

	return true
}
