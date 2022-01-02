package main

import (
	"context"
	"log"
	"os"

	app ".app-nubank-routine-check-email"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	ctx := context.Background()
	if err := funcframework.RegisterEventFunctionContext(ctx, "/", app.ReadEmail); err != nil {
		log.Fatalf("app.ReadEmail: %v\n", err)
	}
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("app.Start: %v\n", err)
	}
}
