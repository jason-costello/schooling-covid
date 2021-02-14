package models

// District is the root level for the data structures used to capture reported count data.  A district has many schools and
// a school has many counts.    A school can only be associated to one district.   A count can only be associated to one
// school
type District struct {
	Name      string
	ShortName string
	Schools   []School
}
