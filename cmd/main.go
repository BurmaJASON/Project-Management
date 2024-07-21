package main

import (
	"log"
	"project-management/internal/app"
)

// @title Project Management API
// @version 1.0
// @description This is a sample server for project management.
// @termsOfService http://swagger.io/terms/

// @contact.name Saw Ye Htet
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	application, err := app.NewApp()

	if err != nil {
		log.Fatalf("failed to start the application: %v", err)
	}
	application.Run()
}
