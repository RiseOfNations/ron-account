package util

import "github.com/json-iterator/go"

func JsonMarshal(data interface{}) ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(data)
}

func JsonUnmarshal(input []byte, data interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(input, data)
}
