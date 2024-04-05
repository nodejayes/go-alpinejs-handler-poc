package toaster

import (
	"fmt"
	anythingparsejson "github.com/nodejayes/anything-parse-json"
	di "github.com/nodejayes/generic-di"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"net/http"
)

func init() {
	di.Injectable(NewHandler)
}

const name = "cosmic_toaster"

type (
	Handler          struct{}
	HandlerArguments struct {
		Operation string  `json:"operation"`
		Value     Message `json:"value"`
	}
)

func NewHandler() *Handler {
	return &Handler{}
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
		state.AddMessage(args.Value)
	case "remove":
		state.RemoveMessage(args.Value)
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
