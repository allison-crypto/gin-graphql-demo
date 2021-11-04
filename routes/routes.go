package routes

import (
	"github.com/gin-gonic/gin"
)

// MountRoutes mount routes for entitle project
func MountRoutes(srv *gin.Engine) {
	checkRoute := srv.Group("/check")
	mountCheckRoutes(checkRoute)

	v1 := srv.Group("/v1/")

	graphqlGroup := v1.Group("/graphql")
	mountGraphqlRoutes(graphqlGroup)
}
