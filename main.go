package main

import (
	"fmt"
	. "github.com/loslem/adventOfCode2020/adv1"
	. "github.com/loslem/adventOfCode2020/adv12/program/puzzleInput"
	. "github.com/loslem/adventOfCode2020/adv12/program/solve"
	. "github.com/loslem/adventOfCode2020/adv2"
	. "github.com/loslem/adventOfCode2020/adv9/Puzzle"
	. "github.com/loslem/adventOfCode2020/adv9/Puzzle1"
	"sort"
)

func main(){

	v := NewValues()
	v.AddValues()
	v.CompareValues()

	d := &DataRepository{} //create repo
	d.ReadAllData()        //fill with data
	d.ProcessCheckTask1()  //process for task 1
	d.ProcessCheckTask2()  //process for task 2

	b := &Values{}
	bb := &Values{}
	b.ReadAndConvert(PuzzleInput)
	bb.ReadAndConvert(PuzzleInput)
	e := &Existing{}
	e = bb.SolveXMAS(25)

	solve := SolvePuzzle(e, b)

	j := bb.CheckRecursive(solve)
	sort.Ints(j)

	println(j[0] + j[len(j)-1])


	cc := &CommandsList{}

	cc.PrepareInput(PuzzleInputAdv12)
	s := &Ship{}
	w := &Waypoint{
		North:  1,
		East:   10,
		South:  0,
		West:   0,
		Degree: 0,
	}

	for _, j := range cc.Com {
		s.ReadCommands(w,j.Direction, j.Amount)
		println(fmt.Sprintf("SHIP: N: %v, E: %v, S: %v, W: %v", s.North, s.East, s.South, s.West))
		println(fmt.Sprintf("WAYPOINT: N: %v, E: %v, S: %v, W: %v", w.North, w.East, w.South, w.West))
	}


}
