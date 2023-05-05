package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Dafaque/ws-gen/examples/generated/api"
	"github.com/Dafaque/ws-gen/examples/generated/client"
	"github.com/Dafaque/ws-gen/examples/generated/iface"
	"github.com/Dafaque/ws-gen/examples/generated/model"
)

type handler struct {
	api.UnimplementedMessageHandler
}

func (h handler) OnTextMessage(ctx context.Context, msg model.TextMessage, _ *api.MessageSender) error {
	fmt.Printf("client got message: %s\n", *msg.Content)
	if msg.ID == 0 {
		return errors.New("TEST client closes on error")
	}
	return nil
}
func (h handler) OnDisconnected(code int, reason string) {
	fmt.Printf("Disconnected from server: code=%d, reason=%s", code, reason)
}

func main() {
	var h handler
	cl, err := client.NewClient(
		"ws://localhost:8080",
		&model.InitParams{
			ChatId: "123",
		},
		h,
		iface.DefaultCoder{},
		log.Default(),
	)
	if err != nil {
		panic(err)
	}
	cl.SendChatEvent(*model.NewChatEvent(123, model.EventEntered, 2.000, false, nil, nil, nil))
	cl.Ping()
	cl.Pong()
	done := make(chan struct{})
	time.AfterFunc(5*time.Second, func() {
		cl.Shutdown()
		done <- struct{}{}
	})
	<-done
}
