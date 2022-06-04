func fizzBuzz(n int) []string {
	result := make([]string, n, n)

	for i := 1; i <= n; i++ {
		isFizz := i%3 == 0
		isBuzz := i%5 == 0

		str := ""
		if isFizz {
			str += "Fizz"
		}
		if isBuzz {
			str += "Buzz"
		}
		if str == "" {
			str += strconv.Itoa(i)
		}
		result[i-1] = str
	}

	return result
}