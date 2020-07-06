package server

import (
	"balance/internal/app/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (s *server) handleClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value(ctxKeyUser).(*models.User)
		data := map[string]interface{}{
			"title": "Создание клиента",
			"user":  u.Username,
			"admin": u.Admin,
		}

		if r.Method == "GET" {
			s.tmpl.ExecuteTemplate(w, "create_client.html", data)
			return
		}
		tc := r.PostFormValue("type")
		if tc == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			data["Error"] = true
			data["ErrorTitle"] = "Клиент не создан"
			data["ErrorMessage"] = "поля не могут быть пустыми, не выбран тип клиента"
			s.tmpl.ExecuteTemplate(w, "create_client.html", data)
			return
		}
		c := models.Client{}
		c.Status = true
		typeClient, err := strconv.Atoi(tc)
		if err != nil {
			typeClient = 666
			c.Status = false
		}
		c.Type = typeClient

		mkup := r.PostFormValue("markup")
		mkup = strings.Replace(mkup, ",", ".", 1)
		markup, err := strconv.ParseFloat(mkup, 64)
		if err != nil {
			markup = 0
		}
		c.Markup = 1 + markup/100

		c.Name = r.PostFormValue("name")
		c.Comment = r.PostFormValue("comment")
		c.User = u.ID
		if err := s.store.Repository().CreateClient(r.Context(), c); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			data["Error"] = true
			data["ErrorTitle"] = "Клиент не создан"
			data["ErrorMessage"] = "внутренняя ошибка сервера"
			s.logger.Errorln("создание нового клиента", err)
			s.tmpl.ExecuteTemplate(w, "create_client.html", data)
			return
		}
		message := fmt.Sprintf("Создан новый клиент %s", c)
		data["Success"] = true
		data["SuccessTitle"] = "Клиент успешно создан"
		data["SuccessMessage"] = message
		s.logger.Infoln(message, "для юзера", u.Username)
		w.WriteHeader(http.StatusCreated)
		s.tmpl.ExecuteTemplate(w, "create_client.html", data)
	}
}
