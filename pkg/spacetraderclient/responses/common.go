package responses

type Faction struct {
	Symbol string `json:"symbol"`
}

type SubmittedBy string

const (
	Cosmic SubmittedBy = "COSMIC"
)

type SystemSymbol string

const (
	X1Df55 SystemSymbol = "X1-DF55"
)

type Meta struct {
	Total int64 `json:"total"`
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}
