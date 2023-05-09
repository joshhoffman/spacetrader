package spacetraderclient

import (
	"encoding/json"
	"fmt"

	"github.com/joshhoffman/spacetrader/pkg/client"
	"github.com/joshhoffman/spacetrader/pkg/spacetraderclient/responses"
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
	resp, err := s.client.MakeGetRequest("my/agent", new(responses.AgentResponse))
	if err != nil {
		fmt.Printf("Error in agent %s\n", err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func (s SpaceTraderClient) Factions() {
	resp, err := s.client.MakeGetRequest("factions", new(responses.FactionsResponse))
	if err != nil {
		fmt.Printf("Error in factions %s\n", err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func (s SpaceTraderClient) Contracts() {
	resp, err := s.client.MakeGetRequest("my/contracts", new(responses.ContractResponse))
	if err != nil {
		fmt.Printf("Error in contracts %s\n", err)
		return
	}
	fmt.Printf("%+v\n", resp)
	// contractResponse, err := responses.UnmarshalWelcome(resp)
}

func (s SpaceTraderClient) Systems() {
	resp, err := s.client.MakeGetRequest("systems", new(responses.SystemsResponse))
	if err != nil {
		fmt.Printf("Error in contracts %s\n", err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func (s SpaceTraderClient) Waypoints(systemId string) {
	urlPath := fmt.Sprintf("systems/%s/waypoints", systemId)
	resp, err := s.client.MakeGetRequest(urlPath, new(responses.WaypointsResponse))
	if err != nil {
		fmt.Printf("Error in contracts %s\n", err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func (s SpaceTraderClient) Ships() (*responses.ShipsResponse, error) {
	resp, err := s.client.MakeGetRequest("my/ships", new(responses.ShipsResponse))
	if err != nil {
		fmt.Printf("Error in contracts %s\n", err)
		return nil, err
	}
	ret := resp.(responses.ShipsResponse)
	fmt.Printf("%+v\n", resp)

	return &ret, nil
}

func getBodyBytes(r *RegisterRequest) ([]byte, error) {
	request, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error")
		return nil, nil
	}

	return request, nil
}
