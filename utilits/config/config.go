package config

// Config ...
type Config struct {
	Server struct {
		Host string `ini:"host"`
		Port string `ini:"port"`
	} `ini:"server"`
	Database struct {
		URL  string `ini:"url"`
	} `ini:"database"`
	Logger struct {
		Level     string `ini:"level"`
		ShortLvl  bool   `ini:"short_lvl"`
		MaxSize   int    `ini:"max_size"`
		MaxAge    int    `ini:"max_age"`
		MaxBackup int    `ini:"max_backup"`
		Compress  bool   `ini:"compress"`
		Localtime bool   `ini:"localtime"`
		LogFile   string `ini:"log_file"`
	} `ini:"logger"`
}

// New ...
func New() *Config {
	//TODO: make default config parameters
	return &Config{}
}
