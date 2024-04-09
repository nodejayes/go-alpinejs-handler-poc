package toaster

import (
	"fmt"
	anythingparsejson "github.com/nodejayes/anything-parse-json"
	di "github.com/nodejayes/generic-di"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"net/http"
	"time"
)

func init() {
	di.Injectable(NewHandler)
}

const name = "cosmic_toaster"

type (
	Handler struct {
		messageRemoveAnimationTime time.Duration
	}
	HandlerArguments struct {
		Operation string  `json:"operation"`
		Value     Message `json:"value"`
	}
)

func NewHandler() *Handler {
	return &Handler{
		messageRemoveAnimationTime: 300 * time.Millisecond,
	}
}

func (ctx *Handler) GetName() string {
	return name
}

func (ctx *Handler) GetActionType() string {
	return fmt.Sprintf("[%s] operation", ctx.GetName())
}

func (ctx *Handler) GetDefaultState() any {
	return NewState()
}

func (ctx *Handler) Handle(msg goalpinejshandler.Message, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools) {
	args, err := anythingparsejson.Parse[HandlerArguments](msg.Payload)
	if err != nil {
		return
	}
	state := di.Inject[State]()
	clientId := tools.GetClientId(req)

	switch args.Operation {
	case "add":
		if args.Value.Ttl == 0 {
			args.Value.Ttl = 5
		}
		args.Value = state.AddMessage(args.Value)
		if args.Value.Ttl > -1 {
			go func() {
				time.Sleep(time.Duration(args.Value.Ttl) * time.Second)
				removeMessage(args.Value, state, messagePool, clientId, ctx.GetName(), ctx.messageRemoveAnimationTime)
				messagePool.Add(goalpinejshandler.ChannelMessage{
					ClientFilter: func(client goalpinejshandler.Client) bool {
						return client.ID == clientId
					},
					Message: goalpinejshandler.Message{
						Type:    fmt.Sprintf("[%s] update", name),
						Payload: state,
					},
				})
			}()
		}
		if state.MessagePoolFull() {
			oldestMessage, err := state.GetOldestMessage()
			if err == nil {
				removeMessage(oldestMessage, state, messagePool, clientId, ctx.GetName(), ctx.messageRemoveAnimationTime)
			}
		}
	case "remove":
		removeMessage(args.Value, state, messagePool, clientId, ctx.GetName(), ctx.messageRemoveAnimationTime)
	case "animation_start":
		state.UpdateOpenState(args.Value.ID, args.Value.Open)
		return
	}

	messagePool.Add(goalpinejshandler.ChannelMessage{
		ClientFilter: func(client goalpinejshandler.Client) bool {
			return client.ID == clientId
		},
		Message: goalpinejshandler.Message{
			Type:    fmt.Sprintf("[%s] update", ctx.GetName()),
			Payload: state,
		},
	})
}

func removeMessage(message Message, state *State, messagePool *goalpinejshandler.MessagePool, clientId string, name string, animationTime time.Duration) {
	state.UpdateOpenState(message.ID, false)
	go func() {
		time.Sleep(animationTime)
		state.RemoveMessage(message)
		messagePool.Add(goalpinejshandler.ChannelMessage{
			ClientFilter: func(client goalpinejshandler.Client) bool {
				return client.ID == clientId
			},
			Message: goalpinejshandler.Message{
				Type:    fmt.Sprintf("[%s] update", name),
				Payload: state,
			},
		})
	}()
}
