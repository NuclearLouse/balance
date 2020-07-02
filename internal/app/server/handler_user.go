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
			s.tmpl.ExecuteTemplate(w, "register_user.html", data)
			return
		}
		// TODO: тут нужен свич в зависимости от типа запроса html/text или html/json
		// разные запросы должны отдавать разные переменные data		
		switch r.Header.Get("Accept") {
		case "application/json":
			// Respond with JSON
			// c.JSON(http.StatusOK, data)
		case "application/xml":
			// Respond with XML
			// c.XML(http.StatusOK, data)
		default:
			// c.HTML(status, templateName, data)
		}

		type request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &models.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.dbStore.UserRepository().Create(r.Context(), u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize() // ! нужна только при ответе json
		s.respond(w, r, http.StatusCreated, u)
		
		data["Success"] = true
		data["SuccessTitle"] = "Пользователь успешно создан"
		data["SuccessMessage"] = fmt.Sprintf("Зарегистрирован новый пользователь %s", u.Email)
		s.tmpl.ExecuteTemplate(w, "register_user.html", data)
	}
}
