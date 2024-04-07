package todo_app

import (
	"fmt"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/toaster"
	"net/http"
	"time"

	"github.com/google/uuid"
	anythingparsejson "github.com/nodejayes/anything-parse-json"
	di "github.com/nodejayes/generic-di"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
)

func init() {
	di.Injectable(NewTodoHandler)
}

type (
	TodoHandler struct {
		Toaster *toaster.Handler
	}
	TodoArguments struct {
		Operation string `json:"operation"`
		Value     any    `json:"value"`
	}
)

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		Toaster: di.Inject[toaster.Handler](),
	}
}

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
		if len(todoToAdd.Name) < 3 {
			ctx.sendMessage(toaster.DangerType, "Activity must have a label with min 3 Chars", res, req, messagePool, tools)
			return
		}
		todoToAdd.ID = uuid.NewString()
		todoToAdd.Open = true
		state.Add(todoToAdd)
		ctx.sendMessage(toaster.SuccessType, fmt.Sprintf("Activity %s added", todoToAdd.Name), res, req, messagePool, tools)
	case "remove":
		todoToRemoveId, err := anythingparsejson.Parse[string](args.Value)
		if err != nil {
			ctx.sendMessage(toaster.DangerType, err.Error(), res, req, messagePool, tools)
			return
		}
		state.Remove(todoToRemoveId)
	case "toggle":
		todoToToggleId, err := anythingparsejson.Parse[string](args.Value)
		if err != nil {
			ctx.sendMessage(toaster.DangerType, err.Error(), res, req, messagePool, tools)
			return
		}
		state.Toggle(todoToToggleId)
		todo, err := state.Get(todoToToggleId)
		if err != nil {
			ctx.sendMessage(toaster.DangerType, err.Error(), res, req, messagePool, tools)
			return
		}
		activeLabel := "open"
		if !todo.Open {
			activeLabel = "finish"
		}
		ctx.sendMessage(toaster.SuccessType, fmt.Sprintf("State of Activity %s was set to %s", todo.Name, activeLabel), res, req, messagePool, tools)
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

func (ctx *TodoHandler) sendMessage(messageTyp, message string, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools) {
	ctx.Toaster.Handle(goalpinejshandler.Message{
		Type: ctx.Toaster.GetActionType(),
		Payload: toaster.HandlerArguments{
			Operation: "add",
			Value: toaster.Message{
				Typ:     messageTyp,
				Message: message,
			},
		},
	}, res, req, messagePool, tools)
}
