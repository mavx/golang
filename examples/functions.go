package main
import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 4)
	fmt.Println("1 + 4 =", res)

	res = plusPlus(1, 3, 5)
	fmt.Println("1 + 3 + 5 =", res)
}
