package msg2api

import (
	"errors"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	upgradeTimeout = 10 * time.Second

	deviceAPIProtocolV1 = "v2.device.msg"
	userAPIProtocolV3   = "v5.user.msg"
)

var errProtocolNegotiationFailed = errors.New("protocol negotiation failed")

type apiBase struct {
	socket *socketWrapper
}

func (b *apiBase) Close() {
	b.socket.Close(websocket.CloseGoingAway, "")
}

func initAPIBaseFromSocket(conn *websocket.Conn) (*apiBase, error) {
	if conn.Subprotocol() == "" {
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseProtocolError, ""))
		conn.Close()
		return nil, errProtocolNegotiationFailed
	}

	conn.SetReadLimit(4096)

	return &apiBase{
		socket: wrapWebsocket(conn),
	}, nil
}

func initAPIBaseFromHTTP(w http.ResponseWriter, r *http.Request, protocols []string) (*apiBase, error) {
	upgrader := websocket.Upgrader{
		HandshakeTimeout: upgradeTimeout,
		Subprotocols:     protocols,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	return initAPIBaseFromSocket(conn)
}
