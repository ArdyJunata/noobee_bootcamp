package main

import (
	"log"
	"math/rand"
	"time"
)

var counter int
var userToken map[string]time.Time

func main() {
	ball := make(chan int)
	done := make(chan bool)

	go play(ball, done, "Reyhan")
	go play(ball, done, "NooBeeID")
	ball <- 1

	if <-done {
		log.Println("Done ...")
	}
}

func play(ball chan int, done chan bool, name string) {
	for {

		userToken["token"] = time.Now()
		counter++

		time.Sleep(1 * time.Second)
		// player 1 hit ball to player 2
		val := <-ball
		log.Println("player", name, "got value", val)
		if val%11 == 0 {
			log.Println("player", name, "fail in value", val)
			done <- true
			break
		}

		// proses / teknik pukulan player
		// val++
		val = rand.Intn(100-1) + 1

		log.Println("player", name, "hit the ball to another player with value", val)
		// ball will be delivered to player 2
		ball <- val
	}
}
