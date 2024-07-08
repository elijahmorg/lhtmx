package main

import (
	"fmt"

	"github.com/elijahmorg/lhmtx/api"
)

func main() {
	err := api.GetData()
	if err != nil {
		fmt.Println("error syncing data with server")
	}
	go api.GetData()

	api.EchoStart()
}
