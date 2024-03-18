package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (ctx *Counter) Authorized(msg goalpinejshandler.Message, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool) error {
	return nil
}

func (ctx *Counter) GetDefaultState() string {
	stream, err := json.Marshal(state.Counter{
		Value: 0,
	})
	if err != nil {
		return ""
	}
	return string(stream)
}

func (ctx *Counter) Handle(msg goalpinejshandler.Message, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool) {
	content, err := json.Marshal(msg.Payload)
	if err != nil {
		return
	}
	var args CounterArguments
	err = json.Unmarshal(content, &args)
	if err != nil {
		return
	}

	clientId := req.Header.Get("clientId")
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
