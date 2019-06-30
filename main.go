package main

import (
	"encoding/json"
	"fmt"

	"gotten/flattened"
)

func main() {
	j_string_map := `
  {
    "a": true,
    "b": false,
    "c": {
      "f": true,
      "g": {
        "m": true,
        "n": false
      }
    }
  }
  `

	var empty map[string]interface{}

	json.Unmarshal([]byte(j_string_map), &empty)

	fmt.Printf("empty: %+v\n", empty)

	fmt.Println(j_string_map)

	f := flattened.FlattenBools("", empty)
	fmt.Println(f)
}
