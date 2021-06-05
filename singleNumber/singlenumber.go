package singlenumber

import (
 	"errors"
)

func GetLowerCount(input []int) (int, error) {
	for _, v := range input {
		count, _ := getCount(v, input)
		if count == 1 {
			return v, nil
		}
	}
	return 0, errors.New("no single numbers")
}
func getCount(i int, input []int) (int, error) {
	if input == nil || len(input) == 0 {
		return 0, errors.New("invalid input")
	}
	counter := 0
	for _, v := range input {
		if v == i {
			counter++
		}
	}
	return counter, nil
}
/*
Example of use

input := "llrrddulrr"
moves, msg := robot.IsRobotMoving(input)
if moves {
fmt.Println(msg)
} else {
fmt.Println("the robot didn't move")
}*/