package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/graphql-go/graphql"
)

// mocked database
var data map[int]level

func main() {
	// read json data from files, mock database
	data = make(map[int]level)
	for i := 1; i < 31; i++ {
		filename := "data/level" + strconv.Itoa(i) + ".json"
		level := level{}
		if importJSONDataFromFile(filename, &level) {
			level.ID = strconv.Itoa(i)
			data[i] = level
		}
	}

	// handle graphql query
	http.HandleFunc("/graphql", func(writer http.ResponseWriter, request *http.Request) {
		result := executeQuery(request.URL.Query().Get("query"), schema)
		json.NewEncoder(writer).Encode(result)
	})

	// start server
	http.ListenAndServe(":8080", nil)
}

/**
 * handle graphql queries
 */
func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
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
