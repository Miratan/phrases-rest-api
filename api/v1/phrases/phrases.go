package phrases

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	phrases := r.Group("/phrases")
	{
		phrases.POST("/", create)
		phrases.GET("/", list)
		phrases.GET("/:id", read)
		phrases.DELETE("/:id", remove)
		phrases.PATCH("/:id", update)
	}
}
