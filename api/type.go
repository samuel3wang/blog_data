package api

import(
	"time"
)

type Blog struct {
	Title 		string		`json:"title"`
	Content 	string		`json:"content"`
	Category 	string		`json:"category"`
	Date		time.Time	`json:"date"`
}