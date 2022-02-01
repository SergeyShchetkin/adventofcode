package main

import (
	"adventofcode/common"
	"bufio"
	"fmt"
)

func main() {
	var (
		step  = 3
		count = -1
		data  []int
		tmp   int
	)

	f := common.OpenFile("input_data.txt")
	sc := bufio.NewScanner(f)
	defer f.Close()

	for sc.Scan() {
		if item, err := common.StrToInt(sc.Text()); err == nil {
			data = append(data, item)
		} else {
			panic(err)
		}
	}

	for i := 0; i < len(data)-1; i++ {
		sum := common.SliceSum(data[i : i+step])
		if sum > tmp {
			count++
		}

		tmp = sum
	}

	fmt.Printf("count=%d\n", count)
}
