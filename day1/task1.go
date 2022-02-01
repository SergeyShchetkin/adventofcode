package main

import (
	"adventofcode/common"
	"bufio"
	"fmt"
)

func main() {
	var (
		step  int
		count = -1
	)

	f := common.OpenFile("input_data.txt")
	sc := bufio.NewScanner(f)
	defer f.Close()

	for sc.Scan() {
		if depth, err := common.StrToInt(sc.Text()); err == nil {
			if depth > step {
				count++
			}
			step = depth
		} else {
			panic(err)
		}
	}

	fmt.Printf("count=%d\n", count)
}
