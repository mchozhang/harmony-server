// Package graphql
// wrapper of graphql query and mutation executor and extractor,
// define the graphql objects
package graphql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"harmony-server/handler/dynamodb"
)

// the schema of a Cell object
var cellType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Cell",
		Fields: graphql.Fields{
			"targetRow": &graphql.Field{
				Type: graphql.Int,
			},
			"steps": &graphql.Field{
				Type: graphql.Int,
			},
			"row": &graphql.Field{
				Type: graphql.Int,
			},
			"col": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})

// the schema of a Level object
var levelType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Level",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"size": &graphql.Field{
				Type: graphql.Int,
			},
			"colors": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"cells": &graphql.Field{
				// 2D list of cell object
				Type: graphql.NewList(graphql.NewList(cellType)),
			},
		},
	})

// the schema of the input Cell object of the mutation
var cellInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CellInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"targetRow": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"steps": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"row": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"col": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	})

// the schema of the input Level object of the mutation
var levelInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "LevelInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"level": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"size": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"colors": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
			},
			"cells": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.NewList(cellInputType)),
			},
		},
	})

// query
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// query the game level data
			"level": &graphql.Field{
				Type: levelType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					levelId, isOk := p.Args["id"].(int)

					if isOk {
						return dynamodb.GetLevel(levelId)
					}
					return nil, nil
				},
			},
		},
	},
)

// mutation type
var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createLevel": &graphql.Field{
			Type: levelType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: levelInputType,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				levelInput, isOk := p.Args["input"]
				if isOk {
					return dynamodb.CreateLevel(levelInput)
				}
				return nil, nil
			},
		},
	},
},
)

// GraphQL root
// define what queries and mutations are supported
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

// ExecuteQuery
// execute the graphql query or mutation to the dynamoDB
func ExecuteQuery(query string, variables map[string]interface{}, operationName string) *graphql.Result {
	// parse the quest body to acquire query parameters
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		VariableValues: variables,
		RequestString:  query,
		OperationName:  operationName,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("Error: %v\n", result.Errors)
	}
	return result
}
