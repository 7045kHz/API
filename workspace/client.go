package main

import (
	//	"bytes"

	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"golang.org/x/net/html/charset"
)

func main() {
	var numRequest SendEnvelope
	numRequest.Body.ws_add.First = 5
	numRequest.Body.ws_add.Second = 2

	// Using bytes and decoder because xml.Unmarshal dies on encoding="ISO-8859-1"
	var p GetEnvelope
	var xmlBytes []byte
	xmlBytes, err := getSUM(numRequest)

	reader := bytes.NewReader(xmlBytes)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&p)
	//fmt.Printf("reader: %v\n\n", reader)
	if err != nil {
		fmt.Println("DECODE ERROR: ", err)
	}

	fmt.Printf("%d + %d = %s ", numRequest.Body.ws_add.First, numRequest.Body.ws_add.Second, p.Body.GetResponse.Response)
	json, _ := json.Marshal(p)
	fmt.Printf("Json: %s", json)
}

func getSUM(numRequest SendEnvelope) ([]byte, error) {
	url := "http://oracle-base.com/webservices/server.php"
	client := &http.Client{}
	sRequestContent := generateRequestContent_sum(numRequest)
	requestContent := []byte(sRequestContent)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestContent))

	if err != nil {
		fmt.Println("Error in req http.newreq")
		return nil, err
	}

	req.Header.Add("SOAPAction", `"http://oracle-base.com/webservices/server.php/ws_add"`)
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Accept", "text/xml")
	resp, err := client.Do(req)

	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return []byte{}, errors.New("Error Respose " + resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error getting contents from ioutil.ReadAll")
		return []byte{}, err
	}

	return data, nil
}

func generateRequestContent_sum(numRequest SendEnvelope) string {
	type QueryData struct {
		First  int
		Second int
	}

	const getTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
  xmlns:xs="http://www.w3.org/2001/XMLSchema">
  <soap:Body>
    <ws_add xmlns="http://oracle-base.com/webservices/" soap:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
      <int1 xsi:type="xsd:integer">{{.First}}</int1>
      <int2 xsi:type="xsd:integer">{{.Second}}</int2>
    </ws_add>
  </soap:Body>
</soap:Envelope>`
	var querydata QueryData
	querydata.First = numRequest.Body.ws_add.First
	querydata.Second = numRequest.Body.ws_add.Second
	tmpl, err := template.New("getCityWeatherByZIPTemplate").Parse(getTemplate)
	if err != nil {
		panic(err)
	}
	var doc bytes.Buffer

	err = tmpl.Execute(&doc, querydata)

	if err != nil {
		fmt.Printf("\n\nError in tmpl.Execute\n\n")
		panic(err)
	}
	return doc.String()
}
