package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type source struct {
	ID string `json:"id"`
}

type sourcesAPI struct {
	Sources []source `json:"sources"`
}

type topicsAPI struct {
	Articles []Topic `json:"articles"`
}

func getSources(category string) []source {
	body := getData(sourceURL(category))

	var sourceAPI sourcesAPI
	err := json.Unmarshal(body, &sourceAPI)
	if err != nil {
		return nil
	}

	return sourceAPI.Sources
}

func getTopics(sources []source) []Topic {
	var topics []Topic

	for _, source := range sources {
		body := getData(topicURL(source.ID))

		var topicsAPI topicsAPI
		err := json.Unmarshal(body, &topicsAPI)
		if err != nil {
			return nil
		}

		topics = append(topics, topicsAPI.Articles...)

	}

	return topics
}

func sourceURL(category string) string {
	return fmt.Sprintf("https://newsapi.org/v2/top-headlines/sources?category=%s&apiKey=f0530199143e49c5a24d164b7fa735cd", category)
}
func topicURL(id string) string {
	return fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&apiKey=f0530199143e49c5a24d164b7fa735cd", id)
}

func getData(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return body
}
