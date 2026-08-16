package main

import (
	"fmt"
	"time"

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

	// read chat history
	chatRes, err := vrageClient.Session.ChatGet()
	panicIfError(err)
	fmt.Println("Messages:")
	for _, msg := range chatRes.Data.Messages {
		sendTime, err := msg.TimestampAsTime()
		panicIfError(err)

		fmt.Printf("%s | %s: %s\n", sendTime.Format(time.TimeOnly), msg.DisplayName, msg.Content)
	}

	// send message
	_, err = vrageClient.Session.ChatSend("hello from go!")
	panicIfError(err)
}
