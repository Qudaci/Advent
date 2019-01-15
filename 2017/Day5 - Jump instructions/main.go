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

// A - returns numbers of steps needed to escape the input jump maze
func A() (int, error) {
	f, err := os.Open("inputA")
	if err != nil {
		return 0, err
	}

	r := bufio.NewReader(f)
	in, err := readInput(r)
	if err != nil {
		return 0, err
	}
	i := 0
	steps := 0

	for {
		if i >= len(in) || i < 0 {
			break
		}

		in[i]++
		i += in[i] - 1

		steps++
	}
	return steps, nil
}

// B - returns numbers of steps needed to escape the input jump maze
func B() (int, error) {
	f, err := os.Open("inputA")
	if err != nil {
		return 0, err
	}

	r := bufio.NewReader(f)
	in, err := readInput(r)
	if err != nil {
		return 0, err
	}
	i := 0
	steps := 0

	for {
		if i >= len(in) || i < 0 {
			break
		}

		newI := i + in[i]
		if in[i] >= 3 {
			in[i]--
		} else {
			in[i]++
		}
		i = newI
		steps++
	}
	return steps, nil
}

func readInput(r *bufio.Reader) ([]int, error) {
	s, err := r.ReadString(0)
	if err != nil && err != io.EOF {
		return nil, err
	}
	sSlice := strings.Split(s, "\r\n")
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
