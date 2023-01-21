package api

import(
	"time"
)

type Blog struct {
	Title_id	int64			`json:"title_id"`
	Title			string		`json:"title"`
	Content		string		`json:"content"`
	Category	string		`json:"category"`
	Image			string		`json:"image"`
	Time		time.Time		`json:"time"`
}
type Category struct {
	Category string `json:"category"`
}