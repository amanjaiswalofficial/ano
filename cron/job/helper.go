package job

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readDataFromFile() {

}

func writeDataToFile(data RSSData, fileName string) {

	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(fileName, file, 0644)

}

func readExistingData(filePath string) (RSSData, bool){
	jsonFile, err := os.Open(filePath)
	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	data := make(RSSData)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &data)

	return data, true
}


// SetupEnvironment is used to set env with variable name/value from config.json
func SetupEnvironment(configPath string) (bool) {
	jsonFile, err := os.Open(configPath)
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
		return false
    }
	
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

	for key, value := range result {
		os.Setenv(key, value.(string))
	}

	return true
}