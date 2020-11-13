package utils

func Sum(slice *[]float64) float64 {
	total := 0.0
	for _, value := range *slice {
		total += value
	}
	return total
}
