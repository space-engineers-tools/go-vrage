package main

import (
	"fmt"

	"github.com/space-engineers-tools/go-vrage"
)

func panicIfError(err error) {
	if nil != err {
		panic(err)
	}
}

func main() {
	vrageClient, err := vrage.NewClient(vrage.ClientConfig{
		UseHTTPS:          false,
		RemoteApiIP:       "127.0.0.1",
		RemoteApiPort:     11202,
		RemoteSecurityKey: "PLEASE USE A STRONG SECURITY KEY AND STORE IT IN A .env FILE",
	})
	panicIfError(err)

	// get current players
	playersRes, err := vrageClient.Session.Players()
	panicIfError(err)
	fmt.Println("Players:")
	for _, player := range playersRes.Data.Players {
		fmt.Printf("- %+v\n", player)
	}
	/*
		Example output:

		Players:
		・{SteamID:76561199082847890 DisplayName:EchterTimo FactionName:Python Haters FactionTag:PYH PromoteLevel:5 Ping:42}
	*/
}
