package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	hCube2x2x2 := createCube2x2x2Helper() // create 2x2x2 helper
	hColors := createColorsHelper()       // create colors helper
	c := cube2x2x2{
		0, 0, 0, 0,
		1, 1, 2, 2,
		3, 3, 4, 4,
		1, 1, 2, 2,
		3, 3, 4, 4,
		5, 5, 5, 5,
	} // Create solved cube

	for i := 0; i < 200; i++ { //scramble the Rubik's Cube
		c = hCube2x2x2.turnFace(c, r1.Intn(3), r1.Intn(2))
	}

	e := createEngine2x2x2(c, hCube2x2x2, hColors) //Create engine to solve

	fmt.Print("Scrambled cube:\n\n")
	e.printXterm(&c) // Print Rubik's Cube

	e.solve() //Solve the Rubik's Cube

	fmt.Println()
	fmt.Printf("Solved in %d steps:\n\n", e.isSolved.step)

	e.showSteps()                  // Print annotations to solve
	e.printXterm(&e.isSolved.cube) // Print solved Rubik's Cube
}
