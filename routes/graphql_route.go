package routes

import (
	"gin-graphql-demo/graphql/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
)

func mountGraphqlRoutes(rg *gin.RouterGroup) {
	graphqlSchema := schemas.Schema()
	introspection.AddIntrospectionToSchema(graphqlSchema)
	rg.POST("/", gin.WrapH(graphql.HTTPHandler(graphqlSchema)))

	rg.StaticFS("/playground", http.Dir("./graphql/playground"))
}
