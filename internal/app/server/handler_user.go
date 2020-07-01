package server

import (
	"net/http"
)


func (s *server) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"title": "Вход",
		}
		if r.Method == "GET" {
			s.tmpl.ExecuteTemplate(w, "login.html", data)
			return
		}
		s.tmpl.ExecuteTemplate(w, "login.html", data)
	}
}

func (s *server) handleRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"title": "Регистрация",
		}
		if r.Method == "GET" {
			s.tmpl.ExecuteTemplate(w, "register_user.html", data)
			return
		}
		s.tmpl.ExecuteTemplate(w, "register_user.html", data)
	}
}
