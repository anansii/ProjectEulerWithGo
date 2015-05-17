/*  Problem 3 (https://projecteuler.net/problem=3)

    The prime factors of 13195 are 5, 7, 13 and 29.
    What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Problem 003")
	input := uint64(600851475143)

	fmt.Printf("  Input = %d\n", input)
	collectedPrimes := collectPrimes(sieve(input))
	fmt.Printf(" Primes = %d\n", collectedPrimes)
	collectedFactors := collectFactors(collectPrimes(sieve(input)), input)
	fmt.Printf("Factors = %d\n", collectedFactors)
	fmt.Printf("The largest prime factor of the number %d is %d\n", input, collectedFactors[len(collectedFactors)-1])

}

func collectFactors(primes []uint64, input uint64) []uint64 {
	factors := make([]uint64, 0)
	for i := uint64(0); i < uint64(len(primes)); i++ {
		if input%primes[i] == 0 {
			factors = append(factors, uint64(primes[i]))
		}
	}
	return factors
}

func collectPrimes(sieve []bool) []uint64 {
	primes := make([]uint64, 0)
	for i := 0; i < len(sieve); i++ {
		if sieve[i] {
			primes = append(primes, uint64(i))
		}
	}
	return primes
}

func sieve(max uint64) []bool {
	// max only needs to go up to the sqrt of max
	max = uint64(math.Floor(math.Sqrt(float64(max))))

	sieve := make([]bool, max)

	for i := 2; i < len(sieve); i++ {
		sieve[i] = true
	}

	for i := 2; i < len(sieve); i++ {
		if sieve[i] {
			for j := i + i; j < len(sieve); j += i {
				sieve[j] = false
			}
		}
	}

	return sieve
}
