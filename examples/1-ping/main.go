package main

import (
	"errors"
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

	ping, err := vrageClient.Server.Ping()

	// Example error handling for common errors
	switch {
	case errors.Is(err, vrage.ErrAPIConnectionFailed):
		fmt.Println("Failed to connect to the server. Check if the server is running and reachable.")
		return
	case errors.Is(err, vrage.ErrAPIRequestTimeout):
		fmt.Println("The server did not respond in time.")
		return
	case errors.Is(err, vrage.ErrAPIInvalidSecurityKey):
		fmt.Println("The security key is invalid or missing. Please check your configuration.")
		return
	default:
		panicIfError(err)
	}

	fmt.Println("The server uses api version:", ping.Meta.APIVersion)
	fmt.Println(ping.Data.Result)
	fmt.Println(ping.Meta.QueryTime)
}
