package utils

// Sum returns the sum of all the elements of the float slice it is given
func Sum(slice *[]float64) float64 {
	total := 0.0
	for _, value := range *slice {
		total += value
	}
	return total
}

// CheckIsErrorRaised is given an error and checks if it needs to panic
func CheckIsErrorRaised(err error) {
	if err != nil {
		panic(err)
	}
}
