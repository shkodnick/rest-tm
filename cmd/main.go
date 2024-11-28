package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"

	router "github.com/go-game-dev/rest-tm/internal/delivery/rest"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}
	log.Printf("Current working directory: %s", dir)

	e := echo.New()

	if err := router.RegisterRoutes(e); err != nil {
		e.Logger.Fatalf("Error registering routes: %v", err)
	}

	for _, route := range e.Routes() {
		fmt.Printf("Route: %s %s\n", route.Method, route.Path)
	}

	fmt.Println("Server started at http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
