package main

import (
	"oauth2bin/oauth2/server"
	"os"
)

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := server.NewOA2Server(port, "config/flowParams.json", "config/ratePolicies.csv")
	server.Start()
}
