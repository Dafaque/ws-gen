// Code generated by wsgen. DO NOT EDIT.
package api
import (
    "time"
    
    "github.com/dafaque/ws-gen/examples/generated/iface"
    "github.com/dafaque/ws-gen/examples/generated/model"
    
    "github.com/gorilla/websocket" 
)

type MessageSender struct {
    conn *websocket.Conn
    coder iface.Coder
}

func NewMessageSender(conn *websocket.Conn, coder iface.Coder) *MessageSender {
    return &MessageSender{conn: conn, coder: coder}
}

func (c *MessageSender) SendBinary(data []byte) error {
    return c.conn.WriteMessage(websocket.BinaryMessage, data)
}

func (c *MessageSender) send(msg interface{}) error {
    data, err := c.coder.Marshal(msg)
    if err != nil {
        return err
    }
    return c.SendBinary(data)
}

func (c *MessageSender) Ping() error {
    return c.conn.WriteControl(
        websocket.PingMessage,
        nil,//@todo
        time.Now().Add(5*time.Second),//@todo
    )
}
func (c *MessageSender) Pong() error {
    return c.conn.WriteControl(
        websocket.PongMessage,
        nil,//@todo
        time.Now().Add(5*time.Second),//@todo
    )
}
// Sends TextMessage message
// To instance TextMessage message use model.NewTextMessage method
func (c *MessageSender) SendTextMessage(msg model.TextMessage) error {
    return c.send(msg)
}
// Sends ChatEvent message
// To instance ChatEvent message use model.NewChatEvent method
func (c *MessageSender) SendChatEvent(msg model.ChatEvent) error {
    return c.send(msg)
}