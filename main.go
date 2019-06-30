package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	j_string_map := `
  {
    "a": 5,
    "b": 6,
    "c": {
      "f": 9,
      "g": {
        "m": 17,
        "n": 3
      }
    }
  }
  `

	var empty map[string]interface{}

	json.Unmarshal([]byte(j_string_map), &empty)

	fmt.Printf("empty: %+v\n", empty)

	fmt.Println(j_string_map)
}
