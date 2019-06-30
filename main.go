package main

import "fmt"

func main() {
	j_string_map := `
  {
    'a': 5,
    'b': 6,
    'c': {
      'f': 9,
      'g': {
        'm': 17,
        'n': 3
      }
    }
  }
  `

	fmt.Println(j_string_map)
}
