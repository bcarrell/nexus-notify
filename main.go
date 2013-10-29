package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

func sendAlert() {

}

func getContent(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return strings.ToLower(string(contents)), nil
}

func main() {
	result, err := getContent("https://play.google.com/store/devices?hl=en")
	if err != nil {
		fmt.Println("There's a problem.  Exiting.")
		os.Exit(1)
	}

	fmt.Println(strings.Contains(result, "nexus 5"))
}
