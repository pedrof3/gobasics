package jsons

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Player struct {
	Name     string `json:"name"`
	Age      uint8  `json:"age"`
	Position string `json:"position"`
	Club     string `json:"club"`
	Country  string `json:"country"`
}

func MarshalExample() {
	p1 := Player{
		"Vinicius Jr",
		23,
		"Left-winger",
		"Real Madrid",
		"Brazil",
	}
	fmt.Println(p1)

	p1Json, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p1Json)
	fmt.Println(string(p1Json))
	fmt.Println(bytes.NewBuffer(p1Json))
}
