/**
import data from json files to the dynamoDB table
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"io/ioutil"
	"os"
	"strconv"
)

type Item struct {
	Level  int      `json:"level,int,omitempty"`
	Size   int      `json:"size,int,omitempty"`
	Colors []string `json:"colors,omitempty"`
	Cells  [][]struct {
		TargetRow int `json:"targetRow,int"`
		Steps     int `json:"steps,int"`
		Col       int `json:"col,int"`
		Row       int `json:"row,int"`
	} `json:"cells,omitempty"`
}

// get table items from JSON file
func getItems() []Item {
	var items []Item
	for i := 1; i < 31; i++ {
		filename := "../data/level" + strconv.Itoa(i) + ".json"
		item := Item{}
		if importJSONDataFromFile(filename, &item) {
			items = append(items, item)
		}
	}
	return items
}

/**
 * helper function to import data from json files
 */
func importJSONDataFromFile(filename string, result interface{}) bool {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return true
}

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile: "harmony-server",
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	items := getItems()

	tableName := "harmony-server-Levels"

	for _, item := range items {
		av, err := dynamodbattribute.MarshalMap(item)
		if err != nil {
			fmt.Println("Got error marshalling map:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Create item in table Movies, item will be updated if it exists
		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(tableName),
		}

		_, err = svc.PutItem(input)
		if err != nil {
			fmt.Println("Got error calling PutItem:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}
