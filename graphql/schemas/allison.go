package schemas

import "github.com/graphql-go/graphql"

var allisonFiled = graphqlField(
	graphql.NewNonNull(graphql.String),
	"allison online",
	func(p graphql.ResolveParams) (interface{}, error) {
		return "developing", nil
	},
)
