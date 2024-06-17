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

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("hello.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	//  get file info
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File name: %s\n", info.Name())
	// fmt.Printf(info.Size())
}
