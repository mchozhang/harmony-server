package dynamodb

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"os"
	"strconv"
)

var svc *dynamodb.DynamoDB

// Init
// initialize the dynamodb client
func Init() {
	region := os.Getenv("AWS_REGION")
	// Initialize a session
	if sess, err := session.NewSession(&aws.Config{
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("Failed to initialize a session to AWS: %s", err.Error()))
	} else {
		// Create DynamoDB client
		svc = dynamodb.New(sess)
	}
}

// GetLevel
// get level item from dynamoDB
func GetLevel(levelId int) (interface{}, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("levelTableName")),
		Key: map[string]*dynamodb.AttributeValue{
			"level": {
				N: aws.String(strconv.Itoa(levelId)),
			},
		},
	}
	result, err := svc.GetItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if result.Item == nil {
		return nil, errors.New("Could not find Level" + strconv.Itoa(levelId) + ".")
	}

	level := &Level{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &level)
	if err != nil {
		return nil, err
	}
	return level, nil
}

// CreateLevel
// create a new game level to the database
func CreateLevel(levelInput interface{}) (interface{}, error) {
	av, err := dynamodbattribute.MarshalMap(levelInput)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(os.Getenv("levelTableName")),
	}

	result, err := svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return nil, err
	}
	level := &Level{}
	err = dynamodbattribute.UnmarshalMap(result.Attributes, &level)
	return level, err
}
