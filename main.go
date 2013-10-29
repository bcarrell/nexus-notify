package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	fmt.Println("Running nexus-notify...")
	for {
		getResult := getContent("https://play.google.com/store/devices?hl=en")
		possibilities := [2]string{"nexus 5", "nexus5"}

		for _, possibility := range possibilities {
			if strings.Contains(getResult, possibility) {
				sendAlert("https://mandrillapp.com/api/1.0/messages/send.json")
				os.Exit(0)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
