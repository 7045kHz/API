package main

import (
	"encoding/xml"
	"fmt"
)

type MyRespEnvelope struct {
	XMLName xml.Name
	Body    Body
}

type Body struct {
	XMLName     xml.Name
	GetResponse completeResponse `xml:"activationPack_completeResponse"`
}

type completeResponse struct {
	XMLName xml.Name `xml:"activationPack_completeResponse"`
	Id      string   `xml:"Id,attr"`
	MyVar   string   `xml:"activationPack_completeResult"`
}

func main() {

	Soap := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/">
<soap:Body>
<activationPack_completeResponse Id="http://tempuri.org/">
<activationPack_completeResult xsi:type="xsd:string">Active</activationPack_completeResult>
</activationPack_completeResponse>
</soap:Body>
</soap:Envelope>`)

	var res MyRespEnvelope
	xml.Unmarshal(Soap, &res)

	fmt.Println(res.Body.GetResponse.Id)
}
