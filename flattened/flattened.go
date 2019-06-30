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

package flattened

import (
	"fmt"
)

// Flatten will flatten any string:<type>
func Flatten(prefix string, v map[string]interface{}) map[string]interface{} {
	flattened := make(map[string]interface{})

	for k, v2 := range v {
		true_prefix := fmt.Sprintf("%s.%s", prefix, k)

		// If the value in the map is not an interface{} append it
		if v3, ok := v2.(float64); ok {

			// Because we want to append in a.b.c fashion, we avoid prefix '.'
			if prefix != "" {
				flattened[true_prefix] = v3
			} else {
				flattened[k] = v3
			}
			continue
		}

		// If the value in the map is not an interface{} append it
		if v3, ok := v2.(bool); ok {

			// Because we want to append in a.b.c fashion, we avoid prefix '.'
			if prefix != "" {
				flattened[true_prefix] = v3
			} else {
				flattened[k] = v3
			}
			continue
		}

		// If the value in the map is not an interface{} append it
		if v3, ok := v2.(string); ok {

			// Because we want to append in a.b.c fashion, we avoid prefix '.'
			if prefix != "" {
				flattened[true_prefix] = v3
			} else {
				flattened[k] = v3
			}
			continue
		}

		// The map is an interface{}, so use recursion
		if v3, ok := v2.(map[string]interface{}); ok {
			var tmp_map map[string]interface{}

			// Because we want to append in a.b.c fashion, we avoid prefix '.'
			if prefix != "" {
				tmp_map = Flatten(true_prefix, v3)
			} else {
				tmp_map = Flatten(k, v3)
			}

			// Merge the map
			for j, v := range tmp_map {
				flattened[j] = v
			}
		}
	}

	return flattened
}

// FlattenInts will flatten any string:<int>
func FlattenInts(prefix string, v map[string]interface{}) map[string]int {
	flattened := make(map[string]int)

	for k, v2 := range v {
		true_prefix := fmt.Sprintf("%s.%s", prefix, k)

		// If the value in the map is not an interface{} append it
		if v3, ok := v2.(float64); ok {

			// Because we want to append in a.b.c fashion, we avoid prefix '.'
			if prefix != "" {
				flattened[true_prefix] = int(v3)
			} else {
				flattened[k] = int(v3)
			}
			continue
		}

		// The map is an interface{}, so use recursion
		if v3, ok := v2.(map[string]interface{}); ok {
			var tmp_map map[string]int

			// Because we want to append in a.b.c fashion, we avoid prefix '.'
			if prefix != "" {
				tmp_map = FlattenInts(true_prefix, v3)
			} else {
				tmp_map = FlattenInts(k, v3)
			}

			// Merge the map
			for j, v := range tmp_map {
				flattened[j] = v
			}
		}
	}

	return flattened
}

// FlattenFloats will flatten any string:<float64>
func FlattenFloats(prefix string, v map[string]interface{}) map[string]float64 {
	flattened := make(map[string]float64)

	for k, v2 := range v {
		true_prefix := fmt.Sprintf("%s.%s", prefix, k)

		// If the value in the map is not an interface{} append it
		if v3, ok := v2.(float64); ok {

			// Because we want to append in a.b.c fashion, we avoid prefix '.'
			if prefix != "" {
				flattened[true_prefix] = v3
			} else {
				flattened[k] = v3
			}
			continue
		}

		// The map is an interface{}, so use recursion
		if v3, ok := v2.(map[string]interface{}); ok {
			var tmp_map map[string]float64

			// Because we want to append in a.b.c fashion, we avoid prefix '.'
			if prefix != "" {
				tmp_map = FlattenFloats(true_prefix, v3)
			} else {
				tmp_map = FlattenFloats(k, v3)
			}

			// Merge the map
			for j, v := range tmp_map {
				flattened[j] = v
			}
		}
	}

	return flattened
}

// FlattenStrings will flatten any string:<string>
func FlattenStrings(prefix string, v map[string]interface{}) map[string]string {
	flattened := make(map[string]string)

	for k, v2 := range v {
		true_prefix := fmt.Sprintf("%s.%s", prefix, k)

		// If the value in the map is not an interface{} append it
		if v3, ok := v2.(string); ok {

			// Because we want to append in a.b.c fashion, we avoid prefix '.'
			if prefix != "" {
				flattened[true_prefix] = v3
			} else {
				flattened[k] = v3
			}
			continue
		}

		// The map is an interface{}, so use recursion
		if v3, ok := v2.(map[string]interface{}); ok {
			var tmp_map map[string]string

			// Because we want to append in a.b.c fashion, we avoid prefix '.'
			if prefix != "" {
				tmp_map = FlattenStrings(true_prefix, v3)
			} else {
				tmp_map = FlattenStrings(k, v3)
			}

			// Merge the map
			for j, v := range tmp_map {
				flattened[j] = v
			}
		}
	}

	return flattened
}

// FlattenBools will flatten any string:<bool>
func FlattenBools(prefix string, v map[string]interface{}) map[string]bool {
	flattened := make(map[string]bool)

	for k, v2 := range v {
		true_prefix := fmt.Sprintf("%s.%s", prefix, k)

		// If the value in the map is not an interface{} append it
		if v3, ok := v2.(bool); ok {

			// Because we want to append in a.b.c fashion, we avoid prefix '.'
			if prefix != "" {
				flattened[true_prefix] = v3
			} else {
				flattened[k] = v3
			}
			continue
		}

		// The map is an interface{}, so use recursion
		if v3, ok := v2.(map[string]interface{}); ok {
			var tmp_map map[string]bool

			// Because we want to append in a.b.c fashion, we avoid prefix '.'
			if prefix != "" {
				tmp_map = FlattenBools(true_prefix, v3)
			} else {
				tmp_map = FlattenBools(k, v3)
			}

			// Merge the map
			for j, v := range tmp_map {
				flattened[j] = v
			}
		}
	}

	return flattened
}
