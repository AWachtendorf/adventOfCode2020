package solve
//
//import (
//	"strconv"
//	"strings"
//)
//
//type CommandsList struct {
//	com []*command
//}
//
//type command struct {
//	direction string
//	amount    int
//}
//
//
//func (ss *Ship) moveShipForward(i int) {
//	switch ss.facing {
//	case dirNorth:
//		ss.moveShipNorth(i)
//	case dirEast:
//		ss.moveShipEast(i)
//	case dirSouth:
//		ss.moveShipSouth(i)
//	case dirWest:
//		ss.moveShipWest(i)
//	}
//}
//
//func (ss *Ship) moveShipNorth(i int) {
//	ss.South -= i
//	ss.North += i
//}
//func (ss *Ship) moveShipEast(i int) {
//	ss.West -= i
//	ss.East += i
//}
//func (ss *Ship) moveShipSouth(i int) {
//	ss.North -= i
//	ss.South += i
//}
//func (ss *Ship) moveShipWest(i int) {
//	ss.East -= i
//	ss.West += i
//}
//
//func (ss *Ship) changeDregreeAndDirection(s string, i int) {
//	if s == "L" {
//		ss.Degree -= i
//	}
//	if s == "R" {
//		ss.Degree += i
//	}
//	if ss.Degree == -90 {
//		ss.Degree = 270
//	}
//	if ss.Degree == -180 {
//		ss.Degree = 180
//	}
//
//	if ss.Degree == -270 {
//		ss.Degree = 90
//	}
//
//	if ss.Degree == 360 {
//		ss.Degree = 0
//	}
//	if ss.Degree == 450 {
//		ss.Degree = 90
//	}
//	if ss.Degree == 540 {
//		ss.Degree = 180
//	}
//	switch ss.Degree {
//	case 0:
//		ss.facing = dirNorth
//	case 90:
//		ss.facing = dirEast
//	case 180:
//		ss.facing = dirSouth
//	case 270:
//		ss.facing = dirWest
//	}
//
//}
//
//func (ss *Ship) readCommands(s string, i int) {
//	if s == "R" || s == "L" {
//		ss.changeDregreeAndDirection(s, i)
//	}
//	if s == "N" {
//		ss.moveShipNorth(i)
//	}
//	if s == "E" {
//		ss.moveShipEast(i)
//	}
//	if s == "S" {
//		ss.moveShipSouth(i)
//	}
//	if s == "W" {
//		ss.moveShipWest(i)
//	}
//	if s == "F" {
//		ss.moveShipForward(i)
//	}
//}
//
//type facingDirection int
//
//type Ship struct {
//	North, East, South, West, Degree int
//	facing                           facingDirection
//}
//
//const (
//	dirNorth facingDirection = iota
//	dirEast
//	dirSouth
//	dirWest
//)
//
//
//
//func (ss *Ship) calculate() int {
//	checkSum := 0
//
//	if ss.North > 0 {
//		checkSum += ss.North
//	}
//	if ss.East > 0 {
//		checkSum += ss.East
//
//	}
//	if ss.South > 0 {
//		checkSum += ss.South
//
//	}
//	if ss.West > 0 {
//		checkSum += ss.West
//
//	}
//
//	return checkSum
//}
//
//func (c *CommandsList) prepareInput(s string) *CommandsList {
//	splitString := strings.SplitAfter(s, "\n")
//	for _, ss := range splitString {
//		ss = strings.ReplaceAll(ss, "\n", "")
//		coms := []rune(ss)
//		comm := string(coms[0])
//		val, err := strconv.Atoi(string(coms[1:]))
//		if err != nil {
//			print(err)
//		}
//		cccc := &command{
//			direction: comm,
//			amount:    val,
//		}
//
//		c.com = append(c.com, cccc)
//
//	}
//
//	return c
//}
//
