package helper

import (
	"encoding/json"
	"log"
)

func ResponseSuccess(meta interface{}, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"meta": meta,
		"data": data,
	}
}

func ResponseError(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"error": data,
	}
}

func EncodeJSON(data map[string]interface{}) []byte {
	dataJSON, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	return dataJSON
}
