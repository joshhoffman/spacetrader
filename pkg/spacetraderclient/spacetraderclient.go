package spacetraderclient

import (
	"encoding/json"
	"fmt"

	"github.com/joshhoffman/spacetrader/pkg/client"
)

type request interface{}

type RegisterRequest struct {
	Symbol  string `json:"symbol"`
	Faction string `json:"faction"`
}

type SpaceTraderClient struct {
	client   client.Client
	Callsign string
}

func (s SpaceTraderClient) init() {
	s.client = client.Client{Url: "https://api.spacetraders.io/v2/"}
}

func (s SpaceTraderClient) Register(faction string) {
	register := &RegisterRequest{Symbol: s.Callsign, Faction: "COSMIC"}
	registerBytes, err := getBodyBytes(register)
	if err != nil {
		return
	}
	resBody, err := s.client.MakePostRequest("register", registerBytes)
	println(resBody)
}

func getBodyBytes(r request) ([]byte, error) {
	request, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error")
		return nil, nil
	}

	return request, nil
}
