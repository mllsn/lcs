func fizzBuzz(n int) []string {
	result := make([]string, n, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			result[i-1] = "FizzBuzz"
		case i%5 == 0:
			result[i-1] = "Buzz"
		case i%3 == 0:
			result[i-1] = "Fizz"
		default:
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result
}