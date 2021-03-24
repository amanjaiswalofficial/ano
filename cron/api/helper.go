package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// read data from the JSON file whose path is given
// Return its data after unmarshalling in an interface{}
func readDataFromFile(filePath string)(map[string]interface{}){
	jsonFile, err := os.Open(filePath)

	var result map[string]interface{}
	if err != nil {
		fmt.Println(err)
		return result
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &result)

	return result
}
