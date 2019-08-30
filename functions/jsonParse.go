package functions

import (
	"encoding/json"
)

func JsonParse(jsonString string) map[string]interface{} {
	var jsonResult map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonResult)

	return jsonResult
}
