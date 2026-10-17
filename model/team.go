package model

type Team struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	City         string `json:"city"`
	Abbreviation string `json:"abbreviation"`
	Conference   string `json:"conference"`
	Division     string `json:"division"`
}
