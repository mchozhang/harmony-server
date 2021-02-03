/**
import data from json files to the dynamoDB table
*/
package script

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
