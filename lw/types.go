package lw

import (
		"polaroid/types"
)

type Lw struct {
	Data *types.Data
}

func NewLw(s *types.Data) (a *Lw) {
	a = new(Lw)
	a.Data = s
	return
}

type Fortune struct {
	Fortune string
	Data []string
	Datas []string
	Datal []string
	I string
}
