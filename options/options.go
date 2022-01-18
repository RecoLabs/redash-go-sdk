package options

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func MapFromString(optionsJSON string) (map[string]interface{}, error) {
	compactOptions := new(bytes.Buffer)
	if err := json.Compact(compactOptions, []byte(optionsJSON)); err != nil {
		return map[string]interface{}{}, fmt.Errorf("options is not a valid json string, options: %v", optionsJSON)
	}

	var options map[string]interface{}
	err := json.Unmarshal(compactOptions.Bytes(), &options)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to unmarshal options: %v, error: %v", compactOptions, err)
	}
	return options, nil
}
