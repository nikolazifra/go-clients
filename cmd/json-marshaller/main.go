package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Stats struct {
	HP  int32 `json:"hp"`
	ATK int32 `json:"atk"`
	DEF int32 `json:"def"`
}

type Character struct {
	ID    int32  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Stats Stats  `json:"stats"`
}

func main() {
	character := Character{
		ID:   0,
		Name: "",
		Stats: Stats{
			HP:  100,
			ATK: 50,
			DEF: 30,
		},
	}
	byteArray, err := json.Marshal(character)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteArray))
}
