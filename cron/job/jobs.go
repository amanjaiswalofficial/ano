package job

import (
	"sync"
)

// SyncRSSFeeds will take a slice of URLs to read RSS data from and store it.
func SyncRSSFeeds(urls []string) {
	data := make(RSSData)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go getDataFromURL(url, &data, &wg)
	}

	wg.Wait()

	existingRSSData, exists := readExistingData("test.json")
	if exists {
		data = filterData(data, existingRSSData)
	}
	writeDataToFile(data, "test.json")

}
