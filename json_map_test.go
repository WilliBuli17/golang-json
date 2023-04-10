package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMap(t *testing.T) {
	jsonString := `{"id":"P-0001","name":"Macbook Pro","price":1000000,"image_url":"http://example.img/image.png"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestJSONMapEncode(t *testing.T) {
	result := map[string]interface{}{
		"id":        "P-0001",
		"name":      "Macbook Pro",
		"price":     1_000_000,
		"image_url": "http://example.img/image.png",
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
