package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

//https://www.callicoder.com/golang-basic-types-operators-type-conversion/
func variableTypesSize() {
	var vUint8 byte
	var vSlice24Uint8 [24]uint8
	var vInt int
	var vSlice24Int [24]int
	var vMyCube cube
	var vMyCube2x2x2 cube2x2x2

	fmt.Printf("uint(byte): %T, %d\n", vUint8, unsafe.Sizeof(vUint8))
	fmt.Printf("Slice Uint8: %T, %d\n", vSlice24Uint8, unsafe.Sizeof(vSlice24Uint8))
	fmt.Printf("int: %T, %d\n", vInt, unsafe.Sizeof(vInt))
	fmt.Printf("Slice int: %T, %d\n", vSlice24Int, unsafe.Sizeof(vSlice24Int))
	fmt.Printf("My Cube: %T, %d\n", vMyCube, unsafe.Sizeof(vMyCube))
	fmt.Printf("My Cube2x2x2: %T, %d\n", vMyCube2x2x2, unsafe.Sizeof(vMyCube2x2x2))
}

func cubeTurnFaceByCube() {
	var c cube
	c = cube{
		2, 2, 2, 2,
		3, 3, 3, 3,
		4, 4, 4, 4,
		1, 1, 1, 1,
		0, 0, 0, 0,
		5, 5, 5, 5,
	}
	fmt.Println(c)
	c1 := c.turnFace(0, 0)
	fmt.Println(c)
	fmt.Println(c1)

}

func cubeTurnFaceByHelper() {
	var c cube
	c = cube{
		2, 2, 2, 2,
		3, 3, 3, 3,
		4, 4, 4, 4,
		1, 1, 1, 1,
		0, 0, 0, 0,
		5, 5, 5, 5,
	}
	fmt.Println(c)
	c1 := cube2x2x2Helper.turnFace(c, 0, 0)
	c2 := cube2x2x2HelperMod.turnFace(c, 0, 0)
	fmt.Println(c)
	fmt.Println(c1, "helper")
	fmt.Println(c2, "helpermod")
}

func cubeGetID() {
	var c cube
	c = cube{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5}
	fmt.Println(c.getIDCustom())
	fmt.Printf("%s", c)
}

func solvee(e *engine2x2x2) {
	loops := 0
	for len(e.ids) != loops {
		e.run(e.ids[loops], 16)
		loops++
	}
	fmt.Println("Permutations:", len(e.permutations))
	fmt.Println("Ids:", len(e.ids))
	fmt.Println("Loops:", loops)
}

func run(e *engine2x2x2, id string, maxLevel byte) {
	for face := 0; face < 3; face++ {
		for direction := range e.helper.Directions {
			child := e.helper.turnFace(e.permutations[id].cube, face, direction)
			childID := string(child[:])
			if _, ok := e.permutations[childID]; !ok {
				if maxLevel == 0 || e.permutations[id].level < maxLevel {
					e.ids = append(e.ids, childID)
					e.permutations[childID] = &permutation{
						cube:   child,
						level:  e.permutations[id].level + 1,
						parent: e.permutations[id],
					}
				}
			}
		}
	}
}

func (e *engine2x2x2) solve() {
	loops := 0
	for len(e.ids) != loops {
		e.run(e.ids[loops], 0)
		loops++
	}

	fmt.Println("Permutations:", len(e.permutations))
	fmt.Println("Ids:", len(e.ids))
	fmt.Println("Loops:", loops)
}

func (e *engine2x2x2) run(id string, maxLevel byte) {
	for face := 0; face < 3; face++ {
		for direction := range e.helper.Directions {
			if e.solution != nil {
				continue
			}
			child := e.helper.turnFace(e.permutations[id].cube, face, direction)
			childID := string(child[:])
			if _, ok := e.permutations[childID]; !ok {
				if maxLevel == 0 || e.permutations[id].level < maxLevel {
					combination := permutation{
						cube:   child,
						level:  e.permutations[id].level + 1,
						parent: e.permutations[id],
					}
					e.ids = append(e.ids, childID)
					e.permutations[childID] = &combination
					if e.helper.isSolved(child) {
						e.solution = &combination
					}
				}
			}
		}
	}
}

func (c *cube2x2x2Mod) isSolved(cc cube) bool {
	solved := [24]byte{
		49, 49, 49, 49, // Front | Red
		50, 50, 50, 50, // Right | Blue
		52, 52, 52, 52, // Back  | Orange
		53, 53, 53, 53, // Left  | Green
		48, 48, 48, 48, // Up   | White
		51, 51, 51, 51, // Back  | Yellow
	}
	return solved == cc
}

func (e *engine2x2x2) mermaid() {
	fmt.Println("graph TD")
	for id, permutation := range e.permutations {
		if permutation.parent != nil {
			fmt.Printf("%s(%d) --> %s(%d)\n", id, permutation.level, permutation.parent.cube, permutation.parent.level)
		} else {
			fmt.Printf("%s(%d)\n", permutation.cube, permutation.level)
		}
	}
}

func engineTest() {
	var solution *permutation
	c := cube{
		49, 49, 49, 49,
		50, 50, 50, 50,
		52, 52, 52, 52,
		53, 53, 53, 53,
		48, 48, 48, 48,
		51, 51, 51, 51,
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 1; i++ {
		c = cube2x2x2HelperMod.turnFace(c, r1.Intn(3), r1.Intn(2))
	}

	id := string(c[:])

	e := engine2x2x2{
		permutations: map[string]*permutation{},
		ids:          []string{},
		helper:       &cube2x2x2HelperMod,
	}
	e.permutations[id] = &permutation{
		cube:  c,
		level: 0,
	}
	solution = nil
	if cube2x2x2HelperMod.isSolved(c) {
		solution = e.permutations[id]
	}
	e.solution = solution
	e.ids = append(e.ids, id)

	e.solve()

	//e.mermaid()
	fmt.Println("Solved in level: ", e.solution.level)

}

func main() {
	printMemUsage()
	//variableTypesSize()

	//cubeTurnFaceByCube()
	//cubeTurnFaceByHelper()

	//cubeGetID()
	//printMemUsage()
	//generateSolutions()
	engineTest()
	printMemUsage()
}
