package models

import "time"

type Count struct {
	ID          int
	Observed    time.Time
	Positive    int
	Symptomatic int
	Exposed     int
}
