package apiv1

import (
	"phrases-rest-api/api/v1/phrases"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		phrases.ApplyRoutes(v1)
	}
}
