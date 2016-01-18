package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
