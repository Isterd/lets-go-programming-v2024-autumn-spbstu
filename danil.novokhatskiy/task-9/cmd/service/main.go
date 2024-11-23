package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/katagiriwhy/task-9/internal/config"
	"github.com/katagiriwhy/task-9/internal/storage"
)

type Contact struct {
	id    int
	name  string
	phone string
}

var testTable []Contact = []Contact{
	{
		id:    1,
		name:  "Nikita",
		phone: "89289019785",
	},
	{
		id:    2,
		name:  "",
		phone: "",
	},
}

func main() {

	pathOfCfg := config.ReadFlag()

	cfgDB := storage.Load("/home/danil/lets-go-programming-v2024-autumn-spbstu/danil.novokhatskiy/task-9/internal/storage/config.yaml")

	db, err := storage.NewStorage(cfgDB)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/contacts", getContacts(db)).Methods("GET")
	router.HandleFunc("contacts/{id}", getContact(db)).Methods("GET")
	router.HandleFunc("/users", createContact(db)).Methods("POST")
	router.HandleFunc("/users/{id}", updateContact(db)).Methods("PUT")
	router.HandleFunc("/contacts/{id}", deleteContact(db)).Methods("DELETE")

	defer db.Db.Close()

	_, err = db.Db.Exec("CREATE TABLE IF NOT EXISTS Contacts(name TEXT NOT NULL, phone TEXT NOT NULL)")
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.LoadConfig(pathOfCfg)

	fmt.Println(cfg)

}

func jsonContentTypeMiddware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getContacts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT name FROM Contacts")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		contacts := []Contact{}
		for rows.Next() {
			var contact Contact
			if err := rows.Scan(&contact.name, &contact.phone); err != nil {
				log.Fatal(err)
			}
			contacts = append(contacts, contact)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func getContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		var contact Contact

		err := db.QueryRow("SELECT name FROM Contacts WHERE id = $1", id).Scan(&contact.id, &contact.name, &contact.phone)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
