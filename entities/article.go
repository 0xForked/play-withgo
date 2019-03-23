package entities

type Article struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Category *Category `json:"category"`
}
