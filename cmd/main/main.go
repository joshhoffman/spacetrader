package main

import (
	"os"

	"github.com/joshhoffman/spacetrader/pkg/spacetraderclient"
)

type RegisterRequest struct {
	Symbol  string `json:"symbol"`
	Faction string `json:"faction"`
}

func main() {
	callsign := os.Getenv("CALLSIGN")
	c := spacetraderclient.SpaceTraderClient{Callsign: callsign}
	c.Register("COSMIC")
}
