package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Email struct {
	Key     string         `json:"key"`
	Message MessageContent `json:"message"`
}

type MessageContent struct {
	Text       string `json:"text"`
	Subject    string `json:"subject"`
	To         []To   `json:"to"`
	From_email string `json:"from_email"`
}

type To struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func sendAlert(postUrl string) {
	msgContent := MessageContent{
		Text: "Hey, the Nexus 5 might be available.  You should go " +
			"check!\n\n" +
			"https://play.google.com/store/devices?hl=en",
		Subject:    "Nexus 5 Available",
		From_email: "nexus5available@nexus5.com",
		To: []To{{
			Email: EmailAddress,
			Name:  EmailName,
		}},
	}

	email := Email{
		Key:     ApiKey,
		Message: msgContent,
	}
	buf, _ := json.Marshal(email)
	body := bytes.NewBuffer(buf)
	res, _ := http.Post(postUrl, "text/json", body)
	defer res.Body.Close()
}
