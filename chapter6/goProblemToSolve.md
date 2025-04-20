techupskills: Explain how to code a function in Go by giving me a problem to complete.

GitHub Copilot: Sure! Let's create a simple function in Go that calculates the factorial of a given number. Here's the problem statement:

**Problem Statement:**
Write a function named `Factorial` in Go that takes an integer `n` as input and returns the factorial of `n`. The factorial of a non-negative integer `n` is the product of all positive integers less than or equal to `n`. The factorial of 0 is defined as 1.

**Steps to solve the problem:**
1. Define the function `Factorial` that takes an integer `n` as an argument.
2. Check if `n` is 0. If it is, return 1.
3. Initialize a variable to store the result and set it to 1.
4. Use a loop to multiply the result by each integer from 1 to `n`.
5. Return the result.

Here's the code in Go:

```go
package main

import (
    "fmt"
)

// Factorial function calculates the factorial of a given number n
func Factorial(n int) int {
    if n == 0 {
        return 1
    }
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}

func main() {
    // Test the Factorial function
    fmt.Println(Factorial(5)) // Output: 120
    fmt.Println(Factorial(0)) // Output: 1
    fmt.Println(Factorial(7)) // Output: 5040
}
```

You can run this code in your Go environment to see the results. The `Factorial` function will correctly compute the factorial of the given number `n`.
