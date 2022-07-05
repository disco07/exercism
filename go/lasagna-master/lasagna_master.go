package lasagna

// PreparationTime return le estimated time to cook.
func PreparationTime(slice []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(slice) * time
}

// Quantities determines quantities of needle and sauce to cook lasagna.
func Quantities(slice []string) (noodles int, sauce float64) {
	var qNoodle int
	var qSauce float64
	for _, s := range slice {
		if s == "sauce" {
			qSauce++
		} else if s == "noodles" {
			qNoodle++
		}
	}
	return qNoodle * 50, qSauce * 0.2
}

// AddSecretIngredient returns my list of ingredient and add secret ingredient.
func AddSecretIngredient(friendsList []string, myList []string) []string {
	myList[len(friendsList)] = friendsList[len(friendsList)-1]
	return myList
}

// ScaleRecipe return a slice of float64 of the amounts needed
// for the desired number of portions.
func ScaleRecipe(amounts []float64, portions int) []float64 {
	var quantities []float64
	for _, quantity := range amounts {
		quantities = append(quantities, quantity*float64(portions/2))
	}
	return quantities
}
