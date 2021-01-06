package main

import (
	"sort"
	"strconv"
)

type cube [24]byte

type permutation struct {
	cube   cube
	level  byte
	parent *permutation
}

func (c *cube) turnFace(face int, direction int) cube {
	turnedFace := *c

	f := cube2x2x2Helper.Faces[face]
	for p, position := range f {
		turnedFace[f[cube2x2x2Helper.Directions[direction][p]]] = c[position]
	}
	return turnedFace
}

func (c *cube) getIDString() string {
	return string(c[:])
}

func (c *cube) getIDCustom() string {
	var faceList [4]int = [4]int{}
	var allFaces []int = []int{}
	var mod byte = 0
	var res string

	for _, v := range c {
		faceList[mod] = int(v) + 1
		mod++
		if mod == 4 {
			mod = 0
			sort.Ints(faceList[:])
			faceHeight := (faceList[0] * 1000) + (faceList[1] * 100) + (faceList[2] * 10) + faceList[3]
			allFaces = append(allFaces, faceHeight)
		}
	}
	sort.Ints(allFaces)
	for _, v := range allFaces {
		res += strconv.Itoa(v)
	}
	return res
}

// Helper

type cube2x2x2 struct {
	FacesMod   map[int]face
	Faces      [6]face      //Front, Right, Back, Left, Up, Down
	Directions [2]direction //clockside, anticlockside
}
type face [12]int

type direction [12]int

type faceByte [12]byte

type directionByte [12]byte

func calculateFaceDirections() cube2x2x2Mod {
	calculatedFaceDirections := [2][6]directionByte{}
	d := [2]directionByte{
		0: {1, 3, 0, 2, 6, 7, 8, 9, 10, 11, 4, 5}, // turn clockside
		1: {2, 0, 3, 1, 10, 11, 4, 5, 6, 7, 8, 9}, // turn anticlockside
	}
	f := [6]faceByte{
		0: {0, 1, 2, 3, 16, 17, 4, 6, 21, 20, 15, 13},    //Front
		1: {4, 5, 6, 7, 17, 19, 8, 10, 23, 21, 3, 1},     //Right
		4: {8, 9, 10, 11, 19, 18, 12, 14, 22, 23, 7, 5},  //Back
		3: {12, 13, 14, 15, 18, 16, 0, 2, 20, 22, 11, 9}, //Left
		2: {18, 19, 16, 17, 9, 8, 5, 4, 1, 0, 13, 12},    //
		5: {20, 21, 22, 23, 2, 3, 6, 7, 10, 11, 14, 15},  //Down
	}
	for faceKey, FaceValues := range f {
		for directionKey, directionValues := range d {
			for p, position := range FaceValues {
				//fmt.Println(faceKey, FaceValues, "Face")
				//fmt.Println(directionKey, directionValues, "Direction")
				//fmt.Println(p, position)
				calculatedFaceDirections[directionKey][faceKey][directionValues[p]] = position
				//fmt.Println(directionValues[0])
			}
			//fmt.Println(calculatedFaceDirections[directionKey][faceKey])
		}
	}

	return cube2x2x2Mod{
		FacesMod: map[int]face{
			0: {0, 1, 2, 3, 16, 17, 4, 6, 21, 20, 15, 13}, //Front
			1: {4, 5, 6, 7, 17, 19, 8, 10, 23, 21, 3, 1},  //Right
			2: {18, 19, 16, 17, 9, 8, 5, 4, 1, 0, 13, 12}, //Up
		},
		Faces:      f,
		Directions: calculatedFaceDirections,
	}
}

type cube2x2x2Mod struct {
	FacesMod   map[int]face
	Faces      [6]faceByte         //Front, Right, Back, Left, Up, Down
	Directions [2][6]directionByte //clockside, anticlockside
}

func (c *cube2x2x2Mod) turnFace(cc cube, face int, direction int) cube {
	turnedFace := cc
	for p, position := range c.Directions[direction][face] {
		turnedFace[c.Faces[face][p]] = cc[position]
	}
	return turnedFace
}

func (c *cube2x2x2Mod) getIDString(cube cube) string {
	return string(cube[:])
}

var cube2x2x2HelperMod = calculateFaceDirections()

var cube2x2x2Helper = cube2x2x2{
	FacesMod: map[int]face{
		0: {0, 1, 2, 3, 16, 17, 4, 6, 21, 20, 15, 13}, //Front
		1: {4, 5, 6, 7, 17, 19, 8, 10, 23, 21, 3, 1},  //Right
		2: {18, 19, 16, 17, 9, 8, 5, 4, 1, 0, 13, 12}, //Up
	},
	Faces: [6]face{
		0: {0, 1, 2, 3, 16, 17, 4, 6, 21, 20, 15, 13},    //Front
		1: {4, 5, 6, 7, 17, 19, 8, 10, 23, 21, 3, 1},     //Right
		2: {18, 19, 16, 17, 9, 8, 5, 4, 1, 0, 13, 12},    //Up
		3: {12, 13, 14, 15, 18, 16, 0, 2, 20, 22, 11, 9}, //Left
		4: {8, 9, 10, 11, 19, 18, 12, 14, 22, 23, 7, 5},  //Back
		5: {20, 21, 22, 23, 2, 3, 6, 7, 10, 11, 14, 15},  //Down
	},
	Directions: [2]direction{
		0: {1, 3, 0, 2, 6, 7, 8, 9, 10, 11, 4, 5}, // turn clockside
		1: {2, 0, 3, 1, 10, 11, 4, 5, 6, 7, 8, 9}, // turn anticlockside
	},
}

func (c *cube2x2x2) turnFace(cc cube, face byte, direction byte) cube {
	turnedFace := cc

	f := c.Faces[face]
	for p, position := range f {
		turnedFace[f[c.Directions[direction][p]]] = cc[position]
	}
	return turnedFace
}

func (c *cube2x2x2) getIDString(cube cube) string {
	return string(cube[:])
}

func (c *cube2x2x2) getIDCustom(cube *cube) string {
	var faceList [4]int = [4]int{}
	var allFaces []int = []int{}
	var mod byte = 0
	var res string

	for _, v := range cube {
		faceList[mod] = int(v) + 1
		mod++
		if mod == 4 {
			mod = 0
			sort.Ints(faceList[:])
			faceHeight := (faceList[0] * 1000) + (faceList[1] * 100) + (faceList[2] * 10) + faceList[3]
			allFaces = append(allFaces, faceHeight)
		}
	}
	sort.Ints(allFaces)
	for _, v := range allFaces {
		res += strconv.Itoa(v)
	}
	return res
}

type cube2x2x2ModByte struct {
	FacesMod   map[int]face
	Faces      [6]faceByte         //Front, Right, Back, Left, Up, Down
	Directions [2][6]directionByte //clockside, anticlockside
}

func (c *cube2x2x2ModByte) turnFace(cc cube, face int, direction int) cube {
	turnedFace := cc
	for p, position := range c.Directions[direction][face] {
		turnedFace[c.Faces[face][p]] = cc[position]
	}
	return turnedFace
}

func (c *cube2x2x2ModByte) getIDString(cube cube) string {
	return string(cube[:])
}

var cube2x2x2HelperModByte = calculateFaceDirections()

type engine2x2x2 struct {
	permutations map[string]*permutation
	ids          []string
	solution     *permutation
	helper       *cube2x2x2Mod
}
