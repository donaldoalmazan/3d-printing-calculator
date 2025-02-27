package calculator

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestCostCalculation tests the Cost() method of ObjectSpec.
func TestCostCalculation(t *testing.T) {
	// Create a sample material for testing.
	mat := Material{
		Name:        "TestMaterial",
		CostPerGram: 0.10,
		Density:     1.0,
	}

	// Define test cases.
	testCases := []struct {
		name               string
		object             ObjectSpec
		machineCostPerHour float64
		laborCostPerHour   float64
		expectedCost       float64
	}{
		{
			name: "simple case",
			object: ObjectSpec{
				Weight:    100, // 100 grams
				PrintTime: 2,   // 2 hours
				Material:  mat,
			},
			machineCostPerHour: 0.50,
			laborCostPerHour:   10.00,
			// Expected cost = (100*0.10) + (2*0.50) + (2*10.00) = 10 + 1 + 20 = 31
			expectedCost: 31.0,
		},
		{
			name: "zero weight",
			object: ObjectSpec{
				Weight:    0,
				PrintTime: 2,
				Material:  mat,
			},
			machineCostPerHour: 0.50,
			laborCostPerHour:   10.00,
			// Expected cost = (0*0.10) + (2*0.50) + (2*10.00) = 0 + 1 + 20 = 21
			expectedCost: 21.0,
		},
		{
			name: "zero print time",
			object: ObjectSpec{
				Weight:    100,
				PrintTime: 0,
				Material:  mat,
			},
			machineCostPerHour: 0.50,
			laborCostPerHour:   10.00,
			// Expected cost = (100*0.10) + (0*0.50) + (0*10.00) = 10
			expectedCost: 10.0,
		},
	}

	// Run each test case.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.object.Cost(tc.machineCostPerHour, tc.laborCostPerHour)
			if actual != tc.expectedCost {
				t.Errorf("expected cost %f but got %f", tc.expectedCost, actual)
			}
		})
	}
}

// TestLoadMaterials tests the LoadMaterials function which reads materials from a JSON file.
func TestLoadMaterials(t *testing.T) {
	// Sample JSON data representing two materials.
	data := `[
		{"name": "PLA", "costPerGram": 0.10, "density": 1.24},
		{"name": "ABS", "costPerGram": 0.12, "density": 1.04}
	]`

	// Create a temporary file to hold the JSON data.
	tmpfile, err := ioutil.TempFile("", "materials_*.json")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	// Clean up the temporary file when done.
	defer os.Remove(tmpfile.Name())

	// Write the sample data into the temporary file.
	if _, err := tmpfile.Write([]byte(data)); err != nil {
		t.Fatalf("failed to write to temporary file: %v", err)
	}
	tmpfile.Close()

	// Load materials using the temporary JSON file.
	materials, err := LoadMaterials(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadMaterials returned error: %v", err)
	}

	// Check that we loaded exactly two materials.
	if len(materials) != 2 {
		t.Errorf("expected 2 materials, got %d", len(materials))
	}

	// Verify that both "PLA" and "ABS" are present.
	if _, ok := materials["PLA"]; !ok {
		t.Error("expected material PLA to be loaded")
	}
	if _, ok := materials["ABS"]; !ok {
		t.Error("expected material ABS to be loaded")
	}
}
