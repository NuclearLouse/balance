package server

import (
	"context"
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
			switch r.Header.Get("Content-Type") {
			case "application/json":
				s.respond(w, r, http.StatusOK, data)
			default:
				s.tmpl.ExecuteTemplate(w, "login.html", data)
			}
			return
		}
		// create session
		var u *models.User
		code, err := func() (int, error) {
			var email, password string
			switch r.Header.Get("Content-Type") {
			case "application/json":
				type request struct {
					Email    string `json:"email"`
					Password string `json:"password"`
				}
				req := &request{}
				if err := json.NewDecoder(r.Body).Decode(req); err != nil {
					return http.StatusBadRequest, err
				}

				email, password = req.Email, req.Password
			default:
				email, password = r.PostFormValue("email"), r.PostFormValue("password")
			}
			var err error
			u, err = s.store.Repository().FindUser(r.Context(), "email", email)
			if err != nil || !u.ComparePassword(password) {
				return http.StatusUnauthorized, errIncorrectEmailOrPassword
			}

			session, err := s.sessionStore.Get(r, sessionName)
			if err != nil {
				return http.StatusInternalServerError, err
			}

			session.Values["user_id"] = u.ID
			session.Values["user_name"] = u.Username
			session.Values["admin"] = u.Admin
			session.Values["prevPage"] = 1
			if err := s.sessionStore.Save(r, w, session); err != nil {
				return http.StatusInternalServerError, err
			}

			return http.StatusOK, nil
		}()
		if err != nil {
			switch r.Header.Get("Content-Type") {
			case "application/json":
				s.error(w, r, code, err)
				return
			default:
				w.WriteHeader(code)
				data["Error"] = true
				data["ErrorTitle"] = "Вход неудачен"
				data["ErrorMessage"] = err.Error()
				s.tmpl.ExecuteTemplate(w, "login.html", data)
				return
			}

		}
		data = map[string]interface{}{
			"title":    "Журнал-главная",
			"admin":    u.Admin,
			"user":     u.Username,
			"prevPage": 1,
			"action":   "выполнен вход",
		}
		if r.Header.Get("Content-Type") == "application/json" {
			s.respond(w, r, http.StatusOK, data)
			return
		}
		//TODO: Тут нужны все данные для страницы Журнал-главная
		// .user .documents .pages
		s.logger.Debug(r.Context().Value(ctxKeyUser))
		s.tmpl.ExecuteTemplate(w, "journal.html", data)
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
			if err := s.store.Repository().CreateUser(r.Context(), u); err != nil {
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
			if err := s.store.Repository().CreateUser(r.Context(), u); err != nil {
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

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"title":      "Вход",
			"Error":      true,
			"ErrorTitle": "Необходима авторизация",
		}
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			code := http.StatusInternalServerError
			switch r.Header.Get("Content-Type") {
			case "application/json":
				s.error(w, r, code, err)
				return
			default:
				w.WriteHeader(code)
				// TODO: Нужна страница с внутренней ошибкой сервера
				data["title"] = "500-Ошибка сервера"
				data["ErrorTitle"] = "Внутренняя ошибка сервера"
				data["ErrorMessage"] = "сессия не определенна"
				s.tmpl.ExecuteTemplate(w, "500.html", data)
				return
			}
		}
		code := http.StatusUnauthorized
		id, ok := session.Values["user_id"]
		if !ok {
			switch r.Header.Get("Content-Type") {
			case "application/json":
				s.error(w, r, code, err)
				return
			default:
				w.WriteHeader(code)
				data["ErrorMessage"] = "сессия не действительна"
				s.tmpl.ExecuteTemplate(w, "login.html", data)
				return
			}
		}

		u, err := s.store.Repository().FindUser(r.Context(), "id", id.(string))
		if err != nil {
			switch r.Header.Get("Content-Type") {
			case "application/json":
				s.error(w, r, code, err)
				return
			default:
				w.WriteHeader(code)
				data["ErrorMessage"] = "пользователь не определен в системе"
				s.tmpl.ExecuteTemplate(w, "login.html", data)
				return
			}
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *server) handleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.sessionStore.Get(r, sessionName)
		session.Options.MaxAge = -1
		session.Save(r, w)
		data := map[string]interface{}{
			"action": "выполнен выход",
			"title":  "Вход",
		}
		switch r.Header.Get("Content-Type") {
		case "application/json":
			s.respond(w, r, http.StatusOK, data)
			return
		default:
			s.tmpl.ExecuteTemplate(w, "login.html", data)
		}
	}
}
