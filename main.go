package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Database connection
var db *sql.DB

func init() {
	var err error

	// Fetch database configurations from environment variables
	dbUser := os.Getenv("DB_USER")       // MySQL username
	dbPassword := os.Getenv("DB_PASSWORD") // MySQL password
	dbName := os.Getenv("DB_NAME")       // Database name
	dbAddress := os.Getenv("DB_ADDRESS") // MySQL server address

	// Validate that all required environment variables are set
	if dbUser == "" || dbPassword == "" || dbName == "" || dbAddress == "" {
		log.Fatal("Missing required environment variables: DB_USER, DB_PASSWORD, DB_NAME, DB_ADDRESS")
	}

	// Data Source Name (DSN) for MySQL connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbAddress, dbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Database is unreachable: %v", err)
	}
	fmt.Println("Connected to the MySQL database!")
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	// Get the current time in Toronto timezone
	location, err := time.LoadLocation("America/Toronto")
	if err != nil {
		http.Error(w, "Failed to load timezone", http.StatusInternalServerError)
		log.Printf("Error loading timezone: %v", err)
		return
	}
	torontoTime := time.Now().In(location)
	currentTimeFormatted := torontoTime.Format("2006-01-02 15:04:05")

	// Insert the time into the database
	query := "INSERT INTO time_log (timestamp) VALUES (?)"
	_, err = db.Exec(query, currentTimeFormatted)
	if err != nil {
		http.Error(w, "Failed to log time to database", http.StatusInternalServerError)
		log.Printf("Error inserting time into database: %v", err)
		return
	}

	// Respond with the current time in JSON format
	response := map[string]string{
		"current_time": currentTimeFormatted,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Setup HTTP route
	http.HandleFunc("/current-time", getCurrentTime)

	// Start the server
	fmt.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
