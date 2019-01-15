package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	a, err := A()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("A:", a)

	b, err := B()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("B:", b)
}

// A - returns checksum being the sum of differences between the highest
// and lowest vaules for each row
func A() (int, error) {
	f, err := os.Open("inputA")
	if err != nil {
		return 0, err
	}

	r := bufio.NewReader(f)
	sum := 0

	for {
		row, err := readRow(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		min := row[0]
		max := row[0]
		for _, n := range row {
			if n < min {
				min = n
			} else if n > max {
				max = n
			}
		}
		sum += max - min
	}
	return sum, nil
}

// B returns a sum of the only possible even division in each row
func B() (int, error) {
	f, err := os.Open("inputA")
	if err != nil {
		return 0, err
	}

	r := bufio.NewReader(f)
	sum := 0

	for {
		row, err := readRow(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		n := onlyEven(row)
		sum += n
	}
	return sum, nil
}

func onlyEven(slice []int) int {
	for i, m := range slice {
		for j, n := range slice {
			if i == j {
				continue
			}
			if m%n == 0 {
				return m / n
			}
		}
	}
	return 0
}

func readRow(r *bufio.Reader) ([]int, error) {
	s, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	s = strings.TrimRight(s, "\n\r")
	sSlice := strings.Split(s, "\t")
	numSlice := make([]int, len(sSlice))
	for i, numString := range sSlice {
		num, err := strconv.Atoi(numString)
		if err != nil {
			return nil, err
		}

		numSlice[i] = num
	}

	return numSlice, nil
}
