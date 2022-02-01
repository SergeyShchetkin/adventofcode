package common

import (
	"os"
	"strconv"
)

func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func OpenFile(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return f
}

func SliceSum(sl []int) int {
	var sum int
	for _, i := range sl {
		sum += i
	}

	return sum
}
