package models

type District struct {
	ID        int
	Name      string
	ShortName string
	Schools   []School
}
