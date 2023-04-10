package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	Age        int       `json:"age"`
	Married    bool      `json:"married"`
	Hobbies    []string  `json:"hobbies"`
	Addresses  []Address `json:"addresses"`
}

type Address struct {
	Street     string `json:"street"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "AAA",
		MiddleName: "BBB",
		LastName:   "CCC",
		Age:        17,
		Married:    false,
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
		Addresses: []Address{
			{
				Street:     "Jalan Dulu Aja",
				Country:    "Fatamorgana",
				PostalCode: "1243455",
			},
			{
				Street:     "Ndak tau, kok tanya saya",
				Country:    "wakanda +26",
				PostalCode: "265398",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
