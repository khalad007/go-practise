// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello \n World")
// 	fmt.Print("This is \t only print")
// 	fmt.Printf("\"This is print F\"")
// }

// package main

// import "fmt"

// // func main() {
// // 	fmt.Println("Hello world")
// // }

// func main() {
// 	var number int16 = 400
// 	fmt.Println(number)
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	f, err := os.Open("hello.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()
// 	//  get file info
// 	info, err := f.Stat()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("File name: %s\n", info.Name())
// 	// fmt.Printf(info.Size())
// }

// package main

// import "fmt"

// func main() {
// 	// short variable
// 	// a := "This is a variable "

// 	// fmt.Println(a)

// 	// long variable
// 	var a string

// 	a = "This is longer longer variable i think "
// 	fmt.Println(a)
// }

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]

	fmt.Println(s)
}
