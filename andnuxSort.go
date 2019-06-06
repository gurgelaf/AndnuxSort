package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var countN, maxr int

	fmt.Printf("bem vindo ao Andnux Sort\n Digite o número de elementos:")
	fmt.Scan(&countN)
	fmt.Println("Digite o valor máximo possível:")
	fmt.Scan(&maxr)

	start := time.Now()

	n := make([]int, countN)

	rand.Seed(time.Now().UnixNano())

	fmt.Println("adicionando valores aleatórios...")

	for i := 0; i < countN; i++ {
		n[i] = rand.Intn(maxr - 1)
	}

	fmt.Println("Ordenando...")

	andnuxSort(n, countN)

	fmt.Println("Valores Ordenados!")

	forr(n, countN)

	fmt.Println("\nValores Mostrados!")

	fmt.Println("tempo de execução: ", time.Since(start))
}

func andnuxSort(n []int, countN int) {
	var max, min, j, vM int = 0, 0, 0, -1

	for i := 0; i < countN; i++ {
		if max < n[i] {
			max = n[i]
		}
	}

	for i := 0; i < countN; i++ {
		if min > n[i] {
			min = n[i]
		}
	}

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

func forr(n []int, countN int) {
	for i := 0; i < countN; i++ {
		if (i+1)%10 == 0 {
			fmt.Println(n[i], "\t")
		} else {
			fmt.Print(n[i], "\t")
		}
	}
}
