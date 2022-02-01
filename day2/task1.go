package main

import (
	"adventofcode/common"
	"bufio"
	"fmt"
	"log"
	"strings"
	"sync"
)

func main() {
	ps := &position{}
	f := common.OpenFile("input_data.txt")
	sc := bufio.NewScanner(f)
	defer f.Close()

	for sc.Scan() {
		ps.wg.Add(1)
		go ps.calc(sc.Text())
	}

	fmt.Printf("Result=%d\n", ps.result())
}

type position struct {
	wg         sync.WaitGroup
	horizontal struct {
		mx  sync.Mutex
		val int
	}
	depth struct {
		mx  sync.Mutex
		val int
	}
}

func (p *position) calc(s string) {
	var (
		forward = "forward"
		up      = "up"
	)

	data := strings.Split(strings.Trim(s, " "), " ")
	dType := data[0]
	dVal, err := common.StrToInt(data[1])
	if len(data) != 2 || err != nil {
		log.Fatalf("Wrong input data '%s'", s)
	}

	if dType == forward {
		p.horizontal.mx.Lock()
		p.horizontal.val += dVal
		p.horizontal.mx.Unlock()
	} else {
		if dType == up {
			dVal = -dVal
		}

		p.depth.mx.Lock()
		p.depth.val += dVal
		p.depth.mx.Unlock()
	}

	p.wg.Done()
}

func (p *position) result() int {
	p.wg.Wait()
	return p.horizontal.val * p.depth.val
}
