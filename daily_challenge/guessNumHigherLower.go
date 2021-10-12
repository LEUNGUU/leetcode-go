package daily_challenge

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right := 1, n
	for {
		if left >= right {
			break
		}
		candidate := (right-left)/2 + left
		res := guess(candidate)
		switch res {
		case -1:
			right = candidate - 1
		case 1:
			left = candidate + 1
		default:
			return candidate
		}
	}
	return left
}
