package main

import (
	"adventofcode/common"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type bits string

func (b bits) parse(cum *[12][2]int) {
	sl := strings.Split(string(b), "")
	for i, sVal := range sl {
		iVal, _ := common.StrToInt(sVal)
		cum[i][iVal]++
	}
}

func main() {
	var cum [12][2]int
	f := common.OpenFile("input_data.txt")
	sc := bufio.NewScanner(f)
	defer f.Close()

	for sc.Scan() {
		(bits(sc.Text())).parse(&cum)
	}

	g, e := makeGammaEpsilon(cum)
	gamma, _ := strconv.ParseInt(g, 2, 64)
	epsilon, _ := strconv.ParseInt(e, 2, 64)
	fmt.Printf("Result=%d\n", gamma*epsilon)
}

func makeGammaEpsilon(cum [12][2]int) (string, string) {
	gSl := make([]string, 12)
	eSl := make([]string, 12)

	for i, val := range cum {
		if val[0] > val[1] {
			gSl[i] = strconv.Itoa(0)
			eSl[i] = strconv.Itoa(1)
		} else {
			gSl[i] = strconv.Itoa(1)
			eSl[i] = strconv.Itoa(0)
		}
	}

	return strings.Join(gSl, ""), strings.Join(eSl, "")
}
