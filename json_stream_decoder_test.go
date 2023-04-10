package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	//yang open ini cuma langkah untuk mengambil data dari json -- cuma contoh kasus
	open, err := os.Open("./resources/Customer.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(open)

	customer := &Customer{}
	err = decoder.Decode(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
