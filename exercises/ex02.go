// CB-2 Create entities using standard ops
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UpdateContextRequest struct {
	ContextElements []ContextElement `json:"contextElements"`
	UpdateAction    string           `json:"updateAction"`
}

type ContextElement struct {
	Id         string      `json:"id"`
	IsPattern  string      `json:"isPattern"`
	Type       string      `json:"type"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
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

type NGSI interface {
	ToNGSI() ContextElement
}

func (r *Room) ToNGSI() ContextElement {
	return ContextElement{
		Id:        r.Name,
		IsPattern: "false",
		Type:      "Room",
		Attributes: []Attribute{
			Attribute{"temperature", "float", strconv.FormatFloat(r.Temperature, 'f', -1, 32)},
			Attribute{"presence", "boolean", strconv.FormatBool(r.Presence)},
			Attribute{"status", "string", r.Status},
		},
	}
}

func (d *Door) ToNGSI() ContextElement {
	return ContextElement{
		Id:        d.Name,
		IsPattern: "false",
		Type:      "Door",
		Attributes: []Attribute{
			Attribute{"locked", "boolean", strconv.FormatBool(d.Locked)},
			Attribute{"closed", "boolean", strconv.FormatBool(d.Closed)},
		},
	}
}

func main() {
	// Create array of context elements
	entities := []NGSI{
		&Room{"Bedroom1", 25.5, false, "OK"},
		&Room{"Bedroom2", 26.0, true, "Needs cleaning"},
		&Room{"Kitchen", 28.9, true, "OK"},
		&Door{"Frontdoor", false, true},
		&Door{"Backdoor", false, false},
	}

	UpdateContext(entities, "APPEND")
}

func UpdateContext(entities []NGSI, action string) error {
	contextElements := make([]ContextElement, len(entities))
	for i, e := range entities {
		contextElements[i] = e.ToNGSI()
	}

	ucr := &UpdateContextRequest{contextElements, action}

	ucr_json, _ := json.Marshal(ucr)
	fmt.Println(string(ucr_json))

	url := "http://localhost:1026/v1/updateContext"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(ucr_json))
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
