package responses

import (
	"encoding/json"

	"github.com/joshhoffman/spacetrader/pkg/client"
)

func (o ShipsResponse) Unmarshal(data []byte) (client.MarshalableObject, error) {
	var r ShipsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (o ShipsResponse) Marshal() ([]byte, error) {
	return json.Marshal(o)
}

type ShipsResponse struct {
	Data []shipsDatum `json:"data"`
	Meta Meta         `json:"meta"`
}

type shipsDatum struct {
	Symbol       string       `json:"symbol"`
	Nav          nav          `json:"nav"`
	Crew         crew         `json:"crew"`
	Fuel         fuel         `json:"fuel"`
	Frame        frame        `json:"frame"`
	Reactor      reactor      `json:"reactor"`
	Engine       engine       `json:"engine"`
	Modules      []module     `json:"modules"`
	Mounts       []mount      `json:"mounts"`
	Registration registration `json:"registration"`
	Cargo        cargo        `json:"cargo"`
}

type cargo struct {
	Capacity  int64       `json:"capacity"`
	Units     int64       `json:"units"`
	Inventory []inventory `json:"inventory"`
}

type inventory struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Units       int64  `json:"units"`
}

type crew struct {
	Current  int64  `json:"current"`
	Capacity int64  `json:"capacity"`
	Required int64  `json:"required"`
	Rotation string `json:"rotation"`
	Morale   int64  `json:"morale"`
	Wages    int64  `json:"wages"`
}

type engine struct {
	Symbol       string             `json:"symbol"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Condition    int64              `json:"condition"`
	Speed        int64              `json:"speed"`
	Requirements engineRequirements `json:"requirements"`
}

type engineRequirements struct {
	Power int64 `json:"power"`
	Crew  int64 `json:"crew"`
}

type frame struct {
	Symbol         string             `json:"symbol"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	ModuleSlots    int64              `json:"moduleSlots"`
	MountingPoints int64              `json:"mountingPoints"`
	FuelCapacity   int64              `json:"fuelCapacity"`
	Condition      int64              `json:"condition"`
	Requirements   engineRequirements `json:"requirements"`
}

type fuel struct {
	Current  int64    `json:"current"`
	Capacity int64    `json:"capacity"`
	Consumed consumed `json:"consumed"`
}

type consumed struct {
	Amount    int64  `json:"amount"`
	Timestamp string `json:"timestamp"`
}

type module struct {
	Symbol       string             `json:"symbol"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Capacity     *int64             `json:"capacity,omitempty"`
	Requirements moduleRequirements `json:"requirements"`
	Range        *int64             `json:"range,omitempty"`
}

type moduleRequirements struct {
	Crew  int64 `json:"crew"`
	Power int64 `json:"power"`
	Slots int64 `json:"slots"`
}

type mount struct {
	Symbol       string             `json:"symbol"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Strength     int64              `json:"strength"`
	Requirements engineRequirements `json:"requirements"`
	Deposits     []string           `json:"deposits,omitempty"`
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

type reactor struct {
	Symbol       string              `json:"symbol"`
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Condition    int64               `json:"condition"`
	PowerOutput  int64               `json:"powerOutput"`
	Requirements reactorRequirements `json:"requirements"`
}

type reactorRequirements struct {
	Crew int64 `json:"crew"`
}

type registration struct {
	Name          string `json:"name"`
	FactionSymbol string `json:"factionSymbol"`
	Role          string `json:"role"`
}
