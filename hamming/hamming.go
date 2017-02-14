package hamming

import "errors"

const testVersion = 5

func Distance(str1, str2 string) (int, error) {
	count := 0
	if len(str1) != len(str2) {
		return count, errors.New("Patrameters not equal")
	}

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			count += 1
		}
	}

	return count, nil

}
