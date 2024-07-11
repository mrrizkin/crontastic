package config

import (
	"fmt"
	"regexp"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	APP_NAME string
	ENV      string
	PORT     int
	PREFORK  bool

	LOG_LEVEL      string
	LOG_CONSOLE    bool
	LOG_FILE       bool
	LOG_DIR        string
	LOG_MAX_SIZE   int
	LOG_MAX_AGE    int
	LOG_MAX_BACKUP int
	LOG_JSON       bool

	DB_DRIVER   string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_SSLMODE  string

	SESSION_DRIVER string

	WHATSAPP_SENDER_TOKEN string
	WHATSAPP_SERVICE_URL  string
}

var validAPP_NAME = regexp.MustCompile(`^[a-zA-Z_0-9]+$`)

func New() *Config {
	envAPP_NAME, _ := envStr("APP_NAME", "boot")

	if !validAPP_NAME.MatchString(*envAPP_NAME) {
		panic(fmt.Errorf("APP_NAME is not valid"))
	}

	envENV, _ := envStr("ENV", "development")
	envPORT, _ := envInt("PORT", 3000)
	envPREFORK, _ := envBool("PREFORK", false)

	envLOG_LEVEL, _ := envStr("LOG_LEVEL", "debug")
	envLOG_CONSOLE, _ := envBool("LOG_CONSOLE", true)
	envLOG_FILE, _ := envBool("LOG_FILE", true)
	envLOG_DIR, _ := envStr("LOG_DIR", "./storage/log")
	envLOG_MAX_SIZE, _ := envInt("LOG_MAX_SIZE", 50)
	envLOG_MAX_AGE, _ := envInt("LOG_MAX_AGE", 7)
	envLOG_MAX_BACKUP, _ := envInt("LOG_MAX_BACKUP", 20)
	envLOG_JSON, _ := envBool("LOG_JSON", true)

	envDBDriver, _ := envStr("DB_DRIVER", "sqlite")
	envDBHost, _ := envStr("DB_HOST", "./storage/db.sqlite")
	envDBPort, _ := envInt("DB_PORT", 5432)
	envDBName, _ := envStr("DB_NAME", "boot")
	envDBUsername, _ := envStr("DB_USERNAME", "boot")
	envDBPassword, _ := envStr("DB_PASSWORD", "boot")
	envDBSSLMode, _ := envStr("DB_SSLMODE", "disable")

	envSESSION_DRIVER, _ := envStr("SESSION_DRIVER", "sqlite")

	envWHATSAPP_SENDER_TOKEN, err := envStr("WHATSAPP_SENDER_TOKEN")
	if err != nil {
		panic(fmt.Errorf("WHATSAPP_SENDER_TOKEN is required"))
	}

	envWHATSAPP_SERVICE_URL, err := envStr("WHATSAPP_SERVICE_URL")
	if err != nil {
		panic(fmt.Errorf("WHATSAPP_SERVICE_URL is required"))
	}

	return &Config{
		APP_NAME: *envAPP_NAME,
		ENV:      *envENV,
		PORT:     *envPORT,
		PREFORK:  *envPREFORK,

		LOG_LEVEL:      *envLOG_LEVEL,
		LOG_CONSOLE:    *envLOG_CONSOLE,
		LOG_FILE:       *envLOG_FILE,
		LOG_DIR:        *envLOG_DIR,
		LOG_MAX_SIZE:   *envLOG_MAX_SIZE,
		LOG_MAX_AGE:    *envLOG_MAX_AGE,
		LOG_MAX_BACKUP: *envLOG_MAX_BACKUP,
		LOG_JSON:       *envLOG_JSON,

		DB_DRIVER:   *envDBDriver,
		DB_HOST:     *envDBHost,
		DB_PORT:     *envDBPort,
		DB_NAME:     *envDBName,
		DB_USERNAME: *envDBUsername,
		DB_PASSWORD: *envDBPassword,
		DB_SSLMODE:  *envDBSSLMode,

		SESSION_DRIVER: *envSESSION_DRIVER,

		WHATSAPP_SENDER_TOKEN: *envWHATSAPP_SENDER_TOKEN,
		WHATSAPP_SERVICE_URL:  *envWHATSAPP_SERVICE_URL,
	}
}
