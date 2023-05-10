package requests

import (
	"encoding/json"

	"github.com/joshhoffman/spacetrader/pkg/client"
)

func (o RegisterRequest) Unmarshal(data []byte) (client.MarshalableObject, error) {
	var r FlightModeRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (o RegisterRequest) Marshal() ([]byte, error) {
	return json.Marshal(o)
}

type RegisterRequest struct {
	Symbol  string `json:"symbol"`
	Faction string `json:"faction"`
}
