package requests

import (
	"encoding/json"

	"github.com/joshhoffman/spacetrader/pkg/client"
)

func (o FlightModeRequest) Unmarshal(data []byte) (client.MarshalableObject, error) {
	var r FlightModeRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (o FlightModeRequest) Marshal() ([]byte, error) {
	return json.Marshal(o)
}

type FlightModeRequest struct {
	FlightMode string `json:"flightMode"`
}
