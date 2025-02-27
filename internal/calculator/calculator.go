package calculator

// ObjectSpec holds the specifications of the 3D printed object.
type ObjectSpec struct {
	Weight    float64  // Weight in grams
	PrintTime float64  // Print time in hours
	Material  Material // Material used for printing
}

// Cost calculates the total production cost for the object.
// It takes machine and labor costs per hour as parameters.
func (spec ObjectSpec) Cost(machineCostPerHour, laborCostPerHour float64) float64 {
	materialCost := spec.Weight * spec.Material.CostPerGram
	machineCost := spec.PrintTime * machineCostPerHour
	laborCost := spec.PrintTime * laborCostPerHour
	return materialCost + machineCost + laborCost
}
