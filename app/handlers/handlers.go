package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/mrrizkin/crontastic/app/config"
	"github.com/mrrizkin/crontastic/app/domains/database"
	"github.com/mrrizkin/crontastic/app/domains/job"
	"github.com/mrrizkin/crontastic/app/domains/permission"
	"github.com/mrrizkin/crontastic/app/domains/role"
	"github.com/mrrizkin/crontastic/app/domains/role_permission"
	"github.com/mrrizkin/crontastic/app/domains/user"
	db "github.com/mrrizkin/crontastic/system/database"
	"github.com/mrrizkin/crontastic/system/session"
	"github.com/mrrizkin/crontastic/system/types"
	"github.com/mrrizkin/crontastic/third_party/logger"
	"github.com/mrrizkin/crontastic/third_party/whatsapp"
)

type Handlers struct {
	logger   *logger.Logger
	database *db.Database
	config   *config.Config
	session  *session.Session
	whatsapp *whatsapp.Whatsapp

	databaseRepo    *database.Repo
	databaseService *database.Service

	jobRepo    *job.Repo
	jobService *job.Service

	permissionRepo    *permission.Repo
	permissionService *permission.Service

	rolePermissionRepo    *role_permission.Repo
	rolePermissionService *role_permission.Service

	roleRepo    *role.Repo
	roleService *role.Service

	userRepo    *user.Repo
	userService *user.Service
}

func New(
	app *types.App,
) *Handlers {

	databaseRepo := database.NewRepo(app.Database)
	databaseService := database.NewService(databaseRepo)

	jobRepo := job.NewRepo(app.Database)
	jobService := job.NewService(jobRepo)

	permissionRepo := permission.NewRepo(app.Database)
	permissionService := permission.NewService(permissionRepo)

	rolePermissionRepo := role_permission.NewRepo(app.Database)
	rolePermissionService := role_permission.NewService(rolePermissionRepo)

	roleRepo := role.NewRepo(app.Database)
	roleService := role.NewService(roleRepo)

	userRepo := user.NewRepo(app.Database)
	userService := user.NewService(userRepo)

	return &Handlers{
		logger:   app.Logger,
		database: app.Database,
		config:   app.Config,
		session:  app.Session,
		whatsapp: app.WhatsApp,

		databaseRepo:    databaseRepo,
		databaseService: databaseService,

		jobRepo:    jobRepo,
		jobService: jobService,

		permissionRepo:    permissionRepo,
		permissionService: permissionService,

		rolePermissionRepo:    rolePermissionRepo,
		rolePermissionService: rolePermissionService,

		roleRepo:    roleRepo,
		roleService: roleService,

		userRepo:    userRepo,
		userService: userService,
	}
}

func (h *Handlers) SendJson(c *fiber.Ctx, resp interface{}, status ...int) error {
	var statusCode int

	if len(status) == 0 {
		statusCode = 200
	} else {
		statusCode = status[0]
	}

	return c.Status(statusCode).JSON(resp)
}

func (h *Handlers) Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, option := range options {
		option(componentHandler)
	}

	return adaptor.HTTPHandler(componentHandler)(c)
}
