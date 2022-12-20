package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")
	supraja := student{"SUPRAJA", 89, "IT"}
	fmt.Println(supraja.Marks)
	fmt.Println("Marks :", student.GetMarks(supraja))
	fmt.Println("Marks :", student.SetMarks(supraja))
	//  the above line doesn't change the marks but just changes the value in the func
	fmt.Println("Marks :", student.GetMarks(supraja))

}

type student struct {
	Name   string
	Marks  int
	Branch string
}

func (S student) GetMarks() int {
	return S.Marks
}
func (s student) SetMarks() int {
	s.Marks = 99
	return s.Marks
}
