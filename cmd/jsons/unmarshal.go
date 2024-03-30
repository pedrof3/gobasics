package jsons

import (
	"encoding/json"
	"fmt"
	"log"
)

type Car struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  uint16 `json:"year"`
}

func UnmarshalExample() {
	var c1 Car
	exJson := []byte(`{
		"brand": "Fiat",
		"model": "Uno",
		"year": 2008
	}`)

	err := json.Unmarshal(exJson, &c1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c1)
}
