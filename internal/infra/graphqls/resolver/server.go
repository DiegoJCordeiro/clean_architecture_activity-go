package resolver

import (
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/infra/graphqls/models"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func CreateGraphQLServer(resolver *Resolver) http.Handler {
	orderType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.String},
			"customer_id": &graphql.Field{Type: graphql.String},
			"price":       &graphql.Field{Type: graphql.Float},
			"tax":         &graphql.Field{Type: graphql.Float},
			"final_price": &graphql.Field{Type: graphql.Float},
			"created_at":  &graphql.Field{Type: graphql.String},
			"updated_at":  &graphql.Field{Type: graphql.String},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"listOrders": &graphql.Field{
				Type: graphql.NewList(orderType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.Query().ListOrders(p.Context)
				},
			},
		},
	})

	createOrderInputType := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "CreateOrderInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"customer_id": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"price":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Float)},
			"tax":         &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Float)},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createOrder": &graphql.Field{
				Type: orderType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(createOrderInputType)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})
					return resolver.Mutation().CreateOrder(p.Context, toCreateOrderInput(input))
				},
			},
		},
	})

	// Schema
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}

func toCreateOrderInput(input map[string]interface{}) models.CreateOrderInput {
	return models.CreateOrderInput{
		CustomerID: input["customer_id"].(string),
		Price:      input["price"].(float64),
		Tax:        input["tax"].(float64),
	}
}
