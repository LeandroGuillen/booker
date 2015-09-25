// Write a program (or web/mobile application) that
// - Asks for user input (one value)
// - Updates Locked attribute of Frontdoor entity using that input
// - Queries the entity and check the result
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Room struct {
	Name        string
	Temperature float64 `json:"temperature"`
	Presence    bool    `json:"presence"`
	Status      string  `json:"status"`
}

func (r *Room) ToNGSI() ContextElement {
	return ContextElement{
		ContextEntity: ContextEntity{
			Id:        r.Name,
			IsPattern: false,
			Type:      "Room",
		},
		Attributes: []Attribute{
			Attribute{"temperature", "float", strconv.FormatFloat(r.Temperature, 'f', -1, 32)},
			Attribute{"presence", "boolean", strconv.FormatBool(r.Presence)},
			Attribute{"status", "string", r.Status},
		},
	}
}

type Door struct {
	Name   string
	Locked bool `json:"locked"`
	Closed bool `json:"closed"`
}

func (d *Door) ToNGSI() ContextElement {
	return ContextElement{
		ContextEntity: ContextEntity{
			Id:        d.Name,
			IsPattern: false,
			Type:      "Door",
		},
		Attributes: []Attribute{
			Attribute{"locked", "boolean", strconv.FormatBool(d.Locked)},
			Attribute{"closed", "boolean", strconv.FormatBool(d.Closed)},
		},
	}
}

func main() {
	// Read value
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Frontdoor.Locked :")
	value, _, err := reader.ReadLine()

	var locked bool
	if locked, err = strconv.ParseBool(string(value)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Update Frontdoor.Locked state
	UpdateContext([]NGSI{&Door{"Frontdoor", locked, false}}, "UPDATE")

	// Check update went ok
	QueryContext([]ContextEntity{ContextEntity{Id: "Frontdoor"}}, []string{"locked"})
}
