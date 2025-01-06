package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Global variables
var db *sql.DB // This variable stores the database connection object
var tmpl *template.Template // This variable stores the parsed templates

// This function is called before the main function.
func init() {
	// Parse and load all templates before starting the server
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

// Creates a new connection to the database and stores it in the db variable
func initDB() {
	var err error

	// Open the database connection
	db, err = sql.Open("mysql", "root:toor@(127.0.0.1:3306)/testdb?parseTime=true") 
	if err != nil { log.Fatal(err) }

	// Check if the connection is successful
	if err = db.Ping(); err != nil { log.Fatal(err) }
}

func main() {
	initDB()
	defer db.Close()

	gRouter := mux.NewRouter()

	gRouter.HandleFunc("/", HomeHandler)
	
	http.ListenAndServe(":8080", gRouter)
}

// This handler function sends a response to the root URL ("/")
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil { 
		http.Error(w, "Error while loading templates: " + err.Error(), http.StatusInternalServerError) 
	}
}