package routes

import (
	"github.com/gin-gonic/gin"
	client "github.com/sofyan48/svc_auth/src/app/v1/api/client/controller"
	health "github.com/sofyan48/svc_auth/src/app/v1/api/health/controller"
	token "github.com/sofyan48/svc_auth/src/app/v1/api/token/controller"
	"github.com/sofyan48/svc_auth/src/middleware"
)

// VERSION ...
const VERSION = "v1"

// V1RouterLoader types
type V1RouterLoader struct {
	Middleware middleware.DefaultMiddleware
	Health     health.HealthControllerInterface
	Client     client.ClientControllerInterface
	Token      token.TokenControllerInterface
}

// V1RouterLoaderHandler ...
func V1RouterLoaderHandler() *V1RouterLoader {
	return &V1RouterLoader{
		Health: health.HealthControllerHandler(),
		Client: client.ClientControllerHandler(),
		Token:  token.TokenControllerHandler(),
	}
}

// V1RouterLoaderInterface ...
type V1RouterLoaderInterface interface {
	V1Routes(router *gin.Engine)
}

// V1Routes Params
// @router: gin.Engine
func (rLoader *V1RouterLoader) V1Routes(router *gin.Engine) {
	rLoader.initDocs(router)
	rLoader.initHealth(router)
	rLoader.initClient(router)
	rLoader.initToken(router)
}