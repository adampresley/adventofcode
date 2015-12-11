package main

type WireCollection map[string]*Wire

var cache = map[string]int{}

func NewWireCollection() WireCollection {
	return make(WireCollection)
}

func (collection WireCollection) ClearCache() {
	cache = make(map[string]int)
}

func (collection WireCollection) EvaluateWireValue(wireName string) int {
	if value, ok := cache[wireName]; ok {
		return value
	}

	result := collection[wireName].Evaluate(collection)
	cache[wireName] = result
	return result
}
