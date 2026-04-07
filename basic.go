package main

import "fmt"

/* ⚙️ Task 3 (If-Else)
👉 Write a program:
variable age := 16 (or any)
if age >= 18 → print "Adult"
else → "Minor" */

func main() {
	age := 13

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}
