package main

import (
	"context"
	"log"
	"strconv"
	"time"
)

func ex_1() {
	for i := 1; i <= 3; i++ {
		log.Println("time now: {milliseconds}")
		time.Sleep(3 * time.Second)
	}
	log.Println("ket thuc")
}

func ex_2() {
	timestamp := time.Now().Unix()
	log.Println("Unix timestamp is:", timestamp)
	now := time.Now()
	log.Println("Current timestamp is:", now)
}

func ex_3(ctx context.Context) error {
	ctx, cancle := context.WithTimeout(ctx, 3*time.Second)
	defer cancle()
	for i := 0; i <= 3; i++ {
		select {
		case <-ctx.Done():
			log.Println("Time out!")
			return ctx.Err()
		default:
			time.Sleep(3 * time.Second)
			log.Println("No :", i)
		}
	}
	log.Println("All done")
	return nil
}

func ex_4() {
	unixTimestamp := "1592190294764144364"
	unixIntValue, err := strconv.ParseInt(unixTimestamp, 10, 64)
	if err != nil {
		log.Println(err)
	}
	timeStamp := time.Unix(unixIntValue, 0)
	// log.Println(timeStamp)
	hr, min, sec := timeStamp.Clock()
	log.Printf("Clock: [%d]hour : [%d]minutes : [%d]second", hr, min, sec)
	log.Println("number_of_minutes of unixTimestamp :", hr*60+min+sec/60)

}

func ex_5() {
	unixTimestamp := "1592190385"
	unixIntValue, err := strconv.ParseInt(unixTimestamp, 10, 64)
	if err != nil {
		log.Println(err)
	}
	newTime := time.Unix(unixIntValue, 0)
	strDate := newTime.Format(time.UnixDate)
	log.Println("Number :", newTime.UTC())
	log.Println("String :", strDate)
}

func x(ctx context.Context) {
	time_now := time.Now().UnixNano()

	select {
	case <-time.After(3 * time.Second):
		now := time.Now().UnixNano() - time_now
		log.Println("Result:", now)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}

func ex_7() {
	// time_now := time.Now().UnixNano()

	// ctx, cancel := context.WithCancel(context.Background())
	// select {
	// case <-time.After(3 * time.Second):
	// 	now := time.Now().UnixNano() - time_now
	// 	log.Println("Result:", now)
	// case <-ctx.Done():
	// 	log.Println(ctx.Err())
	// }
	// cancel()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	x(ctx)
	cancel()
}

func ex_8() {
	ticker := time.NewTicker(100 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				log.Println("${time.Now().Unix()} done", t)
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	ticker.Stop()
	done <- true
	log.Println("Done!")

}

func ex_9() {
	required_Time := time.Duration(100) * time.Millisecond
	// log.Println(required_Time)
	f := func() {
		log.Println("im study")
	}
	Timer1 := time.AfterFunc(required_Time, f)
	defer Timer1.Stop()
	time.Sleep(2 * time.Second)
}

func main() {
	log.Println("----------------------Ex1----------------------")
	ex_1()
	log.Println("----------------------Ex2----------------------")
	ex_2()
	log.Println("----------------------Ex3----------------------")
	ctx := context.Background()
	err := ex_3(ctx)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Success!")
	}
	log.Println("----------------------Ex4----------------------")
	ex_4()
	log.Println("----------------------Ex5----------------------")
	ex_5()
	log.Println("----------------------Ex6----------------------")
	log.Println("Cac moc don vi: yyyy-mm-dd hh:mm:ss + nsec nanoseconds")
	log.Println("----------------------Ex7----------------------")
	ex_7()
	log.Println("----------------------Ex8----------------------")
	ex_8()
	log.Println("----------------------Ex9----------------------")
	ex_9()

}
