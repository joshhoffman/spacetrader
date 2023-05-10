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

type nav struct {
	SystemSymbol   string `json:"systemSymbol"`
	WaypointSymbol string `json:"waypointSymbol"`
	Route          route  `json:"route"`
	Status         string `json:"status"`
	FlightMode     string `json:"flightMode"`
}

type route struct {
	Departure     departure `json:"departure"`
	Destination   departure `json:"destination"`
	Arrival       string    `json:"arrival"`
	DepartureTime string    `json:"departureTime"`
}

type departure struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int64  `json:"x"`
	Y            int64  `json:"y"`
}
