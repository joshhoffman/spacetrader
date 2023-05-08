package spacetraderclient

import (
	"encoding/json"
	"fmt"

	"github.com/joshhoffman/spacetrader/pkg/client"
)

type requestInterface interface{}

type RegisterRequest struct {
	Symbol  string `json:"symbol"`
	Faction string `json:"faction"`
}

type SpaceTraderClient struct {
	client   client.Client
	Callsign string
	Token    string
}

func (s *SpaceTraderClient) Init() {
	s.client = client.MakeClient("https://api.spacetraders.io/v2/", s.Token)
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

func (s SpaceTraderClient) Agent() {
	s.client.MakeGetRequest("my/agent")
}

func getBodyBytes(r *RegisterRequest) ([]byte, error) {
	request, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error")
		return nil, nil
	}

	return request, nil
}
