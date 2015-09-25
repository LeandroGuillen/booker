// Write a program (or web/mobile application) that does the following queries and prints the result
// - Obtain all attributes of Bedroom1 entity
// - Obtain only the Temperature attribute of Kitchen entity
// - Obtain all attributes of Kitchen and Bedroom2 entities in one query
// - Obtain all attributes of entities that match the pattern Bedroom.*
// - Find out whether the doors are closed using the pattern .*door and the Closed attribute
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type QueryContextRequest struct {
	Entities   []ContextEntity `json:"entities"`
	Attributes []string        `json:"attributes,omitempty"`
}

type ContextEntity struct {
	Id        string `json:"id"`
	IsPattern bool   `json:"isPattern"`
	Type      string `json:"type"`
}

type Room struct {
	Name        string
	Temperature float64 `json:"temperature"`
	Presence    bool    `json:"presence"`
	Status      string  `json:"status"`
}

type Door struct {
	Name   string
	Locked bool `json:"locked"`
	Closed bool `json:"closed"`
}

func main() {
	QueryContext([]NGSI{&Room{Name: "Bedroom1"}}, nil)
	QueryContext([]NGSI{&Room{Name: "Kitchen"}}, []string{"temperature"})
	QueryContext([]NGSI{&Room{Name: "Bedroom1"}, &Room{Name: "Kitchen"}}, nil)
	QueryContext([]ContextEntity{ContextEntity{Id: "Bedroom.*", IsPattern: true}}, nil)
	QueryContext([]ContextEntity{ContextEntity{Id: ".*door", IsPattern: true}}, []string{"closed"})
}

func QueryContext(entities []ContextEntity, attributes []string) error {

	var qcr QueryContextRequest
	if attributes != nil {
		qcr = QueryContextRequest{entities, attributes}
	} else {
		qcr = QueryContextRequest{entities, nil}
	}

	qcr_json, _ := json.Marshal(qcr)
	fmt.Println(string(qcr_json))

	url := "http://localhost:1026/v1/queryContext"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(qcr_json))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}
