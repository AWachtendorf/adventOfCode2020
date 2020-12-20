package Puzzle1

import (
	"strconv"
	"strings"
)

type Values struct {
	Value []int
}

type Existing struct {
	Value []int
}

func (v *Values) ReadAndConvert(s string) *Values {
	cleanString := strings.ReplaceAll(s, "\n", " ")
	splitString := strings.Fields(cleanString)
	for _, ss := range splitString {
		value, err := strconv.Atoi(ss)
		if err != nil {
			println(err)
		}
		v.Value = append(v.Value, value)
	}
	return v
}

func (v *Values) SolveXMAS(ii int) *Existing {
	e := &Existing{}
	temp := false
	for i, j := range v.Value {
		max := i
		min := i - (ii + 1)
		if i >= ii {
			temp = false
			for iii, jj := range v.Value {
				for iiii, jjj := range v.Value {
					if j == jj+jjj && iii < max && iii > min && iiii < max && iiii > min {
						if !temp {
							e.Value = append(e.Value, j)
							temp = true
						}
					}
				}
			}

		}
	}
	return e
}

func SolvePuzzle(e *Existing, v *Values) int {
	for _, j := range e.Value {
		for i, jj := range v.Value {
			if j == jj {
				v.Value = append(v.Value[:i], v.Value[i+1:]...)
			}
		}
	}
	return v.Value[len(v.Value)-1]
}

func (v *Values) CheckRecursive(solve int) []int {
var temp []int
	sum := 0
Loop:
	for {
		for i, j := range v.Value {
			sum += j
			if sum == solve{

				println(sum)

				for _,jj := range v.Value{
				temp = append(temp,jj)
					if jj == j{
					return temp
				}

				}


			}
			if i == len(v.Value)-1{
				sum = 0
				v.Value = append(v.Value[:0], v.Value[0+1:]...)
				continue Loop
			}
			}

		break

	}
	return nil
}
