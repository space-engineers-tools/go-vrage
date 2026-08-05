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
		RemoteSecurityKey: "PLEASE USE A STRONG SECURITY KEY",
	})
	panicIfError(err)

	ping, err := vrageClient.Server.Ping()
	panicIfError(err)

	fmt.Println("The server uses api version:", ping.Meta.APIVersion)
	fmt.Println(ping.Data.Result)
	fmt.Println(ping.Meta.QueryTime)
}
