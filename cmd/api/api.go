package api

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shonginger/GoPhonebook/Lehem/services/phonebook"
)

type APIServer struct {
	serverAddress string
	db            *sql.DB
}

func NewAPIServer(serverAddress string, db *sql.DB) *APIServer {
	return &APIServer{
		serverAddress: serverAddress,
		db:            db,
	}
}

func (s *APIServer) Run() error {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	_, err := http.Get("https://golang.org/")
	if err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	contactStore := phonebook.NewStore(s.db)
	phonebook.NewHandler(contactStore).RegisterRoutes(subrouter)

	log.Println("listening on ", s.serverAddress)
	return http.ListenAndServe(s.serverAddress, router)
}
