package resolvers

import (
	"todo-app/db"
	"todo-app/models"

	"github.com/graphql-go/graphql"
)

var todoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"title":     &graphql.Field{Type: graphql.String},
		"completed": &graphql.Field{Type: graphql.Boolean},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"todos": &graphql.Field{
			Type: graphql.NewList(todoType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var todos []models.Todo
				err := db.DB.Select(&todos, "SELECT * FROM todos")
				return todos, err
			},
		},
	},
})

var addTodoMutation = &graphql.Field{
	Type: todoType,
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		title := p.Args["title"].(string)
		todo := models.Todo{Title: title}
		_, err := db.DB.NamedExec("INSERT INTO todos (title) VALUES (:title)", &todo)
		return todo, err
	},
}

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"addTodo": addTodoMutation,
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
