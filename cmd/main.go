package main

import (
	"log"
	"net/http"

	contextstore "github.com/nodejayes/context-store"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
	"github.com/nodejayes/go-alpinejs-handler-poc/counter_app"
	"github.com/nodejayes/go-alpinejs-handler-poc/img"
	"github.com/nodejayes/go-alpinejs-handler-poc/todo_app"

	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
)

var icons = map[string][]byte{
	"/icons/todo_favicon.ico":    img.TodoFavicon,
	"/icons/counter_favicon.ico": img.CounterFavicon,
}

func getPages() []goalpinejshandler.Page {
	return []goalpinejshandler.Page{
		counter_app.NewPage(),
		todo_app.NewPage(),
	}
}

func getConfig() goalpinejshandler.Config {
	return goalpinejshandler.Config{
		ActionUrl:         "/action",
		EventUrl:          "/events",
		ClientIDHeaderKey: "clientId",
		Pages:             getPages(),
	}
}

func linkIcons(router *http.ServeMux) {
	for route, icon := range icons {
		router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("content-type", img.FaviconType)
			w.Write(icon)
		})
	}
}

func main() {
	defer contextstore.Clear()
	
	config := getConfig()
	router := http.NewServeMux()
	goalpinejshandler.Register(router, &config)
	cosmic_ui_global.RegisterGlobalStyles()
	linkIcons(router)
	err := http.ListenAndServe(":40000", router)
	if err != nil {
		log.Fatal(err)
	}
}
