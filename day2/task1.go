package main

import (
	"adventofcode/common"
	"adventofcode/day2/position"
	"bufio"
	"fmt"
)

func main() {
	ps := &position.Position{}
	f := common.OpenFile("input_data.txt")
	sc := bufio.NewScanner(f)
	defer f.Close()

	for sc.Scan() {
		ps.WG.Add(1)
		go ps.Calc(sc.Text())
	}

	fmt.Printf("Result=%d\n", ps.Result())
}
