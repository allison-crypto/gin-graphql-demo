package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

type RoleType int32
type User struct {
	Id        int
	FirstName string
	LastName  string
	Role      RoleType
}

func registerUser(schema *schemabuilder.Schema) {
	object := schema.Object("User", User{})
	object.Key("id")

	object.FieldFunc("fullName", func(u *User) string {
		return u.FirstName + " " + u.LastName
	})
}

type Args struct {
	Role *RoleType
}

func registerQuery(schema *schemabuilder.Schema) {
	object := schema.Query()

	var tmp RoleType
	schema.Enum(tmp, map[string]RoleType{
		"user":          RoleType(1),
		"manager":       RoleType(2),
		"administrator": RoleType(3),
	})

	userListRet := func(ctx context.Context, args Args) ([]*User, error) {
		return []*User{
			{
				Id:        1,
				FirstName: "Bob",
				LastName:  "Johnson",
				Role:      RoleType(1),
			},
			{
				Id:        2,
				FirstName: "Chloe",
				LastName:  "Kim",
				Role:      RoleType(1),
			},
		}, nil
	}

	object.FieldFunc("users", userListRet)

	object.FieldFunc("usersConnection", userListRet, schemabuilder.Paginated)

}

func registerMutation(schema *schemabuilder.Schema) {
	object := schema.Mutation()

	object.FieldFunc("echo", func(ctx context.Context, args struct{ Text string }) (string, error) {
		return args.Text, nil
	})

	object.FieldFunc("echoEnum", func(ctx context.Context, args struct {
		EnumField RoleType
	}) (RoleType, error) {
		return args.EnumField, nil
	})
}

func Schema() *graphql.Schema {
	schema := schemabuilder.NewSchema()

	registerUser(schema)
	registerQuery(schema)
	registerMutation(schema)

	return schema.MustBuild()
}

func mountGraphqlRoutes(rg *gin.RouterGroup) {
	introspection.AddIntrospectionToSchema(Schema())
	rg.POST("/", gin.WrapH(graphql.HTTPHandler(Schema())))
}
