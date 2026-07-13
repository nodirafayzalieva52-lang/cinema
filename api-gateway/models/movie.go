package models

import(
	"time"
)

type CreateMovieRequest struct {
	ID       	int				`json:"ID"`
	Title    	string			`json:"Title"`
	Description string			`json:"Description"`
	Duration 	time.Duration	`json:"Duration"`
	Age_Limit 	int				`json:"Age_Limit"`
	CreatedAt   time.Time		`json:"CreatedAt"`
}