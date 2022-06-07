func maximumWealth(accounts [][]int) int {
	max := 0
	for _, v := range accounts {
		sum := sum(v)
		if sum > max {
			max = sum
		}
	}
	return max
}

func sum(n []int) int {
	sum := 0
	for _, v := range n {
		sum = sum + v
	}
	return sum
}