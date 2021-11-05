package routes

import (
	"encoding/json"
	"gin-graphql-demo/graphql/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type graphBody struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func mountGraphqlRoutes(rg *gin.RouterGroup) {
	rg.POST("/", func(ctx *gin.Context) {
		var body graphBody
		if err := json.NewDecoder(ctx.Request.Body).Decode(&body); err != nil {
			ctx.Status(400)
			return
		}
		result := graphql.Do(graphql.Params{
			Context:        ctx.Request.Context(),
			Schema:         schemas.TestSchema,
			RequestString:  body.Query,
			VariableValues: body.Variables,
			OperationName:  body.Operation,
		})
		ctx.JSON(http.StatusOK, result)
	})

	rg.StaticFS("/playground", http.Dir("./graphql/playground"))
}
