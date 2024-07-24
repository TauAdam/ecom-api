package main

import "github.com/TauAdam/ecom-api/cmd/api"

func main() {
	server := api.NewServer(":8080", nil)
	if err := server.Run(); err != nil {
		panic(err)
	}
}
