package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Client struct {
	conn net.Conn
	send chan string
	name string
}

func (c *Client) readMessages(server *Server) {
	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		message := scanner.Text()
		fullMessage := fmt.Sprintf("%s: %s", c.name, message)
		fmt.Printf("Broadcasting message: %s\n", fullMessage)
		server.broadcast <- fullMessage
	}
	server.unregister <- c
}

func (c *Client) writeMessages() {
	for message := range c.send { 
		
		fmt.Fprintf(c.conn,"%s\n", message)
	}
}

type Server struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan string
	mu         sync.Mutex
}

func NewServer() *Server {
	return &Server{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan string),
	}
}

func (s *Server) run() {
	for {
		select {
		case client := <-s.register:
			s.mu.Lock()
			s.clients[client] = true
			s.mu.Unlock()
			fmt.Println("New client connected")

		case client := <-s.unregister:
			s.mu.Lock()
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.send)
				fmt.Println("Client disconnected")
			}
			s.mu.Unlock()

		case message := <-s.broadcast:
			s.mu.Lock()
			for client := range s.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(s.clients, client)
				}
			}
			s.mu.Unlock()
		}
	}
}

func handleConnection(conn net.Conn, server *Server) {
	defer conn.Close()
	conn.Write([]byte("Enter your name: "))
	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading client name:", err)
		return
	}

	
	client := &Client{
		conn: conn,
		send: make(chan string),
		name: name,
	}
	fmt.Printf("clientname:%s\n", client.name)
	server.register <- client

	go client.readMessages(server)
	client.writeMessages()
}

func main() {
	server := NewServer()
	go server.run()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Chat server started on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn, server)
	}
}
