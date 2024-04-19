package img

import _ "embed"

const FaviconType = "image/x-icon"

//go:embed todo_favicon.ico
var TodoFavicon []byte

//go:embed counter_favicon.ico
var CounterFavicon []byte
