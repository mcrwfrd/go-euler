package main

import "fmt"

func main() {
	f2 := 1
	f1 := 2
	f0 := 0
	sum := 2

	for f0 < 4000000 {
		f0 = f1 + f2

		if (f0 % 2 == 0) {
			fmt.Printf("adding: %d\n", f0)
			sum += f0
		}

		f2 = f1
		f1 = f0
	}

	fmt.Printf("%d\n", sum)
}
