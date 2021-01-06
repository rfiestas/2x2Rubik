package main

import "fmt"

type helpers struct {
	cube2x2x2 *cube2x2x2Helper
	colors    *colorsHelper
}

type engine2x2x2 struct {
	permutations map[string]*permutation // Contains all permutations
	ids          []string                // Slice with the ids
	isSolved     *permutation            // Point to solved permutation or nil
	solution     *cube2x2x2              // Solution
	helpers      helpers                 // cube2x2x2 and color helpers
}

type permutation struct {
	cube      cube2x2x2
	step      byte
	parent    *permutation
	face      byte
	direction byte
}

func createEngine2x2x2(c cube2x2x2, hCube2x2x2 *cube2x2x2Helper, hColors *colorsHelper) *engine2x2x2 {
	e := engine2x2x2{
		permutations: map[string]*permutation{},
		ids:          []string{},
		helpers: helpers{
			cube2x2x2: hCube2x2x2,
			colors:    hColors,
		},
	}
	e.createFirstPermutation(&c)

	return &e
}

func (e *engine2x2x2) generateSolutionPermutation(id string) *cube2x2x2 {
	back := e.permutations[id].cube[19]
	left := e.permutations[id].cube[12]
	down := e.permutations[id].cube[22]
	front := e.helpers.colors.getInverseColor(back)
	right := e.helpers.colors.getInverseColor(left)
	up := e.helpers.colors.getInverseColor(down)

	return &cube2x2x2{
		up, up, up, up,
		left, left, front, front,
		right, right, back, back,
		left, left, front, front,
		right, right, back, back,
		down, down, down, down,
	}
}

func (e *engine2x2x2) printXterm(c *cube2x2x2) {
	format := `
        %s %s
           T
        %s %s
      
%s %s %s %s %s %s %s %s
   L       F       R       B
%s %s %s %s %s %s %s %s

        %s %s
           D
        %s %s

`
	values := e.helpers.colors.printable(c[:], "▐█▌")
	fmt.Printf(format, values...)
}

func (e *engine2x2x2) createFirstPermutation(c *cube2x2x2) {
	id := string(c[:])                 // Set permutation ID
	e.ids = append(e.ids, id)          // Append the ID
	e.permutations[id] = &permutation{ // Create first permutation
		cube:   *c,
		step:   0,
		parent: nil,
	}
	e.solution = e.generateSolutionPermutation(id) // Generate solution permutation
	e.isSolved = nil                               // Set isSolved to nil
	if *c == *e.solution {                         // Check if is solved
		e.isSolved = e.permutations[id] // Set isSolved
	}
}

func (e *engine2x2x2) addPermutation(c *cube2x2x2, cID string, step byte, parent *permutation, face int, direction int) {
	permutation := permutation{
		cube:      *c,
		step:      step,
		parent:    parent,
		face:      byte(face),
		direction: byte(direction),
	}
	e.ids = append(e.ids, cID)
	e.permutations[cID] = &permutation
	if *c == *e.solution {
		e.isSolved = &permutation
	}
}

func (e *engine2x2x2) solve() {
	loops := 0
	for len(e.ids) != loops {
		e.run(e.ids[loops])
		loops++
	}
}

func (e *engine2x2x2) run(id string) {
	for face := 0; face < 3; face++ {
		for direction := range e.helpers.cube2x2x2.Directions {
			if e.isSolved != nil {
				continue
			}
			child := e.helpers.cube2x2x2.turnFace(e.permutations[id].cube, face, direction)
			childID := string(child[:])
			if _, ok := e.permutations[childID]; !ok {
				e.addPermutation(&child, childID, e.permutations[id].step+1, e.permutations[id], face, direction)
			}
		}
	}
}

func (e *engine2x2x2) mermaid() {
	fmt.Println("graph TD")
	for id, permutation := range e.permutations {
		if permutation.parent != nil {
			fmt.Printf("%s(%d) --> %s(%d)\n", id, permutation.step, permutation.parent.cube, permutation.parent.step)
		} else {
			fmt.Printf("%s(%d)\n", permutation.cube, permutation.step)
		}
	}
}

func (e *engine2x2x2) showSteps() {
	var steps string
	if e.isSolved != nil {
		steps = e.recursiveStep(e.isSolved)
	}
	fmt.Println(steps)
}

func (e *engine2x2x2) recursiveStep(p *permutation) string {
	var s string
	if p.parent != nil {
		s = fmt.Sprintf("%s%s ", string(e.helpers.cube2x2x2.Notations.Faces[p.face]), string(e.helpers.cube2x2x2.Notations.Directions[p.direction]))
		pS := e.recursiveStep(p.parent)
		s = pS + s
	}
	return s
}
