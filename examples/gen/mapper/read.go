// Code generated by wsgen. DO NOT EDIT.
package mapper


import (
    "context"
    "github.com/Dafaque/wsgen/examples/gen/model"
    "github.com/Dafaque/wsgen/examples/gen/iface"
    "github.com/Dafaque/wsgen/examples/gen/api"

    "github.com/gorilla/websocket" 
)

type Handler interface {
    GetConn() *websocket.Conn
    GetContext() context.Context
    GetHandler() api.MessageHandler
    GetCoder() iface.Coder
    GetLogger() iface.Logger
	CloseHandler(int, string) error
}

func Read(h Handler) {
    _, data, errRead := h.GetConn().ReadMessage()
    if wsErr, casted := errRead.(*websocket.CloseError); casted {
		h.CloseHandler(wsErr.Code, wsErr.Text)
		return
	}
    if errRead != nil {
        h.GetLogger().Printf("Read(): %v", errRead)
        return
    }
    var msgMeta model.WSMessageMeta
    if errUnmarsh := h.GetCoder().Unmarshal(data, &msgMeta); errUnmarsh != nil {
        h.GetLogger().Printf("Read(): %v", errUnmarsh)
        return
    }
    switch msgMeta.MsgIdx {
    case model.MsgIdxTextMessage:
        var message model.TextMessage
        if err := h.GetCoder().Unmarshal(data, &message); err != nil {
            h.GetLogger().Printf("Read(): %v", err)
            return
        }
        errHandle := h.GetHandler().OnTextMessage(
            h.GetContext(),
            message,
            api.NewMessageSender(h.GetConn(), h.GetCoder()),
        )
        if errHandle != nil {
            h.GetLogger().Printf("Read()::OnTextMessage: %v", errHandle)
            return
        }
    case model.MsgIdxChatEvent:
        var message model.ChatEvent
        if err := h.GetCoder().Unmarshal(data, &message); err != nil {
            h.GetLogger().Printf("Read(): %v", err)
            return
        }
        errHandle := h.GetHandler().OnChatEvent(
            h.GetContext(),
            message,
            api.NewMessageSender(h.GetConn(), h.GetCoder()),
        )
        if errHandle != nil {
            h.GetLogger().Printf("Read()::OnChatEvent: %v", errHandle)
            return
        }
    default:
        h.GetLogger().Printf("Read()::default: recieved malformed message")
    }
    
}