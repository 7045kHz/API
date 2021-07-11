package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Set_MicroService() MicroService {
	var W MicroService

	W.Name = "MicroService for Oracle server calc"
	W.Application = "APEX MicroService"
	W.Description = "Connects to Oracle server.php and converts results to JSON"
	W.Version = 1.1
	W.Status = "success"
	return W
}

func Status(w http.ResponseWriter, r *http.Request) {
	x := Set_MicroService()
	result, _ := json.Marshal(x)
	fmt.Fprintf(w, string(result))

}
func API_documentation_Help(w http.ResponseWriter, r *http.Request) {
	/*
		x := Document_API()
		prettyJSON, _ := json.MarshalIndent(x, "", "    ")
	*/

	fmt.Fprintf(w, `
	<!doctype html> 
	<html> 
		<head> 
			<title>API Documentation Help</title> 
			<link rel="stylesheet" href="http://localhost:50004/static/css/api.css"> 
		</head> 
		<body> 
		<h1>Provides assistance for API</h1>
		</body>
		</html>
	`)
}
func API_documentation(w http.ResponseWriter, r *http.Request) {

	x := Document_API()
	prettyJSON, _ := json.MarshalIndent(x, "", "    ")
	fmt.Fprintf(w, `
	<!doctype html> 
	<html> 
		<head> 
			<title>API Documentation</title> 
			<link rel="stylesheet" href="http://localhost:50004/static/css/api.css"> 
		</head> 
		<body> 
	`)

	h := fmt.Sprintf("<h1>%s %.1f</h1>", x.Description, x.Version)
	fmt.Fprintf(w, h)
	h = fmt.Sprintf(`
	<table>
	<tr>
	 <th>Endpoint</th>
	 <th>Description</th>
	 <th>Path</th>
	 <th>Example</th>
	 <th>Help</th>
	</tr>`)
	fmt.Fprintf(w, h)
	for i := 0; i < len(x.EndPoint); i++ {
		fmt.Println(x.EndPoint[i].Name)
		h = fmt.Sprintf(`
		<tr>
			 <td>%s</td>
			 <td>%s</td>
			 <td>%s</td>
			 <td><a href="%s">%s</a></td>
			 <td><a href="%s">Details</a></td>
		</tr>	`, x.EndPoint[i].Name, x.EndPoint[i].Description, x.EndPoint[i].Path, x.EndPoint[i].Example, x.EndPoint[i].Example, x.EndPoint[i].Help)
		fmt.Fprintf(w, h)
	}
	fmt.Fprintf(w, "</table>")
	fmt.Fprintf(w, "<code>")
	fmt.Fprintf(w, string(prettyJSON))
	fmt.Fprintf(w, "</code>")
	fmt.Fprintf(w, `</body></html>`)

}

func main() {
	// initEvents()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Status)
	//	router.HandleFunc("/event", createEvent).Methods("POST")
	router.PathPrefix("/static/css/").Handler(http.FileServer(http.Dir(".")))

	router.HandleFunc("/api", API_documentation).Methods("GET")
	router.HandleFunc("/api", API_documentation_Help).Methods("GET")
	//	router.HandleFunc("/sum/{first},{second}", getOneEvent).Methods("POST")
	//	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	//	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":50004", router))
}
