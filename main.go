package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isBalanced(s string) string {
	n := -1
	for len(s) != n {
		n = len(s)
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "{}", "")
		s = strings.ReplaceAll(s, "()", "")
	}
	if n == 0 {
		return "YES"
	}
	return "NO"
}

/*
func isBalanced(s string) string {

	fmt.Println("len(s) ", len(s))
	fmt.Println("s ", s)
	stringArray := strings.Split(s, "")
	if len(stringArray) == 0 {
		return "YES"
	}
	if len(stringArray)%2 != 0 {
		return "NO"
	}
	fmt.Println("len(stringArray) ", len(stringArray))
	fmt.Println("stringArray ", stringArray)

	sub1 := stringArray[:(len(stringArray) / 2)]
	sub2 := stringArray[(len(stringArray) / 2):]

	fmt.Println(sub1)
	fmt.Println(sub2)

	for i, s2 := range sub1 {

		switch s2 {
		case "{":
			if sub2[len(sub2)-i-1] != "}" {
				fmt.Println("i= ", i, " s2= ", s2, " compared to =", sub2[len(sub2)-i-1])

				return "NO"
			}
		case "[":
			if sub2[len(sub2)-i-1] != "]" {
				fmt.Println("i= ", i, " s2= ", s2, " compared to =", sub2[len(sub2)-i-1])

				return "NO"
			}
		case "(":
			if sub2[len(sub2)-i-1] != ")" {
				fmt.Println("i= ", i, " s2= ", s2, " compared to =", sub2[len(sub2)-i-1])
				return "NO"
			}
		default:
			return "NO"
		}
	}
	return "YES"
}*/

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer func(stdout *os.File) {
		err := stdout.Close()
		if err != nil {

		}
	}(stdout)

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)
		fmt.Println("\n _____________________ \n")
		result := isBalanced(s)
		_, err := fmt.Fprintf(writer, "%s\n", result)
		if err != nil {
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		return
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
