package main

import (
	appnubankroutinecheckemail "app-nubank-routine-check-email"
	"context"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	ctx := context.Background()
	if err := funcframework.RegisterEventFunctionContext(ctx, "/", appnubankroutinecheckemail.ReadEmail); err != nil {
		log.Fatalf("appnubankroutinecheckemail.ReadEmail: %v\n", err)
	}
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
