package data

type Todo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Open bool   `json:"open"`
}
