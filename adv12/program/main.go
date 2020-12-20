package adv12

import (
	""
	"fmt"
)

func main() {
	cc := &sv.CommandsList{}

	cc.PrepareInput(pz.Test)
	s := &sv.Ship{}
	w := &sv.Waypoint{
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