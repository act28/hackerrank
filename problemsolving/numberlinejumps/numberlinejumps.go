package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// If x2 is further than x1 and jumps further than v1, they will never catch up.
	// Or, if both are moving at the same speed, they will never catch up, coz
	// x2 will always be ahead by (x2-x1).
	if x2 > x1 && v2 >= v1 {
		return "NO"
	}
	// We just need to solve `x1 + (y * v1) = x2 + (y * v2)` where `y` is the
	// number of jumps...
	// i.e.`y = (x2-x1)/(v1-v2)`
	if (x1-x2)%(v1-v2) == 0 {
		// i.e.`y` is a whole number.
		return "YES"
	}
	// ...otherwise, no.
	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	x1Temp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	x1 := int32(x1Temp)

	v1Temp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	v1 := int32(v1Temp)

	x2Temp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	x2 := int32(x2Temp)

	v2Temp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
	checkError(err)
	v2 := int32(v2Temp)

	result := kangaroo(x1, v1, x2, v2)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
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
