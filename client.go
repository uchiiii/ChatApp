package main

import (
	"github.com/gorilla/websocket"
	"time"
)

type client struct {
	socket   *websocket.Conn
	send     chan *message
	room     *room
	userData map[string]interface{}
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c) //interfaceをroomに付け加える.
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send { //ここで無限ループしてほしいのに、してくれなさそう. -> capasityに空きがない場合と, emptyの場合はロックされる.
		err := c.socket.WriteJSON(msg)
		if err != nil {
			break
		}
	}
}
