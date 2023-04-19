package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Dafaque/wsgen/examples/gen/api"
	"github.com/Dafaque/wsgen/examples/gen/client"
	"github.com/Dafaque/wsgen/examples/gen/iface"
	"github.com/Dafaque/wsgen/examples/gen/model"
)

type handler struct {
	api.UnimplementedMessageHandler
}

func (h handler) OnTextMessage(ctx context.Context, msg model.TextMessage, _ *api.MessageSender) error {
	fmt.Printf("client got message: %s\n", msg.Content)
	return nil
}

func main() {
	var h handler
	cl, err := client.NewClient("ws://localhost:8080", h, iface.DefaultCoder{}, log.Default())
	if err != nil {
		panic(err)
	}
	cl.SendTextMessage(*model.NewTextMessage(123, "Hello, world!"))

	done := make(chan struct{})
	time.AfterFunc(5*time.Second, func() {
		cl.Shutdown()
		done <- struct{}{}
	})
	<-done
}
