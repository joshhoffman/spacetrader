package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshhoffman/spacetrader/pkg/server"
	"github.com/joshhoffman/spacetrader/pkg/spacetraderclient"
)

type RegisterRequest struct {
	Symbol  string `json:"symbol"`
	Faction string `json:"faction"`
}

func main() {
	_, present := os.LookupEnv("DEBUG")
	if !present {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Error reading .env")
			return
		}
	}

	callsign := os.Getenv("CALLSIGN")
	token := os.Getenv("TOKEN")

	c := spacetraderclient.SpaceTraderClient{Callsign: callsign, Token: token}
	c.Init()
	ships, _ := c.Ships()
	// c.Ships()
	shipId := ships.Data[0].Registration.Name
	c.Orbit(shipId)

	s := server.Server{}
	s.StartServer()
}
