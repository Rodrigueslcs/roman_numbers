package roman_numbers

var mapRoman map[string]int

func init() {
	mapRoman = make(map[string]int)

	mapRoman["I"] = 1
	mapRoman["V"] = 5
	mapRoman["X"] = 10
	mapRoman["L"] = 50
	mapRoman["C"] = 100
	mapRoman["D"] = 500
	mapRoman["M"] = 500
}

type Request struct {
	Text string `json:"text"`
}

type Response struct {
	Number string `json:"number"`
	Value  int    `json:"value"`
}

func newResponse(number string, value int) *Response {
	return &Response{
		Number: number,
		Value:  value,
	}
}
