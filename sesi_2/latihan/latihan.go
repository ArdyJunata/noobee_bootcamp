package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name string
	Hit  int
}

const BreakPoint int = 11

func main() {
	players := []string{"Budi", "Tono"}

	done := make(chan *Player)
	turn := make(chan *Player)

	for _, v := range players {
		go play(v, turn, done)
	}

	turn <- &Player{
		Hit: 1,
	}

	finish(done)
}

func play(playerName string, turn, done chan *Player) {
	for {
		p := <-turn

		p.Name = playerName

		if p.Hit%11 == 0 {
			done <- p
			break
		}

		p.Hit = rand.Intn(100-1) + 1
		fmt.Println(p.Name, "memukul bola dengan hit =", p.Hit)

		time.Sleep(200 * time.Millisecond)

		turn <- p
	}
}

func finish(done chan *Player) {
	for {
		p := <-done
		fmt.Println(p.Name, "kalah setelah menerima pukulan =", p.Hit)
		break
	}
}
