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
