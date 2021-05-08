package job

import (
	"sync"
)

type Job func()

// SyncRSSFeeds will take a slice of URLs to read RSS data from and store it.
func SyncRSSFeeds(urls []string) {
	data := make(RSSData)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go getDataFromURL(url, &data, &wg)
	}

	wg.Wait()

	existingRSSData, exists := readExistingData("data.json")
	if exists {
		data = filterData(data, existingRSSData)
	}
	writeDataToFile(data, "data.json")

}
