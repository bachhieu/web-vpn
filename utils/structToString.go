package utils

import (
	"encoding/json"
	"fmt"
)

func JsonToString(input struct{}, output *string) {

	// Convert struct to JSON string
	jsonString, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert JSON string to normal string
	*output = string(jsonString)
}
