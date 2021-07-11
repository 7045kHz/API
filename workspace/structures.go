package main

import "encoding/xml"

type GetEnvelope struct {
	//XMLName xml.Name
	Body Body `xml:"Body" json:"add"`
}

type Body struct {
	//XMLName     xml.Name
	GetResponse ws_addResponse `xml:"ws_addResponse" json:"ws_addResponse"`
}

type ws_addResponse struct {
	//XMLName  xml.Name
	Response string `xml:"return" json:"return"`
}

type SendEnvelope struct {
	XMLName xml.Name
	Body    SendBody
}

type SendBody struct {
	ws_add SendResponse `xml:"ws_add" json:"wa_add"`
}

type SendResponse struct {
	First  int `xml:"int1" json:"int1"`
	Second int `xml:"int2" json:"int2"`
}

type MicroService struct {
	Name        string  `xml:"name" json:"name"`
	Application string  `xml:"application" json:"application"`
	Version     float32 `xml:"version" json:"version"`
	Description string  `xml:"description" json:"description"`
	Status      string  `xml:"status" json:"status"`
}

// API.EndPoint.Name.Field[1..X]
type API struct {
	Name        string         `xml:"name" json:"name"`
	Description string         `xml:"description" json:"description"`
	Version     float32        `xml:"version" json:"version"`
	EndPoint    []API_EndPoint `xml:"endpoint" json:"endpoint"`
}
type API_EndPoint struct {
	Name        string       `xml:"name" json:"name"`
	Description string       `xml:"description" json:"description"`
	Path        string       `xml:"path" json:"path"`
	Example     string       `xml:"example" json:"example"`
	Help        string       `xml:"help" json:"help"`
	Fields      []API_Fields `xml:"fields" json:"fields"`
}
type API_Fields struct {
	Description string `xml:"description" json:"description"`
	Name        string `xml:"name" json:"name"`
	Value       string `xml:"value" json:"value"`
}
