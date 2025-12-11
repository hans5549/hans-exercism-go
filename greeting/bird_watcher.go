package greeting

import "slices"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	if len(birdsPerDay) == 0 {
		return 0
	}

	sum := 0
	for i := 0; i < len(birdsPerDay); i++ {
		sum += birdsPerDay[i]
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	if len(birdsPerDay) == 0 {
		return 0
	}

	weekDays := []int{0 + (week-1)*7, 1 + (week-1)*7, 2 + (week-1)*7, 3 + (week-1)*7, 4 + (week-1)*7, 5 + (week-1)*7, 6 + (week-1)*7}

	sum := 0
	for i := 0; i < len(birdsPerDay); i++ {
		index := slices.Index(weekDays, i)
		if index != -1 {
			sum += birdsPerDay[i]
		}
	}

	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	if len(birdsPerDay) == 0 {
		return birdsPerDay
	}

	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}

	return birdsPerDay
}
