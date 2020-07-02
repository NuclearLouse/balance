package logger

import (
	"balance/utilits/config"

	"github.com/orandin/lumberjackrus"
	"github.com/sirupsen/logrus"
)

// New ...
func New(cfg *config.Config) (*logrus.Logger, error) {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	loglev, err := logrus.ParseLevel(cfg.Logger.Level)
	if err == nil {
		log.SetLevel(loglev)
	}
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		PadLevelText:    true,
		ForceColors:     true,
	})

	hookLog, err := hooking(cfg, &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		DisableColors:   true,
		DisableQuote:    true,
	})
	if err != nil {
		return nil, err
	}
	log.AddHook(hookLog)
	return log, nil

}

func hooking(cfg *config.Config, format logrus.Formatter) (*lumberjackrus.Hook, error) {
	logger := &lumberjackrus.LogFile{
		Filename:   cfg.Logger.LogFile,
		MaxSize:    cfg.Logger.MaxSize,
		MaxBackups: cfg.Logger.MaxBackup,
		MaxAge:     cfg.Logger.MaxAge,
		Compress:   cfg.Logger.Compress,
		LocalTime:  cfg.Logger.Localtime,
	}

	hook, err := lumberjackrus.NewHook(
		logger,
		logrus.TraceLevel,
		format,
		// &easy.Formatter{
		// 	TimestampFormat: "2006-01-02 15:04:05.000",
		// 	LogFormat:       "%lvl% [%time%] ▶ %msg% remote_addr=%remote_addr% request_id=%request_id%\n",
		// 	ShortLevel:      cfg.Logger.ShortLvl,
		// },
		&lumberjackrus.LogFileOpts{
			logrus.ErrorLevel: logger,
			logrus.FatalLevel: logger,
			logrus.PanicLevel: logger,
		},
	)
	if err != nil {
		return nil, err
	}
	return hook, nil
}
