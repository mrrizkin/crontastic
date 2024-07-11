package routes

import (
	"github.com/mrrizkin/crontastic/app/handlers"
	"github.com/mrrizkin/crontastic/system/types"
)

func WebRoutes(app *types.App) {
	handler := handlers.New(app)

	app.Get("/", handler.IndexPage)
	app.Get("/docs", handler.DocumentationPage)

	admin := app.Group("/admin")
	admin.Get("/dashboard", handler.DashboardPage)
	admin.Get("/auto-tasks", handler.AutoTasksPage)
	admin.Get("/manual-tasks", handler.ManualTasksPage)
	admin.Get("/users", handler.UserPage)
	admin.Get("/settings", handler.SettingsPage)

	admin.Get("/", handler.IndexPage)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/system_usage", handler.SystemUsage)

	database := v1.Group("/database")
	database.Get("/", handler.DatabaseFindAll)
	database.Get("/:id", handler.DatabaseFindByID)
	database.Post("/", handler.DatabaseCreate)
	database.Put("/:id", handler.DatabaseUpdate)
	database.Delete("/:id", handler.DatabaseDelete)

	job := v1.Group("/job")
	job.Get("/", handler.JobFindAll)
	job.Get("/:id", handler.JobFindByID)
	job.Post("/", handler.JobCreate)
	job.Put("/:id", handler.JobUpdate)
	job.Delete("/:id", handler.JobDelete)

	permission := v1.Group("/permission")
	permission.Get("/", handler.PermissionFindAll)
	permission.Get("/:id", handler.PermissionFindByID)
	permission.Post("/", handler.PermissionCreate)
	permission.Put("/:id", handler.PermissionUpdate)
	permission.Delete("/:id", handler.PermissionDelete)

	rolePermission := v1.Group("/role_permission")
	rolePermission.Get("/", handler.RolePermissionFindAll)
	rolePermission.Get("/:id", handler.RolePermissionFindByID)
	rolePermission.Post("/", handler.RolePermissionCreate)
	rolePermission.Put("/:id", handler.RolePermissionUpdate)
	rolePermission.Delete("/:id", handler.RolePermissionDelete)

	role := v1.Group("/role")
	role.Get("/", handler.RoleFindAll)
	role.Get("/:id", handler.RoleFindByID)
	role.Post("/", handler.RoleCreate)
	role.Put("/:id", handler.RoleUpdate)
	role.Delete("/:id", handler.RoleDelete)

	user := v1.Group("/user")
	user.Get("/", handler.UserFindAll)
	user.Get("/:id", handler.UserFindByID)
	user.Post("/", handler.UserCreate)
	user.Put("/:id", handler.UserUpdate)
	user.Delete("/:id", handler.UserDelete)
}
