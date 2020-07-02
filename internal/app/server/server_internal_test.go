package server

import (
	"balance/internal/app/store/teststore"
	"balance/utilits/config"
	"balance/utilits/logger"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
)

func TestServer_HandleRegister(t *testing.T) {
	cfg := config.New()
	log,_ := logger.New(cfg)
	s := newServer(
		teststore.New(),
		cfg,
		log,
		sessions.NewCookieStore([]byte("secret")),				
	)

	testCases := []struct{
		name string
		payload interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email": "user@example.org",
				"password": "password",
				"repeat_password": "password",
			},
		},
	}

	for _, tc := range testCases {
		rec := httptest.NewRecorder()

		http.NewRequest("POST","/register",)
	}
}