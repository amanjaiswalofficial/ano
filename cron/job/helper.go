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
