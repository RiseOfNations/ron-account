package util

import "github.com/json-iterator/go"

func Marshal(data interface{}) ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(data)
}

func Unmarshal(input []byte, data interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(input, data)
}
