package server

import (
	"balance/internal/app/store"
	"balance/internal/app/store/sqlstore"
	"balance/utilits/config"
	"balance/utilits/database"
	"balance/utilits/logger"
	"context"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

const (
	sessionName        = "sess_balance"
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

type ctxKey int8

type server struct {
	dbStore  store.Store
	router   *mux.Router
	logger   *logrus.Logger
	sesStore sessions.Store
	tmpl     *template.Template
}

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Start ...
func Start() error {
	cfgFile, ok := os.LookupEnv("CFG_FILE")
	if !ok {
		return errors.New("config file not found")
	}

	cfg, err := ini.Load(cfgFile)
	if err != nil {
		return err
	}
	config := config.New()
	if err := cfg.MapTo(config); err != nil {
		return err
	}

	log, err := logger.New(config)
	if err != nil {
		return err
	}

	ctx := context.Background()
	db, err := database.New(ctx, config.Database.URL)
	if err != nil {
		return err
	}
	defer db.Close(ctx)

	databaseStore := sqlstore.New(db)

	sessionKey := uuid.New().String()
	sessionStore := sessions.NewCookieStore([]byte(sessionKey))
	sessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}
	defer func() {
		sessionStore.Options.MaxAge = -1
	}()
	srv := newServer(databaseStore, config, log, sessionStore)
	bindAddr := config.Server.Host + ":" + config.Server.Port
	srv.logger.Infoln("Start server at", bindAddr)
	return http.ListenAndServe(bindAddr, srv)
}

func newServer(databaseStore store.Store, cfg *config.Config, logger *logrus.Logger, sessionStore sessions.Store) *server {
	pattern := filepath.Join("templates", "*.html")
	s := &server{
		router:   mux.NewRouter(),
		logger:   logger,
		dbStore:  databaseStore,
		sesStore: sessionStore,
		tmpl:     template.Must(template.ParseGlob(pattern)),
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()

		rw := &responseWriter{w, http.StatusOK}

		next.ServeHTTP(rw, r)

		logger.Infof(
			"complited with CODE:[%d] %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
		)
	})
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
