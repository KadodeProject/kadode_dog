package route

import (
	"github.com/gin-gonic/gin"
	"backend/action"
)

func GetRouter(r *gin.Engine) *gin.Engine {
	r.GET("/", action.IndexAction)
	return r
}
