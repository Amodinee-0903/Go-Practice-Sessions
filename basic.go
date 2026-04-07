package main

import "fmt"

/* ⚙️ Task 5: Slices (VERY IMPORTANT)

Create a slice of integers: [1,2,3,4,5]
Loop through it
Print each value

⚙️ Task 6: Slight Twist
👉 From that slice:print only even numbers */

func main() {
	a := []int{1, 2, 3, 4, 5}
	for _, k := range a {
		if k%2 == 0 {
			fmt.Println(k)
		}
	}
}
