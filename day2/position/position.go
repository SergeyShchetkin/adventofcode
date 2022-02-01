package position

import (
	"adventofcode/common"
	"log"
	"strings"
	"sync"
)

type Position struct {
	WG  sync.WaitGroup
	Aim struct {
		Use bool
		val int
	}
	horizontal struct {
		mx  sync.Mutex
		val int
	}
	depth struct {
		mx  sync.Mutex
		val int
	}
}

func (p *Position) Calc(s string) {
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
		if p.isUseAim() {
			p.horizontal.val += dVal
			p.depth.val += p.Aim.val * dVal
		} else {
			p.horizontal.mx.Lock()
			p.horizontal.val += dVal
			p.horizontal.mx.Unlock()
		}
	} else {
		if dType == up {
			dVal = -dVal
		}

		if p.isUseAim() {
			p.Aim.val += dVal
		} else {
			p.depth.mx.Lock()
			p.depth.val += dVal
			p.depth.mx.Unlock()
		}
	}

	if !p.isUseAim() {
		p.WG.Done()
	}
}

func (p *Position) isUseAim() bool {
	return p.Aim.Use
}

func (p *Position) Result() int {
	if !p.isUseAim() {
		p.WG.Wait()
	}

	return p.horizontal.val * p.depth.val
}
