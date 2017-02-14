package raindrops

import "strconv"

const testVersion = 2

var values = []struct {
	input    int
	expected string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {
	var mas []int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			mas = append(mas, i)
		}
	}

	result := ""
	for _, el := range values {
		if elementInSlice(el.input, mas) {
			result += el.expected
		}
	}
	if result == "" {
		return strconv.Itoa(number)
	}

	return result

}

func elementInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
