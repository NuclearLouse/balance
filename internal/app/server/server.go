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
	"os/signal"
	"path/filepath"
	"syscall"
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
	store        store.Store
	router       *mux.Router
	logger       *logrus.Logger
	sessionStore sessions.Store
	tmpl         *template.Template
	serv         *http.Server
}

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func templateFiles(path string) *template.Template {
	pattern := filepath.Join(path, "*.html")
	return template.Must(template.ParseGlob(pattern))
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := database.New(ctx, config.Database.URL)
	if err != nil {
		return err
	}
	defer db.Close(ctx)

	dbStore := sqlstore.New(db)

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
	srv := newServer(dbStore, config, log, sessionStore)
	srv.tmpl = templateFiles("templates")
	srv.serv.Handler = srv
	srv.logger.Infoln("старт сервер:", srv.serv.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGQUIT,
		syscall.SIGKILL)
	go func() {
		srv.logger.Infof("получен сигнал прерывания %v", <-quit)
		if err := srv.serv.Close(); err != nil {
			errors.Wrap(err, "сервер отключился с ошибкой")
		}
	}()
	if err := srv.serv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			srv.logger.Info("сервер отключается по запросу")
		} else {
			srv.logger.Errorln("внутренняя ошибка сервера:", err)
		}
	}
	srv.logger.Info("сервер отключен")
	return nil
	// return http.ListenAndServe(bindAddr, srv)
}

func newServer(store store.Store, cfg *config.Config, logger *logrus.Logger, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logger,
		store:        store,
		sessionStore: sessionStore,
		serv: &http.Server{
			Addr: cfg.Server.Host + ":" + cfg.Server.Port,
		},
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
		logger.Infof("получен запрос %s %s", r.Method, r.RequestURI)

		start := time.Now()

		rw := &responseWriter{w, http.StatusOK}

		next.ServeHTTP(rw, r)

		logger.Infof(
			"запрос завершен [%d] %s за %v",
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
