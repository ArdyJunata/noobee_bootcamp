package main

import "fmt"

type Student struct {
	Name  string
	Class string
}

func main() {
	var student1 Student
	student1.SetMyName("naats")
	message := student1.CallMyName()

	fmt.Println(message)
}

func (s *Student) SetMyName(name string) {
	s.Name = name
}

func (s Student) CallMyName() (message string) {
	return fmt.Sprintf("My Name is %s", s.Name)
}
