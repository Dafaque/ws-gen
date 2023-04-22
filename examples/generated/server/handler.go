// Code generated by wsgen. DO NOT EDIT.
package server

import (
    "context"
    "net/http"
    "time"
    "github.com/Dafaque/ws-gen/examples/generated/model"
    "github.com/Dafaque/ws-gen/examples/generated/iface"
    "github.com/Dafaque/ws-gen/examples/generated/mapper"
    "github.com/Dafaque/ws-gen/examples/generated/api"
    
    "github.com/gorilla/websocket"
)

type connectionHandler struct {
    ctx context.Context
    conn *websocket.Conn
    mh api.MessageHandler
    coder iface.Coder
    logger iface.Logger
    wq chan interface{}
    done bool
}

func (c *connectionHandler) GetConn() *websocket.Conn {
    return c.conn
}
func (c *connectionHandler) GetContext() context.Context {
    return c.ctx
}
func (c *connectionHandler) GetHandler() api.MessageHandler {
    return c.mh
}
func (c *connectionHandler) GetWriteChannel() chan interface{} {
    return c.wq
}
func (c *connectionHandler) GetCoder() iface.Coder {
    return c.coder
}
func (c *connectionHandler) GetLogger() iface.Logger {
    return c.logger
}

func (ch *connectionHandler) rloop() {
    for {
		if ch.done {
			break
		}
		mapper.Read(ch)
	}
}

func (ch *connectionHandler) wloop() {
    for {
		select {
		case <-ch.ctx.Done():
			return
		case iface, closed := <-ch.wq:
			if closed  || ch.done {
				return
			}
			data, err := ch.coder.Marshal(iface)
			if err != nil {
				ch.logger.Printf("wloop() Marshal: %v", err)
				return
			}
			errWrite := ch.conn.WriteMessage(websocket.BinaryMessage, data)
			if errWrite != nil {
				ch.logger.Printf("wloop() WriteMessage: %v", errWrite)
				return
			}
		}
	}
}

func (ch *connectionHandler) loop() {
    defer func() {
        close(ch.wq)
        if r := recover(); r != nil {
            ch.logger.Printf("loop() recover: %v\n", r)
            return
        }
    }()
    go ch.wloop()
    ch.rloop()
}

func (ch *connectionHandler) CloseHandler(code int, reason string) error {
    if ch.done {
        return nil
    }
	ch.done = true
	ch.mh.OnDisconnected(code, reason)
	return nil
}

type handlerMaker func() api.MessageHandler

func NewHandler(hm handlerMaker, coder iface.Coder, logger iface.Logger) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var upgrader = websocket.Upgrader{}
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        params := model.NewInitParams(r.URL.Query())
        if errValidateParams := params.Validate(); errValidateParams != nil {
            conn.WriteControl(
                websocket.CloseMessage,
                websocket.FormatCloseMessage(
                    websocket.CloseUnsupportedData,
                    errValidateParams.Error(),
                ),
                time.Now().Add(5*time.Second), //@todo write deadlines from config
            )
            return
        }
        defer conn.Close()
        mh := hm()
        connHandler := connectionHandler{
            wq: make(chan interface{}, 10),
            conn: conn,
            mh: mh,
            coder: coder,
            logger: logger,
            ctx: r.Context(),
        }

		conn.SetCloseHandler(connHandler.CloseHandler)
        connHandler.mh.Init(r.Context(), params)
        connHandler.mh.OnConnected(
            r.Context(),
            api.NewMessageSender(conn, coder),
        )
        connHandler.loop()
    }
}