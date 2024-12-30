package main

import (
	"database/sql"
	"gokit-crud-app/pkg/endpoint"
	"gokit-crud-app/pkg/service"
	"gokit-crud-app/pkg/transport"
	"log"
	"net/http"
	"os"

	// "github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file (only for local development)
	// if err := godotenv.Load(); err != nil {
	//     log.Fatal("Error loading .env file")
	// }
	connStr, errBool := os.LookupEnv("DATABASE_URL")
	if !errBool {
		log.Fatal("DATABASE_URL environment variable is required")
	}
	port, errBool := os.LookupEnv("PORT")
	if !errBool {
		port = "8080"
	}
	log.Println("Connecting to database...")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	svc := service.NewTodoService(db)
	endpoints := endpoint.MakeEndpoints(svc)
	handler := transport.NewHTTPHandler(endpoints)
	log.Println("Backend service running at :", port)
	http.ListenAndServe(":"+port, handler)
}
