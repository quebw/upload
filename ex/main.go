package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
)

func printSlice(s []int) {
	log.Printf("ex6\nlen=%d\ncap=%d\n%v\n", len(s), cap(s), s)
}

type User struct {
	name string
	age int64
	gender bool
	address string
  }

   	func (u*User) GetName() string{
		   return u.name
	   }
   	func (u*User) GetAge() int{
		   return int(u.age)
	   }
   	func (u*User) GetGender() bool{
		   return u.gender
	   }
  	func (u*User) GetAddress() string{
		  return u.address
	  }


 var m =  map[string]User {
	"person1": {"a", 15 , true, "ab"},
	"person2": {"b",16 , false, "bc"},
}

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
	log.Println("ex1\n",sv)





	// ex3 loiiiii
		values := "admin"
		 idx := sort.Search(len(sv), func(i int) bool {
		 	return string(sv[i].Name) >= values
		 })

		 if sv[idx].Name == values {
		 	log.Println("ex3\nFound:", idx, sv[idx])
		 } else {
		 	log.Println("ex3\nFound nothing ")
		 }
		
	 

	//  ex4
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
		log.Println("ex4\n",string(result))
    	// log.Println("RECORDS:", sv)



		// ex5
		log.Printf("ex5\nAddress of elements:\n e[1] %p\n e[2] %p\n e[3] %p\n", &sv[0], &sv[1], &sv[2])



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
	
	log.Println("ex7\n",m)
	
}
