package server

import (
	"balance/internal/app/store/teststore"
	"balance/utilits/config"
	"balance/utilits/logger"
	"testing"

	"github.com/gorilla/sessions"
)

func testServer(t *testing.T) *server {
	cfg := config.New()
	log, _ := logger.New(cfg)

	s := newServer(
		teststore.New(),
		cfg,
		log,
		sessions.NewCookieStore([]byte("secret")),
	)
	s.tmpl = templateFiles("C:\\Users\\android\\go\\balance\\templates")
	return s
}
