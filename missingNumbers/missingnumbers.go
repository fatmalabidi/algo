package missingnumbers

import (
	"errors"
)

func GetMissingNumbers(input []uint) []uint {
	min, max, _ := getMinMax(input)
	missing := []uint{}
	for i := min; i < max; i++ {
		if !exist(i, input) {
			missing = append(missing, i)
		}
	}
	return missing
}

func getMinMax(input []uint) (min uint, max uint, err error) {
	if input == nil || len(input) == 0 {
		err = errors.New("invalid input")
		return
	}
	max = input[0]
	min = input[0]
	for _, v := range input {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return
}
func exist(i uint, input []uint) bool {
	for _, v := range input {
		if v == i {
			return true
		}
	}
	return false
}

/*
Example of use
input := []uint{1, 0, 3, 6,10}
	missing := GetMissingNumbers(input)
	fmt.Println(missing
*/
