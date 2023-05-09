package responses

import (
	"encoding/json"

	"github.com/joshhoffman/spacetrader/pkg/client"
)

func (o AgentResponse) Unmarshal(data []byte) (client.MarshalableObject, error) {
	var r AgentResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (o AgentResponse) Marshal() ([]byte, error) {
	return json.Marshal(o)
}

type AgentResponse struct {
	Data AgentResponseData `json:"data"`
}

type AgentResponseData struct {
	AccountID    string `json:"accountId"`
	Symbol       string `json:"symbol"`
	Headquarters string `json:"headquarters"`
	Credits      int64  `json:"credits"`
}
