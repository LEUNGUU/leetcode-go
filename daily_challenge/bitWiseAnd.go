package daily_challenge

func rangeBitwiseAnd(left int, right int) int {
	and := left
	for i := left; i <= right; i++ {
		and &= i
	}
	return and
}
