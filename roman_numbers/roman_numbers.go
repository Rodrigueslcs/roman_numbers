package roman_numbers

import (
	"strings"
)

func Roman(s string) (response *Response) {

	biggestSum := 0
	biggestSequence := ""

	sum := 0
	sequence := ""
	s = strings.ToUpper(s)
	r := strings.Split(s, "")
	for i := 0; i < len(r); i++ {
		if vn, ok := mapRoman[r[i]]; ok {
			sum += vn
			sequence += r[i]
		} else if sum >
			biggestSum {

			biggestSum = sum
			biggestSequence = sequence
			sum = 0
			sequence = ""
		}
		if len(r) == i+1 {
			if sum >
				biggestSum {

				biggestSum = sum
				biggestSequence = sequence
			}
		}

	}

	return newResponse(biggestSequence, biggestSum)

}
