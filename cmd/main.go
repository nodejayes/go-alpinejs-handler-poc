package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/counter_app/handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/counter_app/template"
)

func getConfig() goalpinejshandler.Config {
	return goalpinejshandler.Config{
		ActionUrl:            "/action",
		EventUrl:             "/events",
		ClientIDHeaderKey:    "clientId",
		SendConnectedAfterMs: 100,
		Handlers: []goalpinejshandler.ActionHandler{
			&handler.Counter{},
		},
	}
}

func main() {
	config := getConfig()
	router := http.NewServeMux()
	goalpinejshandler.Register(router, &config)
	router.Handle("/counter", templ.Handler(template.Index(goalpinejshandler.HeadScripts())))
	http.ListenAndServe(":40000", router)
}