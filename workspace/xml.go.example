package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/net/html/charset"
)

func main() {
	// Print the XML comments from the test file, which should
	// contain most of the printable ISO-8859-1 characters.
	r, err := os.Open("SOAP.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer r.Close()
	fmt.Println("file:", r.Name())

	xmlBytes, _ := ioutil.ReadAll(r)
	fmt.Printf("data type %T\n\n", xmlBytes)
	var soap GetEnvelope

	reader := bytes.NewReader(xmlBytes)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&soap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(soap.Body.GetResponse.Response)
	fmt.Println(soap)

}
