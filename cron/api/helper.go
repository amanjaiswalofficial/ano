package api

import (
	"cron/job"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

type NewsItemArray []job.NewsItem

// read data from the JSON file whose path is given
// Return its data after unmarshalling in an interface{}
func readDataFromFile(filePath string) map[string]NewsItemArray {
	jsonFile, err := os.Open(filePath)

	var result map[string]NewsItemArray
	//var result map[string]interface{}
	if err != nil {
		fmt.Println(err)
		return result
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &result)

	return result
}

// sortFeedByTime will take all the news Items from all the RSS Sources
// and put them in 1 array, sorting them by time and returning them
// Also, it will slice the array of each source upto provided maxNewsLimit
func sortFeedByTime(jsonData map[string]NewsItemArray, maxNewsLimit int) NewsItemArray {

	if jsonData == nil {
		fmt.Println("Unable to sort news")
		return nil
	}

	newsItemsFlattened := make(map[int]job.NewsItem)
	var newsItemsSorted []job.NewsItem
	var timestampMapper []int

	// for each source's array of newsItems
	for _, listOfItems := range jsonData {

		
		newsLimit := len(listOfItems)
		if newsLimit >= maxNewsLimit {
			newsLimit = maxNewsLimit
		}

		// for each newsItem
		for _, newsItem := range listOfItems[0:newsLimit] {

			parsedTimeKey, _ := time.Parse(time.RFC1123, newsItem.Date)
			newstimeStamp := int(parsedTimeKey.Unix())

			// make a map with timestamp of newsItem as key and newsItem as value
			newsItemsFlattened[newstimeStamp] = newsItem

			timestampMapper = append(timestampMapper, newstimeStamp)
		}
	}

	// sort the timestamps in reverse to find out which timestamp is latest
	sort.Sort(sort.Reverse(sort.IntSlice(timestampMapper)))

	// for each timestamp, append its newsItem into a new slice
	for _, timestamp := range timestampMapper {
		newsItemsSorted = append(newsItemsSorted, newsItemsFlattened[timestamp])
	}

	return newsItemsSorted
}
