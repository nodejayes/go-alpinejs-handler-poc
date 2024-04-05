package toaster

import (
	di "github.com/nodejayes/generic-di"
	"slices"
)

func init() {
	di.Injectable(NewState)
}

const (
	DangerType  = "danger"
	WarningType = "warning"
	SuccessType = "success"
)

type (
	Message struct {
		Typ     string `json:"typ"`
		Message string `json:"message"`
	}
	State struct {
		MaxMessages int       `json:"maxMessages"`
		Messages    []Message `json:"messages"`
	}
)

func NewState() *State {
	return &State{
		MaxMessages: 5,
		Messages:    make([]Message, 0),
	}
}

func (ctx *State) SetMaxMessages(count int) {
	if count <= 0 {
		count = 5
	}
	ctx.MaxMessages = count
}

func (ctx *State) AddMessage(msg Message) {
	idx := ctx.getMessageIdx(msg)
	if idx >= 0 {
		return
	}
	start := len(ctx.Messages) - (ctx.MaxMessages - 1)
	if start < 0 {
		ctx.Messages = append(ctx.Messages, msg)
		return
	}
	ctx.Messages = ctx.Messages[start:]
	ctx.Messages = append(ctx.Messages, msg)
}

func (ctx *State) RemoveMessage(msg Message) {
	idx := ctx.getMessageIdx(msg)
	if idx >= 0 {
		ctx.Messages = append(ctx.Messages[:idx], ctx.Messages[idx+1:]...)
	}
}

func (ctx *State) getMessageIdx(msg Message) int {
	return slices.IndexFunc(ctx.Messages, func(message Message) bool {
		return message.Typ == msg.Typ && message.Message == msg.Message
	})
}
