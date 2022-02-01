package main

import (
	"adventofcode/common"
	"adventofcode/day2/position"
	"bufio"
	"fmt"
)

func main() {
	ps := &position.Position{}
	ps.Aim.Use = true

	f := common.OpenFile("input_data.txt")
	sc := bufio.NewScanner(f)
	defer f.Close()

	for sc.Scan() {
		ps.Calc(sc.Text())
	}

	fmt.Printf("Result=%d\n", ps.Result())
}
