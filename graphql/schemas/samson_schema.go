package schemas

import "github.com/graphql-go/graphql"

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
