package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func getContent(getUrl string) string {
	res, _ := http.Get(getUrl)
	defer res.Body.Close()
	contents, _ := ioutil.ReadAll(res.Body)
	return strings.ToLower(string(contents))
}

func main() {
	for {
		getResult := getContent("https://play.google.com/store/devices?hl=en")
		possibilities := [3]string{"nexus 5", "nexus5"}

		for _, possibility := range possibilities {
			if strings.Contains(getResult, possibility) {
				sendAlert("https://mandrillapp.com/api/1.0/messages/send.json")
			}
		}
		time.Sleep(10 * time.Second)
	}
}
