package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"local/3d-printing-calculator/internal/calculator"

	"github.com/sahilm/fuzzy"
	. "github.com/stevegt/goadapt"
)

func main() {
	// Define the -t flag for print time in "HH:MM" format
	printTimeFlag := flag.String("t", "", "Estimated print time (format HH:MM)")
	// Define the -w flag for weight in grams
	weightFlag := flag.Float64("w", 0, "Estimated weight (grams)")
	// Define the -m flag for material abbreviation
	materialFlag := flag.String("m", "", "Material abbreviation")
	flag.Parse()

	if *printTimeFlag == "" {
		log.Fatalf("Print time flag -t is required. Example usage: -t \"02:30\"")
	}

	if *weightFlag <= 0 {
		log.Fatalf("Weight flag -w is required and must be a positive number. Example usage: -w 50.0")
	}

	if *materialFlag == "" {
		log.Fatalf("Material flag -m is required. Example usage: -m PLA")
	}

	// Parse the print time from "HH:MM" to float64 hours
	printTime, err := parseTime(*printTimeFlag)
	if err != nil {
		log.Fatalf("Invalid print time format: %v. Expected format HH:MM", err)
	}

	weight := *weightFlag

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

	// Fuzzy match the abbreviation against available materials
	abbrev := strings.ToLower(*materialFlag)

	var materialNames []string
	var materialNamesLower []string
	for _, material := range config.Materials {
		materialNames = append(materialNames, material.Name)
		materialNamesLower = append(materialNamesLower, strings.ToLower(material.Name))
	}

	matches := fuzzy.Find(abbrev, materialNamesLower)

	if len(matches) == 0 {
		fmt.Println("No materials found matching abbreviation:", *materialFlag)
		fmt.Println("Available materials:")
		for _, name := range materialNames {
			fmt.Println(" -", name)
		}
		os.Exit(1)
	} else if len(matches) > 1 {
		fmt.Println("Multiple materials match the abbreviation:", *materialFlag)
		fmt.Println("Matching materials:")
		for _, match := range matches {
			fmt.Println(" -", materialNames[match.Index])
		}
		os.Exit(1)
	}

	selectedMaterialName := materialNames[matches[0].Index]
	selectedMaterial, ok := config.Materials[selectedMaterialName]
	if !ok {
		log.Fatalf("Material not found: %s", selectedMaterialName)
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

// parseTime converts a time string in "HH:MM" format to float64 hours.
func parseTime(timeStr string) (float64, error) {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("time must be in HH:MM format")
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid hours component: %v", err)
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid minutes component: %v", err)
	}

	if minutes < 0 || minutes >= 60 {
		return 0, fmt.Errorf("minutes must be between 0 and 59")
	}

	return float64(hours) + float64(minutes)/60.0, nil
}
