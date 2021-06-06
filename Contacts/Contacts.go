package Contacts

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'contacts' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts 2D_STRING_ARRAY queries as parameter.
 */

func contacts(queries [][]string) []int32 {
	// Write your code here
	contacts := []string{}
	resArray := []int32{}
	for _, q := range queries {
		var res int32
		switch q[0] {
		case "add":
			contacts = append(contacts, q[1])
		case "find":
			for _, contact := range contacts {
				if strings.HasPrefix(contact, q[1]) {
					res++
				}
			}
			resArray = append(resArray, res)

		}
	}
	return resArray
}

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

	queriesRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries [][]string
	for i := 0; i < int(queriesRows); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []string
		for _, queriesRowItem := range queriesRowTemp {
			queriesRow = append(queriesRow, queriesRowItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := contacts(queries)

	for i, resultItem := range result {
		_, err := fmt.Fprintf(writer, "%d", resultItem)
		if err != nil {
			return
		}

		if i != len(result)-1 {
			_, err := fmt.Fprintf(writer, "\n")
			if err != nil {
				return
			}
		}
	}

	_, err = fmt.Fprintf(writer, "\n")
	if err != nil {
		return
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
