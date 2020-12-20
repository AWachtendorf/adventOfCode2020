package main

import (
	. "./Puzzle"
	. "./Puzzle1"
	"sort"
)

func main() {
	v := &Values{}
	vv := &Values{}
	v.ReadAndConvert(PuzzleInput)
	vv.ReadAndConvert(PuzzleInput)
	e := &Existing{}
	e = v.SolveXMAS(25)

	solve := SolvePuzzle(e, v)

	j := vv.CheckRecursive(solve)
	sort.Ints(j)

	println(j[0] + j[len(j)-1])

}
