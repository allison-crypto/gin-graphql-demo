package schemas

import "github.com/graphql-go/graphql"

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"allison": allisonFiled,
		"samson":  samsonFiled,
	},
})

var allisonFiled = graphqlField(
	graphql.NewNonNull(graphql.String),
	"allison online",
	func(p graphql.ResolveParams) (interface{}, error) {
		return "developing", nil
	},
)

var samsonFiled = graphqlField(
	graphql.String,
	"I am samson",
	func(p graphql.ResolveParams) (interface{}, error) {
		return "woohoo", nil
	},
	&graphql.Field{
		DeprecationReason: "going to the party",
	},
)
