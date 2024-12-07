package functions

//Struct that will store all informations needed
type Graph struct {
	AntCount         int
	AntNames   []string
	Rooms      []ARoom
	AllPaths      [][]string
	StartRoom  ARoom
	EndRoom   ARoom
}

//Struct that will store a single room informations
type ARoom struct {
	Name    string
	XCoordinate int
	YCoordinate int
	Links   []string
	Visited bool
}