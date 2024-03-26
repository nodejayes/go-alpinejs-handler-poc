package toaster

type Message struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}
