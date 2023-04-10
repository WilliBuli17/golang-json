package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonString := `{"FirstName":"AAA","MiddleName":"BBB","LastName":"CCC","Age":17,"Married":false,"Hobbies":["Gaming","Reading","Coding"],"Addresses":[{"Street":"Jalan Dulu Aja","Country":"Fatamorgana","PostalCode":"1243455"},{"Street":"Ndak tau, kok tanya saya","Country":"wakanda +26","PostalCode":"265398"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestDecodeArray(t *testing.T) {
	jsonString := `[{"Street":"Jalan Dulu Aja","Country":"Fatamorgana","PostalCode":"1243455"},{"Street":"Ndak tau, kok tanya saya","Country":"wakanda +26","PostalCode":"265398"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
