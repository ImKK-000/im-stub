package route

import (
	"im-stub/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoute(route *gin.Engine, configRoute model.ConfigRoute) {
	if configRoute.Method == http.MethodGet {
		route.GET(configRoute.Path, func(context *gin.Context) {
			context.String(http.StatusOK, configRoute.Response.(string))
		})
	}
}
