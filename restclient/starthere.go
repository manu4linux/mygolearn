package main

import "fmt"

func main() {
	n, err := fmt.Printf("hello world\n")
	fmt.Printf("%d\n", n)
	fmt.Println(err)
}
