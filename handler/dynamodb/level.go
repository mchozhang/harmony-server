/**
 * domain layer, graphql of go types
 */
package dynamodb

// Level graphql
type Level struct {
	ID     string   `json:"id,omitempty"`
	Size   int      `json:"size,omitempty"`
	Colors []string `json:"colors,omitempty"`
	Cells  [][]struct {
		TargetRow int `json:"targetRow,int"`
		Steps     int `json:"steps,int"`
		Col       int `json:"col,int"`
		Row       int `json:"row,int"`
	} `json:"cells,omitempty"`
}