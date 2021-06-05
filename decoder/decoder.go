package decoder

import (
	"fmt"
	"strconv"
	"strings"
)

func Decode(input string) string {
	//validate input
	if input == "" {
		return ""
	}
	// decode
	stringArray := strings.Split(input, "")
	var resBuilder strings.Builder
	for i, v := range stringArray {
		counter, _ := strconv.Atoi(v)
		for j := 0; j < counter; j++ {
			resBuilder.Write([]byte(stringArray[i+1]))
		}
	}
	return resBuilder.String()
}
func Encode(input []string) string {
	//validate input
	if input == nil || len(input) == 0 {
		return ""
	}

	res := ""
	counter := 1
	prevChar := input[0]
	// encode input
	for i := 1; i < len(input); i++ {
		if input[i] == prevChar {
			counter++
			continue
		}
		//compress
		if counter > 1 {
			res = fmt.Sprintf("%v%v%v", res, counter, prevChar)
		} else {
			res = fmt.Sprintf("%v%v", res, prevChar)
		}
		prevChar = input[i]
		counter = 1
	}
	//treat the last string
	res = fmt.Sprintf("%v%v%v", res, counter, prevChar)
	return res
}
