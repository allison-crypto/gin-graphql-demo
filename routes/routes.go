package routes

import (
	"github.com/gin-gonic/gin"
)

// MountRoutes mount routes for entitle project
func MountRoutes(srv *gin.Engine) {
	checkRoute := srv.Group("/check")
	mountCheckRoutes(checkRoute)

	graphqlGroup := srv.Group("/graphql")
	mountGraphqlRoutes(graphqlGroup)
}
