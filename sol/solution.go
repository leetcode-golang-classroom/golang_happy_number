package sol

func isHappy(n int) bool {
	var getNext = func(n int) int {
		sum := 0
		for n > 0 {
			digit := n % 10
			n /= 10
			sum += digit * digit
		}
		return sum
	}
	slow, fast := n, getNext(n)
	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return fast == 1
}
