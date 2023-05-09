package responses

import (
	"encoding/json"

	"github.com/joshhoffman/spacetrader/pkg/client"
)

func (o FactionsResponse) Unmarshal(data []byte) (client.MarshalableObject, error) {
	var r FactionsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (o FactionsResponse) Marshal() ([]byte, error) {
	return json.Marshal(o)
}

type FactionsResponse struct {
	Data []factionsResponseData `json:"data"`
	Meta Meta                   `json:"meta"`
}

type factionsResponseData struct {
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Headquarters string  `json:"headquarters"`
	Traits       []trait `json:"traits"`
}

type trait struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
