package greeting

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	perHoursCount := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(perHoursCount / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	if carsCount < 0 || carsCount == 0 {
		return 0
	}

	tenCounts := carsCount / 10
	oneCounts := carsCount % 10
	return uint(tenCounts*95000 + oneCounts*10000)
}
