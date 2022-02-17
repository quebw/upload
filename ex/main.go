package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
)

// type Serve []map[string]string
func printSlice(s []int) {
	log.Printf("\nlen=%d\ncap=%d\n%v\n", len(s), cap(s), s)
}

type User struct {
	name string
	age int64
	gender bool
	address string
  }

func (u*User) GetName() string
func (u*User) GetAge() int
func (u*User) GetGender() bool
func (u*User) GetAddress() string


func main() {

	// ex2
	type Serve struct{
		Name string `json:"name"`
		Class string `json:"class"`
	}

	var sv []Serve
	file, err := ioutil.ReadFile("servve.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &sv)
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("Name:%s, Class:%s", sv.Name, sv.Class)
	log.Println(sv)





	// ex3
		needle := "admin"
		idx := sort.Search(len(sv), func(i int) bool {
			return string(sv[i].Name) >= needle
		})

		if sv[idx].Name == needle {
			log.Println("Found:", idx, sv[idx])
		} else {
			log.Println("Found noting ")
		}


	 

	//  ex4
	//  Name := "fileCustome"
	//  Class := "org.cofax.cds.FileServlet.Custome"
		input := []byte(`[{
			"name": "fileCustome",
			"class": "org.cofax.cds.FileServlet.Custome"
		}]`)
		var add []Serve
		err = json.Unmarshal(input, &add)
			if (err != nil) {
				log.Fatal(err)
			}
			sv = append(sv, add...)  
			for _, v := range sv {
				log.Println(v)
			}
			log.Println()
			result, err := json.Marshal(sv)
			if err != nil {
				log.Println(err)
			}
			log.Println(string(result))
    	// log.Println("RECORDS:", sv)



		// ex5
		log.Printf("Address of elements:\n e[1] %p\n e[2] %p\n e[3] %p\n", &sv[0], &sv[1], &sv[2])



	// ex6
	s := []int{11,34,56,77,99,109,66,20, 88, 34}
	sort.Ints(s)
	printSlice(s)

	slc1 := make([]int, 7)
	slc2 := make([]int, 15)
	copy_1 := copy(slc1,s)
	copy_2 := copy(slc2,slc1)
    log.Println("\nSlice:", slc1)
    log.Println("Total number of elements copied:", copy_1)
	log.Println("\nSlice:", slc2)
    log.Println("Total number of elements copied:", copy_2)
	//khong xay ra loi


	// ex7
    value := []string{"a", "b", "c", "d"}
	log.Println(value)

	valueMap := make(map[string]*User)
	for i := 0; i < len(value); i++ {  
		valueMap[i] = value[i]  
	   }  
	   log.Println(valueMap) 
}