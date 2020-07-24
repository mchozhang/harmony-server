/**
 * define the graphql schema including model and query
 */
package main

import "github.com/graphql-go/graphql"

// cell schema
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

// level schema
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
					idQuery, isOk := p.Args["id"].(int)
					if isOk {
						return data[idQuery], nil
					}
					return nil, nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)
