package gotten

import (
	"fmt"
)

func Flatten() {

}

// FlattenInts returns a map of string ints flattened
func FlattenInts(v map[string]interface{}) map[string]int {
	flattened := make(map[string]int)

	for k, v2 := range v {
		if v3, ok := v2.(int); ok {
			flattened[k] = v3
			continue
		}

		if v3, ok := v2.(map[string]interface{}); ok {
			FlattenInts(v3)
		}

		fmt.Println(k)
		fmt.Println(v2)
	}

	fmt.Printf("f %+v", flattened)

	return flattened
}

func FlattenFloats() {

}

func FlattenStrings() {

}
