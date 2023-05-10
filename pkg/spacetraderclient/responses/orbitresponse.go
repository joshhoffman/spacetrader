package responses

import (
	"encoding/json"

	"github.com/joshhoffman/spacetrader/pkg/client"
)

func (o OrbitResponse) Unmarshal(data []byte) (client.MarshalableObject, error) {
	var r OrbitResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (o OrbitResponse) Marshal() ([]byte, error) {
	return json.Marshal(o)
}

type OrbitResponse struct {
	Data orbitResponseData `json:"data"`
}

type orbitResponseData struct {
	Nav nav `json:"nav"`
}
