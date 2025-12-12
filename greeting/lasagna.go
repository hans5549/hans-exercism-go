package greeting

// TODO: define the 'OvenTime' constant
const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// CalculatePreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func CalculatePreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return CalculatePreparationTime(numberOfLayers) + actualMinutesInOven
}
