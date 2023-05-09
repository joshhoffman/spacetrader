package responses

import (
	"encoding/json"

	"github.com/joshhoffman/spacetrader/pkg/client"
)

func (o SystemsResponse) Unmarshal(data []byte) (client.MarshalableObject, error) {
	var r SystemsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (o SystemsResponse) Marshal() ([]byte, error) {
	return json.Marshal(o)
}

type SystemsResponse struct {
	Data []systemsResponseDatum `json:"data"`
	Meta Meta                   `json:"meta"`
}

type systemsResponseDatum struct {
	Symbol       string        `json:"symbol"`
	SectorSymbol SectorSymbol  `json:"sectorSymbol"`
	Type         string        `json:"type"`
	X            int64         `json:"x"`
	Y            int64         `json:"y"`
	Waypoints    []waypoint    `json:"waypoints"`
	Factions     []interface{} `json:"factions"`
}

type waypoint struct {
	Symbol string `json:"symbol"`
	Type   Type   `json:"type"`
	X      int64  `json:"x"`
	Y      int64  `json:"y"`
}

type SectorSymbol string

const (
	X1 SectorSymbol = "X1"
)

type Type string

const (
	AsteroidField  Type = "ASTEROID_FIELD"
	GasGiant       Type = "GAS_GIANT"
	JumpGate       Type = "JUMP_GATE"
	Moon           Type = "MOON"
	OrbitalStation Type = "ORBITAL_STATION"
	Planet         Type = "PLANET"
)
