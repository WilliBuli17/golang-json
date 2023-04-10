package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	//yang open ini cuma langkah untuk menulis data kedalam file json -- cuma contoh kasus
	writer, err := os.Create("./resources/CustomerOut.json")
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(writer)

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

	err = encoder.Encode(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
