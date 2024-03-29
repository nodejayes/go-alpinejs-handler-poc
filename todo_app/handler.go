package todo_app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	anythingparsejson "github.com/nodejayes/anything-parse-json"
	di "github.com/nodejayes/generic-di"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
)

type (
	TodoHandler   struct{}
	TodoArguments struct {
		Operation string `json:"operation"`
		Value     any    `json:"value"`
	}
)

func (ctx *TodoHandler) GetName() string {
	return "todo"
}

func (ctx *TodoHandler) GetActionType() string {
	return fmt.Sprintf("[%s] operation", ctx.GetName())
}

func (ctx *TodoHandler) GetDefaultState() any {
	return NewTodoState()
}

func (ctx *TodoHandler) OnDestroy(clientID string, tools *goalpinejshandler.Tools) {
	go func() {
		time.Sleep(60 * time.Second)
		if tools.HasConnections(clientID) {
			return
		}
		di.Destroy[State](clientID)
	}()
}

func (ctx *TodoHandler) Handle(msg goalpinejshandler.Message, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools) {
	args, err := anythingparsejson.Parse[TodoArguments](msg.Payload)
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
		todoToAdd, err := anythingparsejson.Parse[Todo](args.Value)
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
