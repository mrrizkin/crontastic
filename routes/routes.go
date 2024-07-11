package routes

import (
	"github.com/mrrizkin/crontastic/system/types"
)

func Setup(app *types.App) {
	WebRoutes(app)
}
