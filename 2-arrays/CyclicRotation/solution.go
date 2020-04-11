package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
	// write your code in Go 1.4
	n := len(A)
	if n == 0 {
		return A
	}
	for i := 0; i < K; i++ {
		rotate(n, A)
	}
	return A
}

func rotate(n int, A []int) {
	tmp := A[n-1]
	for i := n - 1; i >= 1; i-- {
		A[i] = A[i-1]
	}
	A[0] = tmp
}
