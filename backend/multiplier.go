package backend

func multiplier(num1 int, num2 int) *Result {
	answer := num1 * num2
	message := messageGenerator(answer)

	res := &Result{
		Answer:  answer,
		Message: message,
	}

	return res
}
