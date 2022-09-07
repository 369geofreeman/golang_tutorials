package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int {
	if prepTime > 0 {
		return len(layers) * prepTime
	}
	return len(layers) * 2
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int = 0
	var sauce float64 = 0.0
	for i := range layers {
		if layers[i] == "sauce" {
			sauce += 0.2
		} else if layers[i] == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scaledQuantities int) []float64 {
	sq := float64(scaledQuantities)
	ret := []float64{}

	for i := range quantities {
		ret = append(ret, (quantities[i]/2)*sq)
	}
	return ret
}
