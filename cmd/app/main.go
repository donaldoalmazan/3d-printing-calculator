package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"local/3d-printing-calculator/internal/calculator"

	. "github.com/stevegt/goadapt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for weight
	fmt.Print("Enter estimated weight (grams): ")
	weightInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading weight: %v", err)
	}
	weightInput = strings.TrimSpace(weightInput)
	weight, err := strconv.ParseFloat(weightInput, 64)
	if err != nil {
		log.Fatalf("Invalid weight input: %v", err)
	}

	// Prompt for print time
	fmt.Print("Enter estimated print time (hours): ")
	printTimeInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading print time: %v", err)
	}
	printTimeInput = strings.TrimSpace(printTimeInput)
	printTime, err := strconv.ParseFloat(printTimeInput, 64)
	if err != nil {
		log.Fatalf("Invalid print time input: %v", err)
	}

	// Load available materials from JSON file.
	// Since config.json is in the project root, use a relative path.
	configPath := filepath.Join("config.json")
	config, err := calculator.LoadConfig(configPath)
	// demo of the goadapt Ck() function
	Ck(err)
	// instead of:
	/*
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
	*/

	// Display available materials.
	fmt.Println("Available materials:")
	for name := range config.Materials {
		fmt.Println(" -", name)
	}

	// Prompt for material selection.
	fmt.Print("Enter material name: ")
	materialInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading material: %v", err)
	}
	materialInput = strings.TrimSpace(materialInput)
	selectedMaterial, exists := config.Materials[materialInput]
	if !exists {
		log.Fatalf("Material '%s' not found.", materialInput)
	}

	// Create object specification using user inputs.
	object := calculator.ObjectSpec{
		Weight:    weight,
		PrintTime: printTime,
		Material:  selectedMaterial,
	}

	// Calculate the total cost.
	totalCost := object.Cost(config)
	fmt.Printf("\nTotal cost to print the object using %s: $%.2f\n", selectedMaterial.Name, totalCost)
}
