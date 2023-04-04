

/* 
In this example, we create a variable number of goroutines.
The number of goroutines is given as a command-line argument.

Simply enter the following prompt: 'go run multiple.go' followed by the number 
of goroutines you wish to create and see the output in the cli.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	count := 10
	if len(os.Args) == 1 {
		fmt.Println("Using default value.")
	} else {
		temp, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Using default value.")
		} else {
			count = temp
		}
	}

	fmt.Printf("Going to create %d goroutines.\n", count)
	// You can create multiple goroutines, especially when you want to create lots of them.	
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}
	// Below is the time.Sleep that delays the termination of the main() function.
	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}