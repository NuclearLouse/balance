package server

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func (s *server) configureRouter() {
	fs := http.FileServer(http.Dir("./static"))
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/", s.handleLogin()).Methods("GET","POST")
	s.router.HandleFunc("/logout", s.handleLogout()).Methods("GET")
	s.router.HandleFunc("/register", s.handleRegister()).Methods("GET","POST")

	private := s.router.PathPrefix("/u").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/journal", s.handleJournal()).Methods("GET")
	private.HandleFunc("/create_client", s.handleClient()).Methods("GET","POST")
}
