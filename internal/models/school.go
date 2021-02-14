package models


// School captures the data for an individual school that can only belong to one district.  A school will have
// many school counts
type School struct {
	DistrictShortName string
	Name              string
	ShortName         string
	Counts            SchoolCounts
}

