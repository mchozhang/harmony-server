/**
 * wrapper of graphql query and mutation executor and extractor,
 * define the graphql objects
 */
package graphql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"harmony-server/handler/dynamodb"
)

// cell graphql
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

// Level graphql
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
				// 2D list
				Type: graphql.NewList(graphql.NewList(cellType)),
			},
		},
	})

// query
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
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

// GraphQL root graphql
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

/**
 * execute the graphql query
 */
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
