package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/clbanning/mxj"
	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/mapstructure"
)

func main() {
	var numRequest = SUBRequest{First: 1, Second: 2}

	weather, err := getSUM(numRequest)
	if err != nil {
		log.Println(err)
	} else {
		spew.Dump(weather)
	}
}

type SUBRequest struct {
	First  int
	Second int
}
type SUBResponse struct {
	Results int
}

func getSUM(numRequest SUBRequest) (*SUBResponse, error) {
	url := "http://oracle-base.com/webservices/server.php"
	client := &http.Client{}
	sRequestContent := generateRequestContent(numRequest)
	requestContent := []byte(sRequestContent)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestContent))
	fmt.Printf("REQ from http.NewRequest: %s\n\n\n", req)
	if err != nil {
		return nil, err
	}

	req.Header.Add("SOAPAction", `"http://oracle-base.com/webservices/server.php/ws_add"`)
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Accept", "text/xml")
	resp, err := client.Do(req)
	fmt.Printf("RESP from client.Do: %s\n\n\n", resp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("Error Respose " + resp.Status)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("CONTENTS: ", string(contents))
	m, _ := mxj.NewMapXml(contents, true)
	return convertResults(&m)
}

func generateRequestContent(class SUBRequest) string {
	type QueryData struct {
		First  int
		Second int
	}
	// const getTemplate = `<?xml version="1.0" encoding="utf-8"?>
	//<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	//  <soap:Body>
	//    <GetCityWeatherByZIP xmlns="http://ws.cdyne.com/WeatherWS/">
	//      <ZIP>{{.PostalCode}}</ZIP>
	//    </GetCityWeatherByZIP>
	//  </soap:Body>
	//</soap:Envelope>`
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
	querydata := QueryData{First: 1, Second: 2}
	tmpl, err := template.New("getCityWeatherByZIPTemplate").Parse(getTemplate)
	if err != nil {
		panic(err)
	}
	var doc bytes.Buffer
	err = tmpl.Execute(&doc, querydata)
	fmt.Printf("\n\nDOC: %s\n\n", doc.String())
	if err != nil {
		fmt.Printf("\n\nError in tmpl.Execute\n\n")
		panic(err)
	}
	return doc.String()
}

func convertResults(soapResponse *mxj.Map) (*SUBResponse, error) {
	successStatus, _ := soapResponse.ValueForPath("Envelope.Body")
	fmt.Println("Returned: ", successStatus)
	fmt.Println("SOAP RESONSE", soapResponse)
	/*
		success := successStatus.(bool)
		if !success {
			errorMessage, _ := soapResponse.ValueForPath("Envelope.Body.GetCityWeatherByZIPResponse.GetCityWeatherByZIPResult.ResponseText")
			return nil, errors.New("Error Respose " + errorMessage.(string))
		}
	*/
	weatherResult, err := soapResponse.ValueForPath("Envelope.Body")
	if err != nil {
		return nil, err
	}

	var result SUBResponse
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
		// add a DecodeHook here if you need complex Decoding of results -> DecodeHook: yourfunc,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(weatherResult); err != nil {
		return nil, err
	}
	return &result, nil
}
