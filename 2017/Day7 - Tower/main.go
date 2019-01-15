package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type program struct {
	name     string
	weight   int
	children []*program
	parent   *program
}

func main() {

	a, err := A()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("A:", a)

	// b, err := B()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("B:", b)
}

// A - returns the name of the program at the bottom of the tower
func A() (string, error) {
	f, err := os.Open("inputA")
	if err != nil {
		return "", err
	}

	r := bufio.NewReader(f)
	tree, err := createTree(r)
	if err != nil {
		return "", err
	}
	root, err := findRoot(tree)
	if err != nil {
		return "", err
	}

	return root.name, nil
}

func createTree(r *bufio.Reader) (map[string]program, error) {
	tree := make(map[string]program)
	for {
		program, err := readRow(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if _, ok := tree[program.name]; ok {
			return nil, errors.New("program duplicate")
		}

		tree[program.name] = program
	}
	return tree, nil
}

func findRoot(tree map[string]program) (program, error) {
	return program{name: "foo"}, nil
}

// B returns the length of the loop
// func B() (int, error) {
// 	f, err := os.Open("inputA")
// 	if err != nil {
// 		return 0, err
// 	}

// 	r := bufio.NewReader(f)
// 	input, err := readRow(r)
// 	if err != nil && err != io.EOF {
// 		return 0, err
// 	}

// 	sum := 0
// 	history := make([][]int, 0)
// 	for {
// 		max := findMaxIndex(input)
// 		realocate(input, max)
// 		sum++
// 		if isALoop(input, history) {
// 			fmt.Println("history len:", len(history))
// 			// answer is (historyLen - c)
// 			break
// 		}
// 		copyS := make([]int, len(input))
// 		copy(copyS, input)
// 		history = append(history, copyS)
// 	}
// 	return sum, nil
// }

func readRow(r *bufio.Reader) (program, error) {
	s, err := r.ReadString('\n')
	if err != nil {
		return program{}, err
	}

	var (
		name      string
		weight    int
		childrenS string
		children  []*program
	)

	_, err = fmt.Sscanf(s, "%s (%d) -> %s\r\n", &name, &weight, &childrenS)
	if err != nil {
		return program{}, err
	}

	return program{name: name, weight: weight, children: children}, nil
}
