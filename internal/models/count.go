// Package models provides the structures for the datasets that will be returned between the repositories and services
package models

import "time"



// SchoolCounts is multiple SchoolCount data points
type SchoolCounts []SchoolCount


// SchoolCount captures the current counts of types of reported cases for a specific day
type SchoolCount struct {
	CountDate   time.Time
	Positive    int
	Symptomatic int
	Exposed     int
	SchoolSn    string
}

