package common

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

func ReadFile(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(content), "\n")
}

func SliceSum(sl []int) int {
	var sum int
	for _, i := range sl {
		sum += i
	}

	return sum
}
