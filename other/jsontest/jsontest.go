package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Food string
}

func encodeTest() {
	m := Message{
		Name: "Paulo",
		Food: "Pizza",
	}

	b, err := json.Marshal(m)

	fmt.Printf("Marshalled JSON Object: %s / Error: %s\n", b, err)
}

func decodeTest() {
	var m Message

	b := []byte(`{"Name":"Bob","Food":"Pickle"}`)

	err := json.Unmarshal(b, &m)

	fmt.Printf("Unmarshalled JSON Object: %s / Error: %s\n", m, err)
}

func main() {
	encodeTest()
	decodeTest()
}
