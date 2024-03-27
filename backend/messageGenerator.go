package backend

func messageGenerator(num int) string {
	if num == 0 {
		return "Zero"
	} else if num < 0 {
		return "Negative Number"
	} else {
		return "Positive Number"
	}
}
