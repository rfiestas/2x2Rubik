package main

type face [12]byte

type direction [12]byte

type notations struct {
	Faces      [6]byte
	Directions [2]byte
}

//Cube2x2x2Helper is a rubik's cube
type cube2x2x2Helper struct {
	Faces      [6]face         //Front, Right, Back, Left, Up, Down
	Directions [2][6]direction //clockside, anticlockside
	Notations  notations       // Notation names
}

func (c *cube2x2x2Helper) turnFace(cc cube2x2x2, face int, direction int) cube2x2x2 {
	turnedFace := cc
	for p, position := range c.Directions[direction][face] {
		turnedFace[c.Faces[face][p]] = cc[position]
	}
	return turnedFace
}

func createCube2x2x2Helper() *cube2x2x2Helper {
	calculatedFaceDirections := [2][6]direction{}

	directions := [2]direction{
		0: {1, 3, 0, 2, 6, 7, 8, 9, 10, 11, 4, 5}, // turn clockside
		1: {2, 0, 3, 1, 10, 11, 4, 5, 6, 7, 8, 9}, // turn anticlockside
	}

	faces := [6]face{
		0: {0, 1, 2, 3, 11, 10, 9, 8, 7, 6, 5, 4},           //Up
		1: {6, 7, 14, 15, 2, 3, 8, 16, 21, 20, 13, 5},       //Front
		2: {8, 9, 16, 17, 3, 1, 10, 18, 23, 21, 15, 7},      //Right
		3: {20, 21, 22, 23, 14, 15, 16, 17, 18, 19, 12, 13}, //Down
		4: {10, 11, 18, 19, 1, 0, 4, 12, 22, 23, 17, 9},     //Back
		5: {4, 5, 12, 13, 0, 2, 6, 14, 20, 22, 19, 11},      //Left
	}

	for faceKey, FaceValues := range faces {
		for directionKey, directionValues := range directions {
			for p, value := range FaceValues {
				calculatedFaceDirections[directionKey][faceKey][directionValues[p]] = value
			}
		}
	}

	notations := notations{
		Faces:      [6]byte{'U', 'F', 'R', 'D', 'B', 'L'},
		Directions: [2]byte{' ', '\''},
	}

	return &cube2x2x2Helper{
		Faces:      faces,
		Directions: calculatedFaceDirections,
		Notations:  notations,
	}
}
