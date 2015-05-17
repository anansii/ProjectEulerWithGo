/*  Problem 5 (https://projecteuler.net/problem=5)

    2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
    What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Problem 005")

	old, candidate, needle := 0, 0, 0

	for needle == 0 {
		for i := 1; i <= 10; i++ {
			if candidate%i != 0 {
				candidate++
			}
		}
		if old == candidate {
			needle = candidate
		} else {
			old = candidate
		}
	}

	fmt.Println(needle)

}
