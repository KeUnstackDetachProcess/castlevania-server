package main

import (
	"cv-server/pkg/api"
	"cv-server/pkg/logging"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	err    error
	router *mux.Router
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {

	logging.Motd()

	// Using 64 instead of 32 cause stupid strconv doesn't have an uint32 formatter
	var port = flag.Uint64("p", 80, "Port used for the web service")

	router = mux.NewRouter()
	router.Use(commonMiddleware)

	// Load environment variables (db configuration)
	logging.Info("Loading .env file")
	err = godotenv.Load()
	if err != nil {
		logging.Error("Error loading .env file")
	}

	// Get database configuration from environment
	DBUSER := os.Getenv("DBUSER")
	DBPASS := os.Getenv("DBPASS")
	DBADDRESS := os.Getenv("DBADDRESS")
	DBPORT := os.Getenv("DBPORT")
	DBNAME := os.Getenv("DBNAME")

	// Initialize MySQL Database connection
	logging.Info("Estabilishing MySQL connection")
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUSER, DBPASS, DBADDRESS, DBPORT, DBNAME)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.Error(err.Error())
	}

	// Migrate tables
	logging.Info("Setupping routes handling")
	db.AutoMigrate()

	// Configure webserver routes
	logging.Info("Setupping routes handling")

	// Handle node users
	router.HandleFunc("/api/user/create", api.UserCreate).Methods("POST")
	router.HandleFunc("/api/user/update", api.UserCreate).Methods("POST")

	// Handle all P2P & GC requests
	router.HandleFunc("/api/E2EE/p2p", api.UserCreate).Methods("POST")
	router.HandleFunc("/api/E2EE/chat", api.UserCreate).Methods("POST")

	router.HandleFunc("/api/E2EE/chat/users/list", api.UserCreate).Methods("POST")
	router.HandleFunc("/api/E2EE/chat/users/ban", api.UserCreate).Methods("POST")
	router.HandleFunc("/api/E2EE/chat/users/add", api.UserCreate).Methods("POST")

	router.HandleFunc("/api/p2p", api.UserCreate).Methods("POST")
	router.HandleFunc("/api/chat", api.UserCreate).Methods("POST")

	// Start HTTP listener
	logging.Info("Starting http webserver")
	if err := http.ListenAndServe(":"+strconv.FormatUint(*port, 10), router); err != nil {
		logging.Error(err.Error())
	}

	os.Exit(1)
}
