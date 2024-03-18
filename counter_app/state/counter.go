package state

import di "github.com/nodejayes/generic-di"

func init() {
	di.Injectable(NewCounter)
}

type Counter struct {
	Value int `json:"value"`
}

func NewCounter() *Counter {
	return &Counter{}
}

func (ctx *Counter) Add(value int) {
	ctx.Value += value
}

func (ctx *Counter) Sub(value int) {
	ctx.Value -= value
}
