package files

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

// ImportJSONFile imports json file to map
func ImportJSONFile(fileName string, data interface{}) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(content, &data)
	if err != nil {
		return err
	}
	return nil
}
