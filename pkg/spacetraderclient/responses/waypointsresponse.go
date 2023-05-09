package responses

import (
	"encoding/json"

	"github.com/joshhoffman/spacetrader/pkg/client"
)

func (o WaypointsResponse) Unmarshal(data []byte) (client.MarshalableObject, error) {
	var r WaypointsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r WaypointsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WaypointsResponse struct {
	Data []waypointsResponseDatum `json:"data"`
	Meta Meta                     `json:"meta"`
}

type waypointsResponseDatum struct {
	SystemSymbol SystemSymbol             `json:"systemSymbol"`
	Symbol       string                   `json:"symbol"`
	Type         string                   `json:"type"`
	X            int64                    `json:"x"`
	Y            int64                    `json:"y"`
	Orbitals     []Faction                `json:"orbitals"`
	Traits       []waypointsResponseTrait `json:"traits"`
	Chart        chart                    `json:"chart"`
	Faction      Faction                  `json:"faction"`
}

type chart struct {
	SubmittedBy SubmittedBy `json:"submittedBy"`
	SubmittedOn string      `json:"submittedOn"`
}

type waypointsResponseTrait struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
