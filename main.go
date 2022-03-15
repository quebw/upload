package main

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Line struct {
	so_dong_hien_tai int
	gia_tri          string
}

var wg sync.WaitGroup
var m sync.Mutex

func ex_1a() {
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
	}()
	log.Print("hello 2")
	time.Sleep(3 * time.Second)
}

func ex_1b() {
	ch1 := make(chan string)
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "hello 3"
	}()
	a := <-ch1
	log.Println(a)
	log.Print("hello 2")
}

func ex_1c() {
	log.Print("hello 1")
	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		wg.Done()
	}()
	wg.Wait()
	log.Print("hello 2")
}

func ex_1d() {
	log.Print("hello 1")
	go func() {
		m.Lock()
		log.Print("hello 3")
		m.Unlock()
	}()
	time.Sleep(1 * time.Second)

	log.Print("hello 2")
}

// func ex_2() {
// 	X := make(map[string]string)
// 	for i := 0; i < 3; i++ {
// 		wg.Add(3)
// 		go func() {
// 			for j := 0; j <= 1000; j++ {
// 				m.Lock()

// 				X["key"] = "value"

// 				log.Println(j, X)
// 				wg.Done()
// 				m.Unlock()

// 			}

// 		}()
// 	}
// 	wg.Wait()

// }

func ex2_add1(X map[string]string) {
	a := "value1_"
	for i := 0; i < 1000; i++ {
		m.Lock()
		add := a + strconv.Itoa(i)
		// int -> string
		X[add] = strconv.Itoa(i)
		m.Unlock()
	}
	wg.Done()
}

func ex2_add2(X map[string]string) {
	a := "value2_"
	for i := 0; i < 1000; i++ {
		m.Lock()
		add := a + strconv.Itoa(i)
		X[add] = strconv.Itoa(i)
		m.Unlock()
	}
	wg.Done()
}

func ex2_add3(X map[string]string) {
	a := "value3_"
	for i := 0; i < 1000; i++ {
		m.Lock()
		add := a + strconv.Itoa(i)
		X[add] = strconv.Itoa(i)
		m.Unlock()
	}
	wg.Done()

}

// ex3
func errFunc() {
	var mt sync.Mutex

	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 10000; j++ {
				mt.Lock()
				if _, ok := m[j]; ok {
					delete(m, j)
					continue
				}
				m[j] = j * 10
				mt.Unlock()
			}
		}()
	}

	log.Print("done")
}

// func ex4() {

// 	data, err := ioutil.ReadFile("file.txt")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	file := string(data)
// 	line := 0
// 	temp := strings.Split(file, "\n")
// 	for _, item := range temp {
// 		log.Println("[", line, "] \t", item)
// 		line++
// 	}
// 	log.Println("xong")
// 	log.Println(runtime.NumGoroutine(), "Goroutines")

// }

func ex4() {
	ch := make(chan string, 10)
	finish := make(chan bool)
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	a := []*Line{}
	line := 0
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ch <- scanner.Text()
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go data(finish, ch, scanner)
			wg.Done()
		}
		wg.Wait()
		time.Sleep(100 * time.Millisecond)
		output := Line{so_dong_hien_tai: line, gia_tri: scanner.Text()}
		temp := append(a, &output)
		line++
		for i := range temp {
			log.Println("No:", temp[i].so_dong_hien_tai, "Output:", temp[i].gia_tri)
		}
		wg.Wait()
	}

	log.Println("xong")
	log.Println(runtime.NumGoroutine(), "Goroutines")

}

func data(finish chan bool, ch chan string, scanner *bufio.Scanner) {
	for range scanner.Text() {
		log.Println(<-ch)
	}
	finish <- false
	close(finish)
	close(ch)
}

func main() {
	log.Println("-------------Ex1---------------")
	ex_1a()
	log.Println("-------------Channel-------------")
	ex_1b()
	log.Println("-------------WaitGroup-----------------")
	ex_1c()
	log.Println("------------------Mutex-----------------------")
	ex_1d()
	log.Println("------------------Ex2-----------------------")
	X := make(map[string]string)
	wg.Add(3)
	go ex2_add1(X)
	go ex2_add2(X)
	go ex2_add3(X)
	wg.Wait()
	count := 0
	for key, value := range X {
		log.Println("Key",key,"Value", value)
		count++
		if count == 15 {
			break
		}
	}
	log.Println("Add value:", count)

	log.Println("------------------Ex3-----------------------")
	errFunc()

	log.Println("------------------Ex4-----------------------")
	ex4()
}
