# gotten
A simple go module to flatten json

## Overview
This tool takes json and flattens it.

### Features
The following functions are a quick description on how to use:
- `Flatten()` => returns `map[string]interface{}`; it's a generic function that will flatten any map with different types.
- `FlattenInts()` => returns `map[string]int`; it will flatten and return `int`s.
- `FlattenFloats()` => returns `map[string]float64`; it will flatten and return `float64`s.
- `FlattenStrings()` => returns `map[string]string`; it will flatten and return `string`s.
- `FlattenBools()` => returns `map[string]bool`; it will flatten and return `bool`s.

### Usage
See `main.go` for a complete example
