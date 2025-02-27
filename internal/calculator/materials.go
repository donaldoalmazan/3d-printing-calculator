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

// Config holds configuration for the calculator.
type Config struct {
	LaborRate   float64             `json:"laborRate"`
	MachineRate float64             `json:"machineRate"`
	Materials   map[string]Material `json:"materials"`
}

// LoadConfig reads materials from a JSON file and returns a map for easy lookup.
func LoadConfig(filepath string) (config Config, err error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return
	}

	return
}
