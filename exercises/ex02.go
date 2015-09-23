// CB-2 Create entities using standard ops
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"fmt"
	"net/http"

	// "os"
)

type UpdateContextRequest struct {
	ContextElements []ContextElement `json: "contextElements"`
	UpdateAction    string           `json: "updateAction"`
}

type ContextElement struct {
	Id         string      `json: "id"`
	IsPattern  string      `json: "isPattern"`
	Type       string      `json: "type"`
	Attributes []Attribute `json: "attributes"`
}

type Attribute struct {
	Name  string `json: "id"`
	Type  string `json: "type"`
	Value string `json: "value"`
}

func main() {

	ucr := &UpdateContextRequest{
		[]ContextElement{
			ContextElement{"id1", "false", "aType", []Attribute{Attribute{"name1", "type1", "val1"}}},
			ContextElement{"id2", "false", "aType", []Attribute{Attribute{"name2", "type2", "val2"}}},
		},
		"APPEND",
	}

	ucr_json, _ := json.Marshal(ucr)
	url := "http://localhost:1026/v1/updateContext"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(ucr_json))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
