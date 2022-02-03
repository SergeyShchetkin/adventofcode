package main

import (
	"adventofcode/common"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var (
		wg           sync.WaitGroup
		oxygenSlice  []string
		co2Slice     []string
		oxygenBitNum int
		co2BitNum    int
	)

	data := common.ReadFile("input_data.txt")

	wg.Add(1)
	go func() {
		oxygenSlice = filter(data, &oxygenBitNum, true)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		co2Slice = filter(data, &co2BitNum, false)
		wg.Done()
	}()

	wg.Wait()
	oxygenStr := getDataItemByLastBit(oxygenSlice, true, oxygenBitNum)
	co2Str := getDataItemByLastBit(co2Slice, false, co2BitNum)
	oxygen, _ := strconv.ParseInt(oxygenStr, 2, 64)
	co2, _ := strconv.ParseInt(co2Str, 2, 64)
	fmt.Printf("Result=%d\n", oxygen*co2)
}

func getDataItemByLastBit(data []string, checkBitMoreZero bool, bitNum int) string {
	if len(data) == 1 {
		return data[0]
	} else {
		var res string
		for _, str := range data {
			sl := strings.Split(str, "")
			checkBit, _ := common.StrToInt(sl[bitNum])
			if checkBitMoreZero && checkBit == 0 {
				continue
			}
			res = str
		}

		return res
	}
}

func filter(data []string, bitNum *int, byMax bool) []string {
	var (
		result []string
		zero   []string
		first  []string
	)

	for _, str := range data {
		sl := strings.Split(str, "")
		iVal, _ := common.StrToInt(sl[*bitNum])
		if iVal == 0 {
			zero = append(zero, str)
		} else {
			first = append(first, str)
		}
	}

	firstMax := len(first) >= len(zero)
	if byMax {
		if firstMax {
			result = first
		} else {
			result = zero
		}
	} else {
		if firstMax {
			result = zero
		} else {
			result = first
		}
	}

	*bitNum++
	if len(result) > 2 {
		result = filter(result, bitNum, byMax)
	}

	return result
}
