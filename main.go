package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type User struct {
	name    string
	age     int64
	gender  bool
	address string
}

func (u *User) GetName() string {
	return u.name
}
func (u *User) GetAge() int64 {
	return u.age
}
func (u *User) GetGender() bool {
	return u.gender
}
func (u *User) GetAddress() string {
	return u.address
}

type Serve struct {
	Name  string `json:name`
	Class string `json:class`
}

var serves []*Serve

func ex2() {
	file, err := ioutil.ReadFile("serve.json")
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(file, &serves)
	if err != nil {
		log.Println(err)
	}
	substr := "Admin"
	for i, v := range serves {
		log.Println(i, v)

		check_1 := strings.Contains(v.Name, substr)
		check_2 := strings.Contains(v.Class, substr)
		if check_1 == true || check_2 == true {
			log.Println(substr,v)
		} else {
			log.Println("Cannot found")
		}
	}

}

// func ex3() {

// }

func ex4() []*Serve {
	input := Serve{Name: "fileCustome", Class: "org.cofax.cds.FileServlet.Custome"}
	new := append(serves, &input)
	return new
}

func ex6() {
	slice := []int{11, 34, 56, 77, 99, 109, 66, 20, 88, 34}
	log.Println(slice)
	sort.Ints(slice)
	log.Println(slice)
	slice1 := make([]int, 7)
	copy_1 := copy(slice1, slice)
	log.Println(slice1, copy_1)
	copy_2 := copy(slice, slice[1:15])
	log.Println(copy_2)
}

func ex7() {
	// users := User{}
	u1 := &User{name: "A", age: 20, gender: false, address: "abc"}
	u2 := &User{name: "B", age: 21, gender: true, address: "abc"}
	u3 := &User{name: "C", age: 20, gender: false, address: "abc"}
	u4 := &User{name: "D", age: 21, gender: true, address: "abc"}

	new := map[string]*User{u1.GetName(): u1, u2.GetName(): u2, u3.GetName(): u3, u4.GetName(): u4}
	for i, value := range new {
		log.Println(i, value)
	}
}

func main() {
	ex2()
	// log.Println("------------------Ex3--------------------")
	// ex3()
	log.Println("------------------Ex4--------------------")
	input := ex4()
	for i, v := range input {
		log.Println(i, v)
	}
	log.Println("------------------Ex5--------------------")
	address := ex4()
	log.Println(address)
	log.Println("------------------Ex7--------------------")
	ex7()
	log.Println("------------------Ex6--------------------")
	ex6()

}
