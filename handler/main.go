package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"harmony-server/handler/dynamodb"
	"harmony-server/handler/graphql"
)

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

// request body struct of a graphql query
type RequestBody struct {
	Query         string
	Variables     map[string]interface{}
	OperationName string
}

/**
 * Handler is our lambda handler invoked by the `lambda.Start` function call
 */
func Handler(ctx context.Context, request Request) (Response, error) {
	requestBody := RequestBody{}

	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return Response{Body: err.Error(), StatusCode: 500}, err
	}
	graphResult := graphql.ExecuteQuery(requestBody.Query, requestBody.Variables, requestBody.OperationName)
	responseJson, err := json.Marshal(graphResult)

	if err != nil {
		return Response{Body: err.Error(), StatusCode: 400}, err
	}

	return Response{
		Body:       string(responseJson),
		StatusCode: 200, Headers: map[string]string{
			"Access-Control-Allow-Headers": "Content-Type, Access-Control-Allow-Origin",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "OPTIONS,POST,GET",
		}}, nil
}

// init function will be automatically invoked before main
func init() {
	fmt.Println("Initialize handler.")
	dynamodb.Init()
}

func main() {
	lambda.Start(Handler)
}
