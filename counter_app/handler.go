package counter_app

import (
	"fmt"
	"net/http"
	"time"

	anythingparsejson "github.com/nodejayes/anything-parse-json"
	di "github.com/nodejayes/generic-di"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
)

type (
	CounterHandler   struct{}
	CounterArguments struct {
		Operation string `json:"operation"`
		Value     int    `json:"value"`
	}
)

func (ctx *CounterHandler) GetName() string {
	return "counter"
}

func (ctx *CounterHandler) GetActionType() string {
	return fmt.Sprintf("[%s] operation", ctx.GetName())
}

func (ctx *CounterHandler) Authorized(msg goalpinejshandler.Message, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools) error {
	return nil
}

func (ctx *CounterHandler) GetDefaultState() any {
	return NewCounter()
}

func (ctx *CounterHandler) OnDestroy(clientID string, tools *goalpinejshandler.Tools) {
	go func() {
		time.Sleep(60 * time.Second)
		if tools.HasConnections(clientID) {
			return
		}
		di.Destroy[Counter](clientID)
	}()
}

func (ctx *CounterHandler) Handle(msg goalpinejshandler.Message, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools) {
	args, err := anythingparsejson.Parse[CounterArguments](msg.Payload)
	if err != nil {
		return
	}

	clientId := tools.GetClientId(req)
	state := di.Inject[Counter](clientId)
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
