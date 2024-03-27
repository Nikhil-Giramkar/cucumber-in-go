package backend

func Calculate(num1 int, num2 int, operation string) *Result {
	switch operation {
	case "Add":
		return adder(num1, num2)
	}

	return &Result{
		Answer:  0,
		Message: "Invalid Operation",
	}
}
