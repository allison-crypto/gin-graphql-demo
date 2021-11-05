package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/imdario/mergo"
)

var (
	TestSchema graphql.Schema
)

func init() {
	TestSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
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
