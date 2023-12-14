// web_frontend/main.go

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Implement web frontend routes and logic

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Web Frontend is running on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
