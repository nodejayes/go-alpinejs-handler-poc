package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	anythingparsejson "github.com/nodejayes/anything-parse-json"
	di "github.com/nodejayes/generic-di"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/todo_app/data"
	"github.com/nodejayes/go-alpinejs-handler-poc/todo_app/state"
)

type (
	Todo          struct{}
	TodoArguments struct {
		Operation string `json:"operation"`
		Value     any    `json:"value"`
	}
)

func (ctx *Todo) GetName() string {
	return "todo"
}

func (ctx *Todo) GetActionType() string {
	return fmt.Sprintf("[%s] operation", ctx.GetName())
}

func (ctx *Todo) GetDefaultState() any {
	return state.NewTodoState()
}

func (ctx *Todo) OnDestroy(clientID string) {
	di.Destroy[state.Todo](clientID)
}

func (ctx *Todo) Handle(msg goalpinejshandler.Message, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools) {
	args, err := anythingparsejson.Parse[TodoArguments](msg.Payload)
	if err != nil {
		return
	}
	clientId := tools.GetClientId(req)
	if len(clientId) < 1 {
		return
	}
	state := di.Inject[state.Todo](clientId)

	switch args.Operation {
	case "add":
		todoToAdd, err := anythingparsejson.Parse[data.Todo](args.Value)
		if err != nil {
			return
		}
		todoToAdd.ID = uuid.NewString()
		state.Add(todoToAdd)
	case "remove":
		todoToRemoveId, err := anythingparsejson.Parse[string](args.Value)
		if err != nil {
			return
		}
		state.Remove(todoToRemoveId)
	case "toggle":
		todoToToggleId, err := anythingparsejson.Parse[string](args.Value)
		if err != nil {
			return
		}
		state.Toggle(todoToToggleId)
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
