package handler

import (
	"fmt"
	"net/http"

	anythingparsejson "github.com/nodejayes/anything-parse-json"
	di "github.com/nodejayes/generic-di"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/counter_app/state"
)

type (
	Counter          struct{}
	CounterArguments struct {
		Operation string `json:"operation"`
		Value     int    `json:"value"`
	}
)

func (ctx *Counter) GetName() string {
	return "counter"
}

func (ctx *Counter) GetActionType() string {
	return fmt.Sprintf("[%s] operation", ctx.GetName())
}

func (ctx *Counter) Authorized(msg goalpinejshandler.Message, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools) error {
	return nil
}

func (ctx *Counter) GetDefaultState() any {
	return state.NewCounter()
}

func (ctx *Counter) OnDestroy(clientID string) {
	di.Destroy[state.Counter](clientID)
}

func (ctx *Counter) Handle(msg goalpinejshandler.Message, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools) {
	args, err := anythingparsejson.Parse[CounterArguments](msg.Payload)
	if err != nil {
		return
	}

	clientId := tools.GetClientId(req)
	state := di.Inject[state.Counter](clientId)
	if state == nil {
		return
	}

	switch args.Operation {
	case "add":
		state.Add(args.Value)
	case "sub":
		state.Sub(args.Value)
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
