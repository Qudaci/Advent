package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {

	a, err := A()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("A:", a)
	}

	b, err := B()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("B:", b)
	}
}

// A - returns a sum of all digits that match the next digit in the list
func A() (int, error) {
	f, err := os.Open("inputA")
	if err != nil {
		return 0, err
	}

	r := bufio.NewReader(f)
	d, err := readDigit(r)
	if err == io.EOF {
		return 0, errors.New("input is empty")
	}
	if err != nil {
		return 0, err
	}
	first := d
	last := first
	sum := 0

	for {
		d, err := readDigit(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		if last == d {
			sum += d
		}

		last = d
	}
	if first == last {
		sum += last
	}
	return sum, nil
}

// B returns a sum of matching digits halway around the circular list
func B() (int, error) {
	f, err := ioutil.ReadFile("inputA")
	if err != nil {
		return 0, err
	}

	s := string(f)
	if len(s)%2 != 0 {
		return 0, errors.New("input is uneven")
	}
	list := make([]int, len(s))

	for i, c := range s {
		d, err := strconv.Atoi(string(c))
		if err != nil {
			return 0, err
		}

		list[i] = d
	}

	sum := 0
	half := len(list) / 2

	for i := 0; i < half; i++ {
		if list[i] == list[i+half] {
			sum += list[i] * 2
		}
	}

	return sum, nil
}

func readDigit(r *bufio.Reader) (int, error) {
	c, _, err := r.ReadRune()
	if err != nil {
		return 0, err
	}
	d, err := strconv.Atoi(string(c))
	if err != nil {
		return 0, err
	}
	if d < 0 || d > 9 {
		return 0, fmt.Errorf("rune %s is not a number", string(c))
	}
	return d, nil
}
