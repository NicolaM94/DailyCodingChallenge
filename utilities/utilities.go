package utilities

// Checks for a value in an array
func Contains (a []int, b int) bool {
	for _,value := range a {
		if value == b {
			return true
		}
	}
	return false
}

func Products (a []int) int {

	var result = 1
	for _,value := range(a) {
		result *= value
	}
	return result
}