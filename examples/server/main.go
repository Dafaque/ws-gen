package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Dafaque/ws-gen/examples/generated/api"
	"github.com/Dafaque/ws-gen/examples/generated/iface"
	"github.com/Dafaque/ws-gen/examples/generated/model"
	"github.com/Dafaque/ws-gen/examples/generated/server"
)

type handler struct {
	api.UnimplementedMessageHandler
}

func (h handler) OnChatEvent(ctx context.Context, msg model.ChatEvent, sender *api.MessageSender) error {
	fmt.Printf("server got chat event: %s : %f\n", msg.Event, msg.TestSnakeCaseConvertor)
	if msg.ID == 0 {
		return errors.New("TEST server closes on error")
	}
	message := "Hello there!"
	var f float64 = 0.1
	return sender.SendTextMessage(
		*model.NewTextMessage(0, &message, []int64{1, 2, 3}, []*float64{&f, nil}))
}
func (h handler) OnDisconnected(code int, reason string) {
	fmt.Printf("Disconnected client: code=%d, reason=%s", code, reason)
}

func main() {
	handler := server.NewHandler(func() api.MessageHandler { return handler{} }, iface.DefaultCoder{}, log.Default())
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatalln(server.ListenAndServe())
}
