package main

import(
	//"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"log"
)


//websocket code--->  (client message echo from server)



var upgrader= websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
    http.HandleFunc("/ws", handleWebSocket)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("upgrade error:", err)
        return
    }
    defer conn.Close()

    for {message, 
        messageType, meerr := conn.ReadMessage()
        if err != nil {
            log.Println("read error:", err)
            break
        }
        log.Printf("recv: %s", message)

        // Echo the message back to the client
        // err = conn.WriteMessage(messageType, message)
        // if err != nil { 
        //     log.Println("write error:", err)
        //     break
        // }
    }
}


//websocket code --->chatting room for multiple users



// Define a new type to represent a WebSocket client
// type client struct {
//     conn *websocket.Conn // WebSocket connection
//     send chan []byte     // Channel for outgoing messages
// }

// // Define a new type to represent the WebSocket chat room
// type room struct {
//     clients map[*client]bool // Connected clients
//     broadcast chan []byte   // Channel for incoming messages
//     join chan *client       // Channel for new clients joining the room
//     leave chan *client      // Channel for clients leaving the room
// }

// func (r *room) run() {
//     for {
//         select {
//         case client := <-r.join:
//             r.clients[client] = true
//         case client := <-r.leave:
//             if _, ok := r.clients[client]; ok {
//                 delete(r.clients, client)
//                 close(client.send)
//             }
//         case message := <-r.broadcast:
//             for client := range r.clients {
//                 select {
//                 case client.send <- message:
//                 default:
//                     delete(r.clients, client)
//                     close(client.send)
//                 }
//             }
//         }
//     }
// }

// func (c *client) read(r *room) {
//     defer func() {
//         r.leave <- c
//         c.conn.Close()
//     }()
//     for {
//         _, message, err := c.conn.ReadMessage()
//         if err != nil {
//             break
//         }
//         r.broadcast <- message
//     }
// }

// func (c *client) write() {
//     defer c.conn.Close()
//     for message := range c.send {
//         err := c.conn.WriteMessage(websocket.TextMessage, message)
//         if err != nil {
//             break
//         }
//     }
// }

// func wsHandler(w http.ResponseWriter, r *http.Request) {
//     // Upgrade HTTP request to WebSocket
//     conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
//     if err != nil {
//         return
//     }

//     // Create a new client and add it to the room
//     client := &client{conn: conn, send: make(chan []byte, 256)}
//     room.join<-client

//     // Start separate goroutines to handle reading and writing from the WebSocket connection
//     go client.write()
//     go client.read(room)
// }

// // var room = &room{
// //     clients: make(map[*client]bool),
// //     broadcast: make(chan []byte, 256),
// //     join: make(chan *client),
// //     leave: make(chan *client),
// // }

// func main() {
//     go room.run()

//     http.HandleFunc("/ws", wsHandler)
//     err := http.ListenAndServe(":8080", nil)
//     if err != nil {
//         fmt.Println("ListenAndServe error:", err)
//     }
// }