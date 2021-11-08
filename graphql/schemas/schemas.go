package schemas

import (
	"github.com/graphql-go/graphql"
)

var (
	TestSchema graphql.Schema
)

func init() {
	TestSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"allison":    allisonFiled,
				"samson":     samsonFiled,
				"pagination": paginationFiled,
			},
		}),
	})
}
