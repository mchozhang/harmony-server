package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"harmony-server/handler/dynamodb"
	"harmony-server/handler/graphql"
	"os"
)

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

// RequestBody request body struct of a graphql query
type RequestBody struct {
	Query         string
	Variables     map[string]interface{}
	OperationName string
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request Request) (Response, error) {
	requestBody := RequestBody{}

	err := json.Unmarshal([]byte(request.Body), &requestBody)

	if err != nil {
		// fail to interpret GraphQl query data from the request
		return Response{Body: err.Error(), StatusCode: 500}, err
	}

	graphResult := graphql.ExecuteQuery(requestBody.Query, requestBody.Variables, requestBody.OperationName)
	responseJson, err := json.Marshal(graphResult)

	if err != nil {
		// fail to interpret GraphQl result
		return Response{Body: err.Error(), StatusCode: 400}, err
	}

	// use additional headers to enable CORS
	return Response{
		Body:       string(responseJson),
		StatusCode: 200, Headers: map[string]string{
			"Access-Control-Allow-Headers": "Content-Type, Access-Control-Allow-Origin",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "OPTIONS,POST,GET",
		}}, nil
}

// init function will be automatically invoked before main
// it only needs to run one time after the program start.
func init() {
	// Initialize a session
	region := os.Getenv("AWS_REGION")
	if sess, err := session.NewSession(&aws.Config{
		Region: &region,
	}); err != nil {
		fmt.Printf("Failed to initialize a session to AWS: %s\n", err.Error())
	} else {
		dynamodb.Init(sess)
		fmt.Println("Initialized database client.")
	}
}

func main() {
	lambda.Start(Handler)
}
