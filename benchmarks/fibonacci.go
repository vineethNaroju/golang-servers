package benchmarks

func RecFib(n uint) uint {
	if n <= 1 {
		return n
	}
	return RecFib(n-1) + RecFib(n-2)
}

func SeqFib(n uint) uint {
	if n <= 1 {
		return uint(n)
	}

	var a,b,c  uint = 0, 1, 1

	for i := uint(2); i <= n; i++ {
		c = a + b
		a = b
		b = c
	}

	return c
}