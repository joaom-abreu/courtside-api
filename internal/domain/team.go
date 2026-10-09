package domain

type Conference string

const (
	ConferenceEast Conference = "East"
	ConferenceWest Conference = "West"
)

func (c Conference) IsValid() bool {
	return c == ConferenceEast || c == ConferenceWest
}

type Team struct {
	ID           int64
	Name         string
	City         string
	Abbreviation string
	Conference   Conference
	Division     string
}

func (t Team) FullName() string {
	return t.City + " " + t.Name
}
