package responses

import (
	"encoding/json"

	"github.com/joshhoffman/spacetrader/pkg/client"
)

func (o ContractResponse) Unmarshal(data []byte) (client.MarshalableObject, error) {
	var r ContractResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (o ContractResponse) Marshal() ([]byte, error) {
	return json.Marshal(o)
}

type ContractResponse struct {
	Data []contractResponseData `json:"data"`
	Meta Meta                   `json:"meta"`
}

type contractResponseData struct {
	ID            string `json:"id"`
	FactionSymbol string `json:"factionSymbol"`
	Type          string `json:"type"`
	Terms         terms  `json:"terms"`
	Accepted      bool   `json:"accepted"`
	Fulfilled     bool   `json:"fulfilled"`
	Expiration    string `json:"expiration"`
}

type terms struct {
	Deadline string    `json:"deadline"`
	Payment  payment   `json:"payment"`
	Deliver  []deliver `json:"deliver"`
}

type deliver struct {
	TradeSymbol       string `json:"tradeSymbol"`
	DestinationSymbol string `json:"destinationSymbol"`
	UnitsRequired     int64  `json:"unitsRequired"`
	UnitsFulfilled    int64  `json:"unitsFulfilled"`
}

type payment struct {
	OnAccepted  int64 `json:"onAccepted"`
	OnFulfilled int64 `json:"onFulfilled"`
}
