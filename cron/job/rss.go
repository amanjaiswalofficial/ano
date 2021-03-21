package job

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"sync"
)

type PostData struct {
	Channel struct {
		Posts []struct {
			Title       string   `xml:"title"`
			Link        string   `xml:"link"`
			Category    []string `xml:"category"`
			Creator     string   `xml:"creator"`
			PubDate     string   `xml:"pubDate"`
			Updated     string   `xml:"updated"`
			License     string   `xml:"license"`
			Encoded     string   `xml:"encoded"`
			Description string   `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}

type NewsItem struct {
	Date        string
	Title       string
	Description string
	Link        string
}

type NewsItems []NewsItem

type RSSData map[string]NewsItems

func ParseXMLFromResponse(req []byte) *PostData {
	data := &PostData{} // Create and initialise a data variable as a PostData struct
	err := xml.Unmarshal(req, data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func filterData(fetchedData RSSData, existingData RSSData) RSSData {

	var wg sync.WaitGroup
	filteredRSSData := make(RSSData)

	for key := range fetchedData {
		if !reflect.DeepEqual(fetchedData[key], existingData[key]) {
			var unfilteredItems []NewsItem

			for _, val := range fetchedData[key] {
				unfilteredItems = append(unfilteredItems, val)
			}

			for _, val := range existingData[key] {
				unfilteredItems = append(unfilteredItems, val)
			}

			wg.Add(1)
			go fetchUniqueItems(filteredRSSData, key, unfilteredItems, &wg)
		} else {
			log.Printf("Same datas in %v, no rewrite required\n", key)
			filteredRSSData = fetchedData
		}
	}

	wg.Wait()
	return filteredRSSData
}

func fetchUniqueItems(data RSSData, key string, unfilteredData []NewsItem, wg *sync.WaitGroup) {

	listOfItems := make([]NewsItem, 0)
	title2ItemMapping := make(map[string]NewsItem)

	// for each item in the array
	for _, val := range unfilteredData {

		// check if it's title already exists in the map
		_, present := title2ItemMapping[val.Title]

		// if doesn't, meaning its a unique newsItem, so add it to listOfItems
		if !present {
			title2ItemMapping[val.Title] = val // also add it to map for future
			listOfItems = append(listOfItems, val)
		}
	}

	data[key] = listOfItems

	defer wg.Done()
}

// func (item1 interface{}) union(item2 interface{}) ([]string){
// 	fmt.Println(item1)
// 	fmt.Println(item2)

// 	return item1
// }

func getDataFromURL(url string, data *RSSData, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting data from : %v\n", url)
		fmt.Printf(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error decoding body for %v\n", url)
		fmt.Printf(err.Error())
	}

	var newsItems []NewsItem

	RSSResponse := ParseXMLFromResponse(bodyBytes)
	for _, val := range RSSResponse.Channel.Posts {
		newsItem := NewsItem{
			val.PubDate,
			val.Title,
			val.Description,
			val.Link,
		}
		newsItems = append(newsItems, newsItem)
	}

	data2 := *data
	data2[url] = newsItems

	defer wg.Done()
}
