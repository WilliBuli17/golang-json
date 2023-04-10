package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageUrl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P-0001",
		Name:     "Macbook Pro",
		Price:    1_000_000,
		ImageUrl: "http://example.img/image.png",
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P-0001","name":"Macbook Pro","price":1000000,"image_url":"http://example.img/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	err := json.Unmarshal(jsonBytes, product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.Price)
	fmt.Println(product.ImageUrl)
}
