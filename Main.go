package main

import "fmt"

func main() {
	// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
	appleCount := applePrice * 9
	pearCount := float64(pearPrice) * 8
	res := appleCount + pearCount
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?\n", "Відповідь: ", res)
	// 2. Скільки груш ми можемо купити?
	res2 := float64(cardBalance) / applePrice
	fmt.Println("Скільки груш ми можемо купити?\n", "Відповідь: ", res2)
	// 3. Скільки яблук ми можемо купити?
	res3 := float64(cardBalance) / pearPrice
	fmt.Println("Скільки яблук ми можемо купити?\n", "Відповідь: ", res3)
	// 4. Чи ми можемо купити 2 груші та 2 яблука?
	res4 := (applePrice * 2) + (pearPrice * 2)
	if float64(cardBalance) > res4 {
		fmt.Println("Так, можем")
	} else {
		fmt.Println("Ні, не можем")
	}

}

const (
	applePrice  = 5.99
	pearPrice   = 7
	cardBalance = 23
)
