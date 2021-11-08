package schemas

import (
	"github.com/graphql-go/graphql"
)

type testData struct {
	Name string
	Task string
}

var paginationFiled = graphqlField(
	pagenationType("pagination", developerType),
	"pagination recrods",
	paginationReslove,
)

func paginationReslove(p graphql.ResolveParams) (interface{}, error) {
	return PaginationResp{
		Data: []testData{
			{Name: "allison", Task: "api"},
			{Name: "samson", Task: "lambda"},
		},
		Page:     1,
		PageSize: 2,
		Total:    2,
	}, nil
}

var developerType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "developer",
	Description: "A character in the Star Wars Trilogy",
	Fields: graphql.Fields{
		"name": graphqlField(
			graphql.String,
			"the name of developer", nil,
		),
		"task": graphqlField(
			graphql.String,
			"the task of developer", nil,
		),
	},
})
