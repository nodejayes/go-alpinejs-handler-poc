package toaster

import (
	"fmt"
	anythingparsejson "github.com/nodejayes/anything-parse-json"
	di "github.com/nodejayes/generic-di"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"net/http"
)

type (
	Handler   struct{}
	Arguments struct {
		Operation string `json:"operation"`
		Value     any    `json:"value"`
	}
)

func (ctx *Handler) GetName() string {
	return "toaster"
}

func (ctx *Handler) GetActionType() string {
	return fmt.Sprintf("[%s] operation", ctx.GetName())
}

func (ctx *Handler) GetDefaultState() any {
	return NewToasterState()
}

func (ctx *Handler) OnDestroy(clientID string) {
	di.Destroy[State](clientID)
}

func (ctx *Handler) Handle(msg goalpinejshandler.Message, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools) {
	args, err := anythingparsejson.Parse[Arguments](msg.Payload)
	if err != nil {
		return
	}
	clientId := tools.GetClientId(req)
	if len(clientId) < 1 {
		return
	}
	state := di.Inject[State](clientId)

	switch args.Operation {
	case "add":
		message, err := anythingparsejson.Parse[Message](args.Value)
		if err != nil {
			return
		}
		state.Add(message)
	case "remove":
		message, err := anythingparsejson.Parse[Message](args.Value)
		if err != nil {
			return
		}
		state.Remove(message)
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
