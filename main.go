package main

import (
	. "github.com/loslem/adventOfCode2020/adv1"
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

}
