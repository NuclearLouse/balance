package server

import (
	"balance/internal/app/store/teststore"
	"testing"

	"github.com/NuclearLouse/logging"
	"github.com/gorilla/sessions"
)

func testServer(t *testing.T) *server {
	cfg := NewConfig()
	cfgLog := logging.NewConfig()
	cfgLog.Level = "error"
	log, _ := logging.New(cfgLog)
	s := newServer(
		teststore.New(),
		cfg,
		log,
		sessions.NewCookieStore([]byte("secret")),
	)
	s.tmpl = templateFiles("C:\\Users\\android\\go\\balance\\templates")
	return s
}
