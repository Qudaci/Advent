package main

import (
	"fmt"
)

func main() {
	inputA := 277678
	inputB := 277678

	fmt.Println("A:", A(inputA))
	fmt.Println("B:", B(inputB))
}

// A - calculates the shortest path to the dest memory square from square 1
// in a spiral square structure starting at one, going right, and spiraling
// counterclockwise
func A(dest int) int {
	if dest == 1 {
		return 0
	}

	sum := 0
	depth := 1
	min, max := 0, 1

	for {
		min = max + 1
		max = (min - 1) + (depth*2)*4
		if dest > max {
			depth++
			continue
		}
		break
	}

	dest -= min
	mid := depth - 1
	side := (dest % (depth * 2))
	sum = mid - side
	if sum < 0 {
		sum = -sum
	}
	sum += depth

	return sum
}

// B return the lowest number in the spiral higher then input, given that
// all squares are filled with the sum of all of it lower neighbors
func B(input int) int {
	if input < 1 {
		return 1
	}

	prev := []int{1}
	depth := 1

	for {
		armLen := depth * 2
		next := make([]int, armLen*4)
		var armNum int
		adjSide := make([]int, armLen-1)
		for i := range next {
			sum := 0
			armPos := i % armLen
			if armPos == 0 {
				armNum = i / armLen
				for j := range adjSide[1:] {
					adjSide[j+1] = prev[armNum*(armLen-2)+(j)]
				}
				firstPos := armNum*(armLen-2) - 1
				if firstPos < 0 {
					adjSide[0] = prev[len(prev)-1]
				} else {
					adjSide[0] = prev[firstPos]
				}
				if armNum == 3 {
					adjSide = append(adjSide, next[0])
				}
			}
			if i > 0 {
				sum += next[i-1]
				if armPos == 0 {
					sum += next[i-2]
				}
			}
			if armPos < (len(adjSide)) {
				sum += adjSide[armPos]
			}
			if armPos < (len(adjSide) - 1) {
				sum += adjSide[armPos+1]
			}
			if armPos > 0 {
				sum += adjSide[armPos-1]
			}
			if sum > input {
				return sum
			}
			next[i] = sum
		}
		prev = next
		depth++
	}
}
