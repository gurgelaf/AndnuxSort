package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Main

func main() {
	var countN, maxr int

	fmt.Printf("Welcome to Andnux Sort\n type the number of elements:")
	fmt.Scan(&countN)
	fmt.Println("Type the maximum value:")
	fmt.Scan(&maxr)

	start := time.Now()

	n := make([]int, countN)

	rand.Seed(time.Now().UnixNano())

	fmt.Println("adding random values...")

	for i := 0; i < countN; i++ {
		n[i] = rand.Intn(maxr - 1)
	}

	fmt.Println("sorting...")

	andnuxSort(n, countN)

	fmt.Println("The values ware sorted!")

	forr(n, countN)

	fmt.Println("\nDisplayed values")

	fmt.Println("run time: ", time.Since(start))
}

//Andnux Sort

func andnuxSort(n []int, countN int) {
	var j, vM, max int = 0, -1, max(n, countN)
	aux, count := make([]int, max+1), make([]int, max+1)

	for i := 0; i < max; i++ {
		aux[i] = vM
	}

	for i := 0; i < countN; i++ {
		if n[i] == aux[n[i]] {
			count[aux[n[i]]]++
		}
		if n[i] != vM {
			aux[n[i]] = n[i]
		}
	}

	for i := 0; i <= max; i++ {
		if aux[i] != vM {
			if count[aux[i]] > 0 {
				for k := 0; k <= count[aux[i]]; k++ {
					n[j] = aux[i]
					j++
				}
			} else {
				n[j] = aux[i]
				j++
			}
		}
	}
}

//Secondary functions

func max(n []int, countN int) int {
	max := 0
	for i := 0; i < countN; i++ {
		if max < n[i] {
			max = n[i]
		}
	}
	return max
}

func forr(n []int, countN int) {
	for i := 0; i < countN; i++ {
		if (i+1)%10 == 0 {
			fmt.Println(n[i], "\t")
		} else {
			fmt.Print(n[i], "\t")
		}
	}
}
