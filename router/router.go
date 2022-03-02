package router

import (
	"github.com/salleh/bp-ekyc-callback-server-demo/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	callback := app.Group("/callback", logger.New())

	idOnlyCallback := callback.Group("/idonly")
	idOnlyCallback.Post("/", handler.IDOnly)

	idFaceCallback := callback.Group("/idface")
	idFaceCallback.Post("/", handler.IDFace)

	idLiveCallback := callback.Group("/idlive")
	idLiveCallback.Post("/", handler.IDLive)

	idReflectCallback := callback.Group("/idreflect")
	idReflectCallback.Post("/", handler.IDReflect)
}
