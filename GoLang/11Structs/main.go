package main

import "fmt"

func main() {
	fmt.Println("structs in Golang")
	// now inheritance in go lang but we have interface
	// interface: inheriting from different parent class by the child

	student := User{"supraja", 1290, "IT", 78.9}
	fmt.Println(student.Name)
	fmt.Printf("student details are :%+v ", student)

}

type User struct {
	Name   string
	Rollno int
	Branch string
	Marks  float32
}
