package main

import "fmt"

func main() {
	fmt.Println("2222")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("bug-fix")

<<<<<<< HEAD
	fmt.Println("master-fix")
=======
	fmt.Println("hot-fix")
>>>>>>> hot-fix
}
