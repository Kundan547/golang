package main

import "fmt"

// Name and Age hold basic user information
var (
	Name = "John"
	Age  = 30
)

// PrintName prints the current value of Name
func PrintName() {
	fmt.Println("Name:", Name)
}

// PrintAge prints the current value of Age
func PrintAge() {
	fmt.Println("Age:", Age)
}

func main() {
	PrintName()
	PrintAge()
}
