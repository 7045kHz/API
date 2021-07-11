package main

// Creating base API Document
func Document_API() API {
	A := API{
		Name:        "MicroService API",
		Description: "Documentation for MicroService API",
		Version:     1.1,
		EndPoint: []API_EndPoint{
			API_EndPoint{
				Name:        "Status",
				Description: "Provides Status of MicroService",
				Path:        "/status",
				Help:        "http://hostname:PORT/status/help",
				Example:     "http://hostname:PORT/status",
			},
			API_EndPoint{
				Name:        "api",
				Description: "Provides help for api values",
				Path:        "/api",
				Help:        "http://hostname:PORT/api/help",
				Example:     "http://hostname:PORT/api",
			},
			API_EndPoint{
				Name:        "add",
				Description: "Takes input for oracle server.php",
				Path:        "/add",
				Help:        "http://hostname:PORT/add/help",
				Example:     "http://hostname:PORT/add",
				Fields: []API_Fields{
					API_Fields{
						Description: "Takes first value for adding",
						Name:        "INT1",
						Value:       "int",
					},
					API_Fields{
						Description: "Takes second value for adding",
						Name:        "INT2",
						Value:       "int",
					},
				},
			},
		},
	}

	return A
}
