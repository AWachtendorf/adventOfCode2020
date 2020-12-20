package solve

import (
	"strconv"
	"strings"
)


type CommandsList struct {
	Com []*command
}

type command struct {
	Direction string
	Amount    int
}


func (ss *Ship) moveShipForward(w *Waypoint,i int) {
	switch ss.facing {
	case dirNorth:
		ss.moveShipNorth(w,i)
	case dirEast:
		ss.moveShipEast(w,i)
	case dirSouth:
		ss.moveShipSouth(w,i)
	case dirWest:
		ss.moveShipWest(w,i)
	}
}

func (ss *Ship) moveShipNorth(w *Waypoint,i int) {
	ss.South -= i * w.North
	w.South = ss.South + w.South
	ss.North += i * w.North
	w.North += ss.North + w.North
}

func (ss *Ship) moveShipEast(w *Waypoint,i int) {
	ss.West -= i * w.West

}
func (ss *Ship) moveShipSouth(w *Waypoint,i int) {
	ss.North -= i
	ss.South += i
}
func (ss *Ship) moveShipWest(w *Waypoint,i int) {
	ss.East -= i
	ss.West += i
}

func (ss *Ship) changeDregreeAndDirection(s string, i int) {
	if s == "L" {
		ss.Degree -= i
	}
	if s == "R" {
		ss.Degree += i
	}
	if ss.Degree == -90 {
		ss.Degree = 270
	}
	if ss.Degree == -180 {
		ss.Degree = 180
	}

	if ss.Degree == -270 {
		ss.Degree = 90
	}

	if ss.Degree == 360 {
		ss.Degree = 0
	}
	if ss.Degree == 450 {
		ss.Degree = 90
	}
	if ss.Degree == 540 {
		ss.Degree = 180
	}
	switch ss.Degree {
	case 0:
		ss.facing = dirNorth
	case 90:
		ss.facing = dirEast
	case 180:
		ss.facing = dirSouth
	case 270:
		ss.facing = dirWest
	}

}

func (ss *Ship) ReadCommands(w*Waypoint,s string, i int) {
	if s == "R" || s == "L" {
		ss.changeDregreeAndDirection(s, i)
	}
	if s == "N" {
		ss.moveShipNorth(w,i)
	}
	if s == "E" {
		ss.moveShipEast(w,i)
	}
	if s == "S" {
		ss.moveShipSouth(w,i)
	}
	if s == "W" {
		ss.moveShipWest(w,i)
	}
	if s == "F" {
		ss.moveShipForward(w,i)
	}
}

type facingDirection int

type Ship struct {
	North, East, South, West, Degree int
	facing                           facingDirection
}

type Waypoint struct {
	North, East, South, West, Degree int

}

const (
	dirNorth facingDirection = iota
	dirEast
	dirSouth
	dirWest
)




func (c *CommandsList) PrepareInput(s string) *CommandsList {
	splitString := strings.SplitAfter(s, "\n")
	for _, ss := range splitString {
		ss = strings.ReplaceAll(ss, "\n", "")
		coms := []rune(ss)
		comm := string(coms[0])
		val, err := strconv.Atoi(string(coms[1:]))
		if err != nil {
			print(err)
		}
		cccc := &command{
			Direction: comm,
			Amount:    val,
		}

		c.Com = append(c.Com, cccc)

	}

	return c
}

