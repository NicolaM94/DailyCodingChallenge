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

func FindZero(arrayOfInts []int) bool {
	var output int
	for _,n := range(arrayOfInts) {
		if n == output {
			return true
		}
	}
	return false
}

func FindMinimun (arrayOfInts []int) int {
	var counter int
	for _,value := range(arrayOfInts) {
		if value < counter && value > 0 {
			counter = value
		}
	}
	return counter
}
