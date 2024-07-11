package types

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/crontastic/app/config"
	"github.com/mrrizkin/crontastic/system/database"
	"github.com/mrrizkin/crontastic/system/session"
	"github.com/mrrizkin/crontastic/third_party/logger"
	"github.com/mrrizkin/crontastic/third_party/whatsapp"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Debug   string      `json:"debug"`
	Data    interface{} `json:"data"`
}

type App struct {
	*fiber.App

	Logger   *logger.Logger
	Database *database.Database
	Config   *config.Config
	Session  *session.Session
	WhatsApp *whatsapp.Whatsapp
}
