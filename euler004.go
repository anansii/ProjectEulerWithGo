/*  Problem 4 (https://projecteuler.net/problem=4)

    A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
    Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	//"strconv"
)

func main() {

	fmt.Println("Problem 004")

	maxPalindrome := 1
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			product := i * j
			if isPalindrome(product) && product > maxPalindrome {
				maxPalindrome = product
			}
		}
	}

	fmt.Println(maxPalindrome)

}

func isPalindrome(input int) bool {
	return input == reverse(input)
}

func reverse(num int) int {
	reversed := 0
	for num > 0 {
		reversed *= 10
		reversed += num % 10
		num /= 10
	}
	return reversed
}
