package main

import (
	"log"
	"time"

	socketio "github.com/redt1de/go-socket.io"
)

// github.com\/redt1de\/pkg\/go-socket.io
// github.com\/googollee\/go-socket.io

// ws://cnc:8000/socket.io/?
// token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjBkNWJiZmEwLTVjYjQtNDI4ZS1iZWMxLWU1MjQxMWUzNGZlYiIsIm5hbWUiOiJjbmMiLCJpYXQiOjE2OTk4MDEyOTMsImV4cCI6MTcwMjM5MzI5M30.-beYe89MRSe1fbGUd-TnYfX7-cHjtlGKql0ZAA-RZyU
// EIO=3
// transport=websocket
// sid=-P4sIQ92f1wR2LSCAAA

func main() {
	// Simple client to talk to default-http example
	uri := "http://cnc:8000"

	// "/socket.io/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjBkNWJiZmEwLTVjYjQtNDI4ZS1iZWMxLWU1MjQxMWUzNGZlYiIsIm5hbWUiOiJjbmMiLCJpYXQiOjE2OTk4MDEyOTMsImV4cCI6MTcwMjM5MzI5M30.-beYe89MRSe1fbGUd-TnYfX7-cHjtlGKql0ZAA-RZyU"

	client, err := socketio.NewClient(uri, nil)
	if err != nil {
		panic(err)
	}

	client.OnConnect(func(s socketio.Conn) error {
		log.Println("connected:", s.ID())
		return nil
	})

	// Handle an incoming event

	// client.OnEvent("/", func(s socketio.Conn, msg string) {
	// 	log.Println("Receive Message /reply: ", "reply", msg)
	// })

	// client.OnEvent("serialport:list", func(s socketio.Conn, msg string) {
	// 	log.Println("Receive Message /reply: ", "reply", msg)
	// })

	err = client.Connect()
	if err != nil {
		panic(err)
	}

	// client.Emit("list", "")

	time.Sleep(1 * time.Second)
	err = client.Close()
	if err != nil {
		panic(err)
	}
}
