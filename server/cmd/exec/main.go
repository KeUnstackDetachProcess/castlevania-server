package main

import (
	"cv-server/pkg/api"
	"cv-server/pkg/log"
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

	// print castelvania motd
	log.Motd()

	// using 64 instead of 32 cause stupid strconv doesn't have an uint32 formatter
	var port = flag.Uint64("p", 80, "Port used for the web service")
	if *port > 65535 {
		log.Error("Entered port is too large for the TCP protocol, please enter any port below 65535 or leave default (80)")
	}

	router = mux.NewRouter()
	router.Use(commonMiddleware)

	// load environment variables (db configuration)
	log.Info("Loading .env file")
	err = godotenv.Load()
	if err != nil {
		log.Error("An error occurred while attempting to load .env file")
	}

	// get database configuration from environment
	DBUSER := os.Getenv("DBUSER")
	DBPASS := os.Getenv("DBPASS")
	DBADDY := os.Getenv("DBADDRESS")
	DBPORT := os.Getenv("DBPORT")
	DBNAME := os.Getenv("DBNAME")

	// initialize MySQL database connection
	log.Info("Estabilishing MySQL connection")
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUSER, DBPASS, DBADDY, DBPORT, DBNAME)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err.Error())
	}

	// migrate tables
	log.Info("Setupping routes handling")
	db.AutoMigrate()

	// configure webserver routes
	log.Info("Setupping routes handling")
	// configure user routes
	api.UserSetupRoutes(&router)

	// retrieve all user friends
	router.HandleFunc("/api/user/#:tag/friend", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")
	// retrieve a specific user friend
	router.HandleFunc("/api/user/#:tag/friend/:id", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")
	// remove an user from friends
	router.HandleFunc("/api/user/#:tag/friend/:id/delete", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")

	// handle group actions
	router.HandleFunc("/api/group/create", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")
	router.HandleFunc("/api/user/delete", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")
	router.HandleFunc("/api/user/edit", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")
	router.HandleFunc("/api/user/:id", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	router.HandleFunc("/api/user/", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")

	// Start HTTP listener
	log.Info("Starting http webserver")
	if err := http.ListenAndServe(":"+strconv.FormatUint(*port, 10), router); err != nil {
		log.Error(err.Error())
	}

	os.Exit(1)
}
