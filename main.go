// hehnope json flattener
// Copyright (C) 2019 hehnope
//
// slurp is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// slurp is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Foobar. If not, see <http://www.gnu.org/licenses/>.
//

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
