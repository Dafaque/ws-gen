package main

import (
	"context"
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

func (h handler) OnTextMessage(ctx context.Context, msg model.TextMessage, sender *api.MessageSender) error {
	fmt.Printf("server got message: %s\n", msg.Content)
	return sender.SendTextMessage(*model.NewTextMessage(321, "Hello there!"))
}
func (h handler) OnChatEvent(ctx context.Context, msg model.ChatEvent, sender *api.MessageSender) error {
	fmt.Printf("server got chat event: %s : %f\n", msg.Event, msg.TestSnakeCaseConvertor)
	return sender.SendChatEvent(msg)
}

func main() {
	handler := server.NewHandler(handler{}, iface.DefaultCoder{}, log.Default())
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatalln(server.ListenAndServe())
}
