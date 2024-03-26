package toaster

import (
	di "github.com/nodejayes/generic-di"
	"slices"
)

func init() {
	di.Injectable(NewToasterState)
}

type State struct {
	Messages []Message `json:"messages"`
}

func NewToasterState() *State {
	return &State{}
}

func (ctx *State) Add(message Message) {
	if len(ctx.Messages) >= 10 {
		ctx.Messages = ctx.Messages[1:]
	}
	ctx.Messages = append(ctx.Messages, message)
}

func (ctx *State) Remove(message Message) {
	foundIdx := slices.IndexFunc(ctx.Messages, func(m Message) bool {
		return m.Type == message.Type && m.Content == message.Content
	})
	if foundIdx < 0 {
		return
	}
	ctx.Messages = append(ctx.Messages[:foundIdx], ctx.Messages[foundIdx+1:]...)
}
