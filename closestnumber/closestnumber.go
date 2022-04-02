package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'closestNumbers' function below.
 *
 * The function accepts INTEGER_ARRAY numbers as parameter.
 */
func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func closestNumbers(numbers []int32) {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	var minDiff int32
	minDiff = numbers[1] - numbers[0]

	var i int32
	for i = 2; i <= int32(len(numbers)-1); i++ {
		minDiff = min(minDiff, numbers[i]-numbers[i-1])
	}

	for i = 1; i <= int32(len(numbers)-1); i++ {
		if numbers[i]-numbers[i-1] == minDiff {
			fmt.Println(numbers[i-1], numbers[i])
		}
	}
	// var i int32
	// for i = 1; i < int32(len(numbers)); i++ {
	// 	fmt.Println(numbers[i-1], nearestNumberBinarySearch(numbers, i, int32(len(numbers)-1), numbers[i-1]))
	// }
}

func nearestNumberBinarySearch(numbers []int32, start int32, end int32, current int32) int32 {
	mid := int32((start + end) / 2)
	if numbers[mid] == current {
		return numbers[mid]
	}
	if start == end-1 {
		if numbers[end]-current >= numbers[start]-current {
			return numbers[start]
		} else {
			return numbers[end]
		}
	}
	if numbers[mid] > current {
		return nearestNumberBinarySearch(numbers, start, mid, current)
	} else {
		return nearestNumberBinarySearch(numbers, mid, end, current)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	numbersCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var numbers []int32

	for i := 0; i < int(numbersCount); i++ {
		numbersItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		numbersItem := int32(numbersItemTemp)
		numbers = append(numbers, numbersItem)
	}

	closestNumbers(numbers)
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
