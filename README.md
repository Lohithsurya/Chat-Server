# Chat Server

## Overview

This project implements a simple chat server in Go, allowing multiple clients to connect and communicate with each other in real-time. Each client is assigned a name and can send messages that are broadcasted to all other connected clients.

## Features

- Multiple client support: The server can handle multiple clients connecting and chatting simultaneously.
- Real-time broadcasting: Messages from clients are immediately broadcasted to all other connected clients.
- Graceful client disconnect: The server handles client disconnections gracefully and cleans up resources.

## Project Structure

- `Client`: Represents a single client connection. Manages reading and writing messages to the client.
- `Server`: Manages client connections, broadcasting messages, and handling client registration and unregistration.
- `main`: The entry point of the application. Sets up the server and handles incoming client connections.

## How It Works

1. **Server Initialization**: The server is initialized and starts listening for incoming TCP connections on port 8080.
2. **Client Connection**: When a new client connects, they are prompted to enter their name.
3. **Message Handling**: 
   - The client's messages are read and broadcasted to all other connected clients.
   - Messages are formatted to include the client's name.
4. **Client Registration and Unregistration**: 
   - Clients are registered when they connect and unregistered when they disconnect.
   - The server ensures proper cleanup of resources when a client disconnects.

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Running the Server

1. **Clone the repository**:
   ```sh
   git clone https://github.com/your-username/chat-server.git
   cd chat-server


