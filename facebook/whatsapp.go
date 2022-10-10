package facebook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func SendMsg() {

	uri := "https://graph.facebook.com/v14.0/105114369050724/messages"

	client := http.Client{}
	body := new(Message)

	body.Messaging_product = "whatsapp"
	body.Recipient_type = "individual"
	body.To = "2347062942291"
	body.Type = "text"
	body.Text = &Text{Body: "hello salemzii from the goland Again", Preview_url: true}

	json_data, err := json.Marshal(body)

	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer EAAK3LCZBTUA4BAFUZAHsIuoEHQaK7ZBzESafL27gk2pS52T9AZBgTcGV19kxZCeuFBIKc4XJ076zBAof99VD3p6if2TU0LWfztpOuLK9ZAPKEkDEgfN6hljRJZAtvEAZC579SJK8aXL3CC6MDssasgHvAos8RIQfRm9myiZC7S38hFaSh0su888yDwReiOR12yUANeq3ow1ZB1uQZDZD"}, //[]string{os.Getenv("META_ACCESS_TOKEN")},
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Status)
	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// decode bytes response to a readable json

	var respJson MessageResponse
	err = json.Unmarshal(respByte, &respJson)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(respJson)
}

/*
func SendMediaMsg() {

	uri := "https://graph.facebook.com/v13.0/app_id/messages"

	client := http.Client{}
	body := new(Message)

	body.Messaging_product = "whatsapp"
	body.Recipient_type = "individual"
	body.To = ""
	body.Image = Image{}

	json_data, err := json.Marshal(body)

	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{""},
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Status)
	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// decode bytes response to a readable json

	var respJson MessageResponse
	err = json.Unmarshal(respByte, &respJson)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(respJson)
}

/*
func SendTemplMsg() {

	uri := "https://graph.facebook.com/v13.0/app_id/messages"

	client := http.Client{}
	body := new(Message)

	body.Messaging_product = "whatsapp"
	body.To = ""
	body.Type = "template"
	body.Template = Template{Name: "hello_world", Language: Language{Code: "en_US"}}

	json_data, err := json.Marshal(body)

	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{os.Getenv("META_ACCESS_TOKEN")},
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Status)
	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// decode bytes response to a readable json

	var respJson MessageResponse
	err = json.Unmarshal(respByte, &respJson)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(respJson)
}
*/
