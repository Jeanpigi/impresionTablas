package main

import "fmt"

func main() {
	var n, k, multi, divi int

	fmt.Println("Escriba el primer valor a multiplicar")
	fmt.Scanf("%d", &n)

	fmt.Println("Escriba el segundo valor ")
	fmt.Scanf("%d", &k)

	for i := 1; i <= k; i++ {
		multi = n * i
		divi = multi / i
		fmt.Printf("El resultado de la multiplicación de %d por %d es %v\n ", n, i, multi)
		fmt.Printf("El resultado de la división de %d por %d es %v\n", multi, i, divi)
	}
}
