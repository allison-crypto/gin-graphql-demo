package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/imdario/mergo"
)

type PaginationResp struct {
	Data     interface{}
	Page     int
	PageSize int
	Total    int
}

func graphqlField(
	typ graphql.Type,
	description string,
	resolve graphql.FieldResolveFn,
	otherProps ...*graphql.Field,
) *graphql.Field {
	filed := &graphql.Field{
		Type:        typ,
		Resolve:     resolve,
		Description: description,
	}

	if len(otherProps) != 0 && otherProps[0] != nil {
		mergo.Merge(filed, *otherProps[0])
	}
	return filed
}

var pagenationFiledArgs = graphql.FieldConfigArgument{
	"page": &graphql.ArgumentConfig{
		Type:        graphql.Int,
		Description: "The start page of request",
	},
	"pageSize": &graphql.ArgumentConfig{
		Type:        graphql.String,
		Description: "The number of records in each page",
	},
}

func pagenationType(
	name string,
	dataType graphql.Type,
	otherProps ...*graphql.Field,
) *graphql.Object {
	pagenationFiled := graphql.Field{
		Args: pagenationFiledArgs,
	}
	if len(otherProps) != 0 && otherProps[0] != nil {
		mergo.Merge(&pagenationFiled, otherProps[0])
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name: name,
		Fields: graphql.Fields{
			"data": graphqlField(
				graphql.NewList(dataType),
				"the records",
				nil,
				&pagenationFiled,
			),
			"page": graphqlField(
				graphql.Int,
				"The start page of request",
				nil,
			),
			"pageSize": graphqlField(
				graphql.Int,
				"The number of records in each page",
				nil,
			),
			"total": graphqlField(
				graphql.Int,
				"The total number of the records",
				nil,
			),
		},
	})
}
