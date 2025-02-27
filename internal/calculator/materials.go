package calculator

import (
	"encoding/json"
	"io/ioutil"
)

// Material holds properties for a printing material.
type Material struct {
	Name        string  `json:"name"`
	CostPerGram float64 `json:"costPerGram"`
	Density     float64 `json:"density"`
}

// LoadMaterials reads materials from a JSON file and returns a map for easy lookup.
func LoadMaterials(filepath string) (map[string]Material, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var materials []Material
	if err := json.Unmarshal(data, &materials); err != nil {
		return nil, err
	}

	matMap := make(map[string]Material)
	for _, mat := range materials {
		matMap[mat.Name] = mat
	}
	return matMap, nil
}
