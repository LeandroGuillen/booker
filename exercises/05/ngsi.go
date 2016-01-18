package main

type UpdateContextRequest struct {
	ContextElements []ContextElement `json:"contextElements"`
	UpdateAction    string           `json:"updateAction"`
}

type QueryContextRequest struct {
	Entities   []ContextEntity `json:"entities"`
	Attributes []string        `json:"attributes,omitempty"`
}

type ContextEntity struct {
	Id        string `json:"id"`
	IsPattern bool   `json:"isPattern"`
	Type      string `json:"type"`
}

type ContextElement struct {
	ContextEntity
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type NGSI interface {
	ToNGSI() ContextElement
}
