package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/*
 * Complete the function below.
 * https://en.wikipedia.org/w/api.php?action=parse&section=0&prop=text&format=json&page=[topic]
 */

type Article struct {
	Title  string            `json:"title"`
	PageId int               `json:"pageid"`
	Text   map[string]string `json:"text"`
}

type ResponseBody struct {
	Parse Article `json:"parse"`
}

func GetTopicCount(topic string) int {
	responseBody := &ResponseBody{}
	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=parse&section=0&prop=text&format=json&page=%s", topic)
	resp, err := http.Get(url)
	if err != nil {
		return 0
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(responseBody)

	return strings.Count(responseBody.Parse.Text["*"], topic)
}

// func main() {
// 	fmt.Println(getTopicCount("pizza"))
// }
