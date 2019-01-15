package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

// A - performs block redistribution until a loop is detected.
// the number of steps until the loop happens is returned.
func A() (int, error) {
	f, err := os.Open("inputA")
	if err != nil {
		return 0, err
	}

	r := bufio.NewReader(f)
	input, err := readRow(r)
	if err != nil && err != io.EOF {
		return 0, err
	}

	sum := 0
	history := make([][]int, 0)
	for {
		max := findMaxIndex(input)
		realocate(input, max)
		sum++
		if isALoop(input, history) {
			break
		}
		copyS := make([]int, len(input))
		copy(copyS, input)
		history = append(history, copyS)
	}
	return sum, nil
}

func findMaxIndex(in []int) int {
	max := math.MinInt32
	maxI := -1
	for i, v := range in {
		if v > max {
			max = v
			maxI = i
		}
	}
	return maxI
}

func realocate(in []int, maxI int) []int {
	pool := in[maxI]
	in[maxI] = 0
	cycles := pool / len(in)
	rest := pool % len(in)

	for i := range in {
		in[i] += cycles
		if (i > maxI && i <= maxI+rest) ||
			i <= maxI+rest-len(in) {
			in[i]++
		}
	}

	return in
}

func isALoop(in []int, history [][]int) bool {
	for c, sample := range history {
		broken := false
		for i, v := range in {
			if sample[i] != v {
				broken = true
				break
			}
		}
		if !broken {
			fmt.Println("c:", c)
			return true
		}
	}
	return false
}

// B returns the length of the loop
func B() (int, error) {
	f, err := os.Open("inputA")
	if err != nil {
		return 0, err
	}

	r := bufio.NewReader(f)
	input, err := readRow(r)
	if err != nil && err != io.EOF {
		return 0, err
	}

	sum := 0
	history := make([][]int, 0)
	for {
		max := findMaxIndex(input)
		realocate(input, max)
		sum++
		if isALoop(input, history) {
			fmt.Println("history len:", len(history))
			// answer is (historyLen - c)
			break
		}
		copyS := make([]int, len(input))
		copy(copyS, input)
		history = append(history, copyS)
	}
	return sum, nil
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
