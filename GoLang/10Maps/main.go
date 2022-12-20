package main

import "fmt"

func main() {
	fmt.Println("maps in Golang")

	// key value pairs and retrival is fast

	student := make(map[int]string)

	student[1290] = "supraja"
	student[1291] = "Tom"
	student[1292] = "Bob"
	student[1293] = "Emma"

	fmt.Println("students are ", student)
	// print name of student with rollno 1290
	fmt.Println(student[1290])

	// delete rollno 1293
	delete(student, 1293)
	fmt.Println(student)

	for key, val := range student {
		fmt.Printf("%v  rollno is %v\n", key, val)

	}

	// var supraja = User("supraja",1290,"IT",76)

}
