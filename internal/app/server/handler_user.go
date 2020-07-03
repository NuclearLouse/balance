package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"balance/internal/app/models"
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
			switch r.Header.Get("Content-Type") {
			case "application/json":
				s.respond(w, r, http.StatusOK, data)
			default:
				s.tmpl.ExecuteTemplate(w, "register.html", data)
			}
			return
		}

		switch r.Header.Get("Content-Type") {
		case "application/json":
			u := &models.User{}
			if err := json.NewDecoder(r.Body).Decode(u); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}
			defer r.Body.Close()
			if err := s.dbStore.UserRepository().Create(r.Context(), u); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			u.Sanitize()
			u.Status = true
			s.respond(w, r, http.StatusCreated, u)
			return
		case "application/xml":
			// Respond with XML
		default:
			u := &models.User{
				Email:          r.PostFormValue("email"),
				Password:       r.PostFormValue("password"),
				RepeatPassword: r.PostFormValue("passRepeat"),
				Username:       r.PostFormValue("username"),
			}
			if err := s.dbStore.UserRepository().Create(r.Context(), u); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				data["Error"] = true
				data["ErrorTitle"] = "Регистрация не прошла!"
				data["ErrorMessage"] = err.Error()
				s.tmpl.ExecuteTemplate(w, "register.html", data)
				return
			}

			message := fmt.Sprintf("Зарегистрирован новый пользователь %s", u.Email)
			data["Success"] = true
			data["SuccessTitle"] = "Пользователь успешно создан"
			data["SuccessMessage"] = message
			s.logger.Info(message)
			w.WriteHeader(http.StatusCreated)
			s.tmpl.ExecuteTemplate(w, "register.html", data)
		}
	}
}
