package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	// a, err := A()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("A:", a)

	b, err := B()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("B:", b)
}

// A - returns the number of valid passphrases
func A() (int, error) {
	f, err := os.Open("inputA")
	if err != nil {
		return 0, err
	}

	r := bufio.NewReader(f)
	sum := 0

LOOP:
	for {
		row, err := readRow(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		for i, m := range row {
			for _, n := range row[i+1:] {
				if m == n {
					continue LOOP
				}
			}

		}
		sum++
	}
	return sum, nil
}

func readRow(r *bufio.Reader) ([]string, error) {
	s, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	s = strings.TrimRight(s, "\n\r")

	return strings.Split(s, " "), nil
}

// B - returns the number of valid annagram passphrases
func B() (int, error) {
	f, err := os.Open("inputA")
	if err != nil {
		return 0, err
	}

	r := bufio.NewReader(f)
	sum := 0

LOOP:
	for {
		row, err := readRow(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		for i, m := range row {
			for _, n := range row[i+1:] {
				if isAnnagram(m, n) {
					continue LOOP
				}
			}

		}
		sum++
	}
	return sum, nil
}

func isAnnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
LOOP:
	for _, m := range s1 {
		for i, n := range s2 {
			if m == n {
				s2 = strings.Join([]string{s2[:i], s2[i+1:]}, "")
				continue LOOP
			}
		}
		return false
	}

	return true
}
