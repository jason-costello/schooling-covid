package models

type School struct {
	ID         int
	DistrictID int
	Name       string
	ShortName  string
	Counts     []Count
}
