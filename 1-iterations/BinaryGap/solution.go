package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	// write your code in Go 1.4
	x := N
	d := 0
	max := 0
	started := false
    for x > 0 {
		a := x % 2
		if a == 0 {
			if started {
				d++
			}
		} else {
			if d > max {
				max = d
			}
			d = 0
			started = true
		}
		x /= 2
	}
    
    return max
}
