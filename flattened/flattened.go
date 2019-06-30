package flattened

import (
	"fmt"
)

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

// FlattenInts returns a map of string ints flattened
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

func FlattenStrings() {

}
