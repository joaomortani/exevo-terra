package adapter

import (
	"encoding/json"
	"fmt"
)

// ResourceData é apenas um apelido para map[string]interface{}
// Isso facilita a leitura do código: onde ler "ResourceData", entenda "Mapa Genérico"
type ResourceData map[string]interface{}

// ToMap converte QUALQUER struct (interface{}) para um Mapa
func ToMap(input interface{}) (ResourceData, error) {
	// 1. Truque de Mestre: Converter Struct -> JSON
	// O json.Marshal sabe ler os ponteiros da AWS e transformar em valores reais
	data, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter para JSON: %w", err)
	}

	// 2. JSON -> Mapa (ResourceData)
	// Agora que virou JSON, podemos jogar dentro de um mapa flexível
	var result ResourceData
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("erro ao converter para Mapa: %w", err)
	}

	return result, nil
}

// BatchToMap faz a mesma coisa, mas para uma lista de structs
func BatchToMap(inputs []interface{}) ([]ResourceData, error) {
	result := make([]ResourceData, len(inputs))
	for i, v := range inputs {
		res, err := ToMap(v)
		if err != nil {
			return nil, err
		}
		result[i] = res
	}
	return result, nil
}
