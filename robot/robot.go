package robot

import (
	"fmt"
	"strings"
)

func IsRobotMoving(input string) (bool, string) {
	stringArray := strings.Split(strings.ToLower(input), "")
	rmv := 0
	lmv := 0
	umv := 0
	dmv := 0
	for _, v := range stringArray {
		switch v {
		case "l":
			lmv++
		case "r":
			rmv++
		case "d":
			dmv++
		case "u":
			umv++
		}
	}
	moves := lmv != rmv || umv != dmv
	var resBuilder strings.Builder
	if moves {
		resBuilder.Write([]byte("the robot moves:"))
	}
	if dmv-umv > 0 {
		resBuilder.Write([]byte(fmt.Sprintf("\ndown by=%v", dmv-umv)))
	} else if dmv-umv < 0 {
		resBuilder.Write([]byte(fmt.Sprintf("\nup by=%v", dmv-umv)))
	}
	if lmv-rmv > 0 {
		resBuilder.Write([]byte(fmt.Sprintf("\nleft by=%v", lmv-rmv)))
	} else if lmv-rmv < 0 {
		resBuilder.Write([]byte(fmt.Sprintf("\nright by=%v", lmv-rmv)))
	}
	return moves, resBuilder.String()
}
