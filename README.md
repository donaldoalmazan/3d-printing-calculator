# 3D Printing Cost Calculator

## Overview
This is a command-line tool written in Go that calculates the cost of producing a 3D-printed object based on user-provided specifications. It uses a JSON file to store material data, allowing easy updates to material costs without modifying the code.

## Features
- Interactive CLI for entering object weight, print time, and material selection.
- Dynamic loading of materials from a JSON file.
- Simple cost calculation based on material, machine, and labor costs.
- Easy extensibility for additional cost factors.

## Installation
Ensure you have Go installed (version 1.17 or later). Then, clone the repository and initialize the project:

```bash
# Clone the repository
git clone https://github.com/yourusername/3d-printing-calculator.git
cd 3d-printing-calculator

# Initialize Go module
go mod tidy
```

## Usage
Run the program from the project root:

```bash
go run cmd/app/main.go
```

The program will prompt you to enter:
1. Estimated weight (in grams)
2. Estimated print time (in hours)
3. Material selection (from the available options)

It will then calculate and display the total cost.

## Materials Configuration
Material data is stored in `materials.json` in the project root. You can modify this file to update material prices or add new materials.

Example format:

```json
[
  { "name": "PLA", "costPerGram": 0.10, "density": 1.24 },
  { "name": "ABS", "costPerGram": 0.12, "density": 1.04 }
]
```

## Contributing
1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m "Added new feature"`)
4. Push to the branch (`git push origin feature-branch`)
5. Open a Pull Request

## License
This project is licensed under the MIT License. See `LICENSE` for details.