package http

import (
	"context"
	"encoding/json"
	"net"
	"net/http"

	"github.com/YnMann/chat_backend/internal/chat"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

// Client management
type ClientManager struct {
	//The client map stores and manages all long connection clients, online is TRUE, and those who are not there are FALSE
	clients map[*Client]bool
	//Web side MESSAGE we use Broadcast to receive, and finally distribute it to all clients
	broadcast chan []byte
	//Newly created long connection client
	register chan *Client
	//Newly canceled long connection client
	unregister chan *Client
}

// Client
type Client struct {
	//User ID
	uID string
	//Connected socket
	socket *websocket.Conn
	//Message
	send chan []byte
	//ctx
	ctx context.Context
}

// Create a client manager
var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

// Will formatting Message into JSON
type Message struct {
	//Message Struct
	SenderID    string `json:"sender_id,omitempty"`
	RecipientID string `json:"recipient_id,omitempty"`
	Content     string `json:"content,omitempty"`
	ServerIP    string `json:"serverIp,omitempty"`
	SenderIP    string `json:"senderIp,omitempty"`
}

// For waiting upload data in db
var wait = make(chan struct{})

func (manager *ClientManager) start(uc chat.UseCase) {
	for {
		select {
		//If there is a new connection access, pass the connection to conn through the channel
		case conn := <-manager.register:
			// Set the client connection to true
			manager.clients[conn] = true
			// Change on db - isOnline
			go func() {
				err := uc.SetUserOnlineStatus(conn.ctx, conn.uID, true)

				if err != nil {
					jsonMessage, _ := json.Marshal(
						&Message{
							Content:  "Set online status err",
							ServerIP: LocalIp(),
							SenderIP: conn.socket.RemoteAddr().String(),
						},
					)
					manager.send(jsonMessage, conn)
				}
				wait <- struct{}{}
			}()
			<-wait
			//Format the message of returning to the successful connection JSON
			jsonMessage, _ := json.Marshal(
				&Message{
					Content:  "/A new socket has connected. ",
					ServerIP: LocalIp(),
					SenderIP: conn.socket.RemoteAddr().String(),
				},
			)
			//Call the client's send method and send messages
			manager.send(jsonMessage, conn)

		//If the connection is disconnected
		case conn := <-manager.unregister:
			//Determine the state of the connection, if it is true, turn off Send and delete the value of connecting client
			if _, ok := manager.clients[conn]; ok {
				go func() {
					err := uc.SetUserOnlineStatus(conn.ctx, conn.uID, false)

					if err != nil {
						jsonMessage, _ := json.Marshal(
							&Message{
								Content:  "Disconnect err",
								ServerIP: LocalIp(),
								SenderIP: conn.socket.RemoteAddr().String(),
							},
						)
						manager.send(jsonMessage, conn)
					}
					wait <- struct{}{}
				}()
				<-wait

				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(
					&Message{
						Content:  "/A socket has disconnected. ",
						ServerIP: LocalIp(),
						SenderIP: conn.socket.RemoteAddr().String(),
					},
				)
				manager.send(jsonMessage, conn)
			}
			//broadcast
		case message := <-manager.broadcast:
			//Traversing the client that has been connected, send the message to them
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

// Define the send method of client management
func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		//Send messages not to the shielded connection
		if conn != ignore {
			conn.send <- message
		}
	}
}

// Define the read method of the client structure
func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		_ = c.socket.Close()
	}()

	for {
		//Read message
		_, message, err := c.socket.ReadMessage()
		//If there is an error message, cancel this connection and then close it
		if err != nil {
			manager.unregister <- c
			_ = c.socket.Close()

			break
		}
		//If there is no error message, put the information in Broadcast
		jsonMessage, _ := json.Marshal(
			&Message{
				SenderID: c.uID,
				Content:  string(message),
				ServerIP: LocalIp(),
				SenderIP: c.socket.RemoteAddr().String(),
			},
		)
		manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		_ = c.socket.Close()
	}()

	for {
		select {
		//Read the message from send
		case message, ok := <-c.send:
			//If there is no message
			if !ok {
				_ = c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//Write it if there is news and send it to the web side
			_ = c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 1024 * 1024,
	WriteBufferSize: 1024 * 1024 * 1024,
	//Solving cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(c *gin.Context) {
	// Upgrade the HTTP protocol to the websocket protocol
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	var msg Message
	if err := conn.ReadJSON(&msg); err != nil {
		http.Error(c.Writer, "Invalid message format", http.StatusBadRequest)
		return
	}
	userID := msg.SenderID

	// Every connection will open a new client, client.id generates through UUID to ensure that each time it is different
	client := &Client{
		uID:    userID,
		socket: conn,
		send:   make(chan []byte),
		ctx:    c,
	}

	// Register a new link
	manager.register <- client

	// Start the message to collect the news from the web side
	go client.read()
	// Start the corporation to return the message to the web side
	go client.write()
}

func healthHandler(c *gin.Context) {
	_, _ = c.Writer.Write([]byte("ok"))
}

func LocalIp() string {
	address, _ := net.InterfaceAddrs()
	var ip = viper.GetString("host")
	for _, address := range address {
		if ipAddress, ok := address.(*net.IPNet); ok && !ipAddress.IP.IsLoopback() {
			if ipAddress.IP.To4() != nil {
				ip = ipAddress.IP.String()
			}
		}
	}
	return ip
}
