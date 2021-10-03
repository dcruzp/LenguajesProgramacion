package main

func add(a int, b int) int {
	return a + b
}

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func palindromo(w string) bool {
	return palindromoAux(w, 0, len(w)-1)
}

func palindromoAux(w string, l int, r int) bool {
	if l == r {
		return true
	}

	if w[l] != w[r] {
		return false
	}

	if l < r+1 {
		return palindromoAux(w, l+1, r-1)
	}

	return true
}
