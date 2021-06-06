package FuzzBuzz

import (
"bufio"
"fmt"
"io"
"os"
"strconv"
"strings"
)



/*
 * Complete the 'fizzBuzz' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func fizzBuzz(input int) {

	for n:=1; n<=int(input);n++{

		// Write your code here
		if(n%3!=0 && n%5!=0){
			fmt.Println(n)
		}
		if(n%3!=0 && n%5==0){
			fmt.Println("Buzz")
		}
		if(n%3==0 && n%5!=0){
			fmt.Println("Fuzz")
		}

		if(n%3==0 && n%5==0){
			fmt.Println("FuzzBuzz")
		}

	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	fizzBuzz(int(n))

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

