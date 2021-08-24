package helper

import (
	"encoding/json"
	"log"
)

func ResponseSuccess(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"message": message,
		"data":    data,
	}
}

func ResponseError(message string) map[string]interface{} {
	return map[string]interface{}{
		"success": false,
		"message": message,
		"data":    map[string]interface{}{},
	}
}

func EncodeJSON(data map[string]interface{}) []byte {
	dataJSON, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	return dataJSON
}
