package greeting

func PreparationTime(layers []string, avgPrepTime int) int {
	if len(layers) == 0 {
		return 0
	} else if avgPrepTime < 0 {
		return 0
	}

	if avgPrepTime == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * avgPrepTime
	}
}

func Quantities(layers []string) (int, float64) {
	noodleCount := 0
	sauceCount := 0.0

	if len(layers) == 0 {
		return noodleCount, sauceCount
	}

	for index := 0; index < len(layers); index++ {
		switch layers[index] {
		case "noodles":
			noodleCount += 50
		case "sauce":
			sauceCount += 0.2
		}
	}

	return noodleCount, sauceCount
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	if len(friendsList) == 0 || len(myList) == 0 {
		return myList
	}

	myList[len(myList)-1] = friendsList[len(friendsList)-1]

	return myList
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	if len(quantities) == 0 || portions <= 0 {
		return quantities
	}

	result := make([]float64, len(quantities))
	scaleFactor := float64(portions) / 2.0
	for index := 0; index < len(quantities); index++ {
		result[index] = quantities[index] * scaleFactor
	}

	return result
}
