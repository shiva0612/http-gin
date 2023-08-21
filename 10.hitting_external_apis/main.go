package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

}

func handleRequest(c *gin.Context) {
	// Set a deadline of 5 seconds for the request
	ctx, cancel := context.WithDeadline(c.Request.Context(), time.Now().Add(5*time.Second))
	defer cancel()

	// Make the external HTTP request using the context
	_, err := http.NewRequestWithContext(ctx, "GET", "https://api.example.com", nil)
	if err != nil {
		// Handle error
	}

	// Process response and return to client
}
