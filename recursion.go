package main

import "fmt"

// ins't technically an algorithm
// a function that calls itself
// it's important to be careful with stack overflows while using recursion, since it can easily lead to an infinite loop
// there enters the concept of a base case, or a stopping condition
// recursive case vs base case -> one continues the call to the function, the other gets out of the function

func recursive(counter int) {
	if counter < 10 {
		counter++
		fmt.Println(counter)
		recursive(counter)
	} else {
		fmt.Println("Base case achieved")
	}
}

//O(n) time complexity
func findFactorialIterative(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res = res * i
	}
	return res
}

// O(n) time complexity
func findFactorialRecursive(n int) int {
	res := 1
	if n == 2 {
		return 2
	} else {
		res = n * findFactorialRecursive(n-1)
	}

	return res
}

// O(n)
func fibonacciIterative(n int) int {
	if n < 2 {
		return n
	} else {
		var a, b int = 0, 1
		for i := 2; i <= n; i++ {
			a, b = b, a+b
		}
		return b
	}
}

// O(2^n)
func fibonacciRecursive(n int) int {
	if n < 2 {
		return n
	} else {
		return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
	}
}

func main() {
	fmt.Println(findFactorialRecursive(5))

	fmt.Println(fibonacciRecursive(2))
}

// 0 1 1 2 3 5
