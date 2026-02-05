package provider

import (
	"encoding/json"
)

// EnhancedResource é um container genérico que combina o dado bruto da AWS
// com campos extras calculados (ex: listas de IDs achatadas).
type EnhancedResource[T any] struct {
	Raw    T                      // O struct original da AWS SDK
	Extras map[string]interface{} // Seus campos sintéticos (ex: "SimpleSGs")
}

// MarshalJSON é o pulo do gato 🐈.
// Ele funde o Raw e o Extras em um único JSON plano.
// Assim, o Adapter e o Template não sabem que existe separação.
func (e EnhancedResource[T]) MarshalJSON() ([]byte, error) {
	rawBytes, err := json.Marshal(e.Raw)
	if err != nil {
		return nil, err
	}
	var mergedMap map[string]interface{}
	if err := json.Unmarshal(rawBytes, &mergedMap); err != nil {
		return nil, err
	}

	for k, v := range e.Extras {
		mergedMap[k] = v
	}
	return json.Marshal(mergedMap)
}

// EnrichSlice é o helper funcional para aplicar a transformação em lista.
// T = Tipo original da AWS (ex: types.Instance)
func EnrichSlice[T any](items []T, enricher func(item T) map[string]interface{}) []EnhancedResource[T] {
	result := make([]EnhancedResource[T], len(items))
	for i, item := range items {
		result[i] = EnhancedResource[T]{
			Raw:    item,
			Extras: enricher(item),
		}
	}
	return result
}
