package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/joshhoffman/spacetrader/pkg/spacetraderclient"
)

const serverPort = 8080

type Server struct{}

func (s Server) StartServer() {
	callsign := os.Getenv("CALLSIGN")
	token := os.Getenv("TOKEN")

	c := spacetraderclient.SpaceTraderClient{Callsign: callsign, Token: token}
	c.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/ships", func(w http.ResponseWriter, r *http.Request) {
		resp, err := c.Ships()
		if err != nil {
			os.Exit(1)
		}
		ret, err := (*resp).Marshal()
		if err != nil {
			os.Exit(1)
		}
		w.Write(ret)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s \n", r.Method)
	})
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error running http server: %s\n", err)
		}
	}
}
