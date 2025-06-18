package arrays01

import "fmt"

func Array01Func() {
	ar := [...]int{1, 2, 3} //array
	ar1 := []int{1, 2, 3}   //slice

	fmt.Println("ar ", ar)
	fmt.Println("ar1 slice", ar1)
}
