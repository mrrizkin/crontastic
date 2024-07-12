package server

import (
	"fmt"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"

	_ "github.com/joho/godotenv/autoload"

	"github.com/mrrizkin/crontastic/app/config"
	"github.com/mrrizkin/crontastic/system/session"
	"github.com/mrrizkin/crontastic/third_party/logger"
)

type Server struct {
	*fiber.App

	config *config.Config
}

type Site struct {
	Menus []MenuItem
}

type MenuItem struct {
	Name string
	URL  string
	Icon string
}

func New(config *config.Config, logger *logger.Logger, session *session.Session) *Server {
	// engine.AddFunc("site", func() Site {
	// 	return Site{
	// 		Menus: []MenuItem{
	// 			{Name: "Dashboard", URL: "/admin/dashboard", Icon: "house"},
	// 			{Name: "Auto Tasks", URL: "/admin/auto-tasks", Icon: "calendar-check"},
	// 			{Name: "Manual Tasks", URL: "/admin/manual-tasks", Icon: "clipboard-list"},
	// 			{Name: "Users", URL: "/admin/users", Icon: "users-round"},
	// 			{Name: "Settings", URL: "/admin/settings", Icon: "settings"},
	// 			{Name: "Documentation", URL: "/docs", Icon: "book-open-text"},
	// 		},
	// 	}
	// })

	app := fiber.New(fiber.Config{
		Prefork: config.PREFORK,
		AppName: config.APP_NAME,
	})

	csrfConfig := csrf.Config{
		KeyLookup:         "header:" + csrf.HeaderName,
		CookieName:        "cronstactic_csrf_token",
		CookieSameSite:    "Lax",
		CookieSecure:      false,
		CookieSessionOnly: true,
		CookieHTTPOnly:    true,
		SingleUseToken:    true,
		Expiration:        1 * time.Hour,
		KeyGenerator:      utils.UUIDv4,
		ErrorHandler:      csrf.ConfigDefault.ErrorHandler,
		Extractor:         csrf.CsrfFromHeader(csrf.HeaderName),
		Session:           session.Store,
		SessionKey:        "fiber.csrf.token",
		HandlerContextKey: "fiber.csrf.handler",
	}

	app.Use(csrf.New(csrfConfig))

	app.Static("/", "public")

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger.Logger,
	}))
	app.Use(requestid.New())
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(idempotency.New())

	return &Server{
		App:    app,
		config: config,
	}
}

func (s *Server) Serve() error {
	return s.Listen(fmt.Sprintf(":%d", s.config.PORT))
}

func (s *Server) Stop() error {
	return s.Shutdown()
}
