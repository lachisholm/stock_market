// This file is using websocket to connect and read the data

package main

import (
	"context" // Import the context package for managing context
	"fmt"     // Import the fmt package for formatted I/O
	"log"     // Import the log package for logging errors and information
	"time"    // Import the time package for handling time-related operations

	"nhooyr.io/websocket"        // Import the websocket package for WebSocket connections
	"nhooyr.io/websocket/wsjson" // Import the wsjson package for JSON over WebSockets
)

const websocketURL = "wss://ws-api.iextrading.com/1.0/tops" // Define the WebSocket URL

// startWebSocket function starts a WebSocket connection and reads data
func startWebSocket() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute) // Create a context with timeout
	defer cancel()                                                        // Ensure the context is cancelled when function exits

	conn, _, err := websocket.Dial(ctx, websocketURL, nil) // Establish a WebSocket connection
	if err != nil {
		log.Fatal(err) // Log and exit if connection fails
	}
	defer conn.Close(websocket.StatusInternalError, "Internal error") // Ensure the connection is closed on exit

	symbol := "AAPL"                                              // Define the stock symbol
	wsjson.Write(ctx, conn, map[string]string{"symbols": symbol}) // Send the stock symbol over WebSocket

	// Continuously read data from WebSocket
	for {
		var data map[string]interface{}      // Define a variable to hold the data
		err := wsjson.Read(ctx, conn, &data) // Read data from WebSocket
		if err != nil {
			log.Fatal(err) // Log and exit if read fails
		}
		fmt.Println(data) // Print the received data
	}
}
