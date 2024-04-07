package toaster

import (
	"github.com/google/uuid"
	di "github.com/nodejayes/generic-di"
	"slices"
	"sync"
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
		ID      string `json:"id"`
		Typ     string `json:"typ"`
		Message string `json:"message"`
		Ttl     int    `json:"ttl"`
	}
	State struct {
		m           *sync.Mutex
		MaxMessages int       `json:"maxMessages"`
		Messages    []Message `json:"messages"`
	}
)

func NewState() *State {
	return &State{
		m:           &sync.Mutex{},
		MaxMessages: 5,
		Messages:    make([]Message, 0),
	}
}

func (ctx *State) SetMaxMessages(count int) {
	ctx.m.Lock()
	defer ctx.m.Unlock()

	if count <= 0 {
		count = 5
	}
	ctx.MaxMessages = count
}

func (ctx *State) AddMessage(msg Message) Message {
	ctx.m.Lock()
	defer ctx.m.Unlock()

	if len(msg.ID) < 1 {
		msg.ID = uuid.NewString()
	}
	idx := ctx.getMessageIdx(msg)
	if idx >= 0 {
		return msg
	}
	start := len(ctx.Messages) - (ctx.MaxMessages - 1)
	if start < 0 {
		ctx.Messages = append(ctx.Messages, msg)
		return msg
	}
	ctx.Messages = ctx.Messages[start:]
	ctx.Messages = append(ctx.Messages, msg)
	return msg
}

func (ctx *State) RemoveMessage(msg Message) {
	ctx.m.Lock()
	defer ctx.m.Unlock()

	idx := ctx.getMessageIdx(msg)
	if idx >= 0 {
		ctx.Messages = append(ctx.Messages[:idx], ctx.Messages[idx+1:]...)
	}
}

func (ctx *State) getMessageIdx(msg Message) int {
	return slices.IndexFunc(ctx.Messages, func(message Message) bool {
		return message.ID == msg.ID
	})
}
