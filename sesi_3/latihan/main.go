package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

func main() {
	now := time.Now()
	asynchornous()
	// synchronous()
	log.Println("Done in", time.Since(now).Seconds())
}

func asynchornous() {
	// chunk file & read file

	users, _ := readFileChunkConcurrent("./data.json")

	dataCh := usdToIdrChunkConcurrent(users)

	done := make(chan bool)

	writeToFileConcurrent(dataCh, done)

	if <-done {
		log.Println("Donee")
	}

	// write file
}

func usdToIdrChunkConcurrent(users <-chan []User) <-chan []User {
	usersCh := make(chan []User)
	usersData := []User{}

	go func() {
		for u := range users {
			wg := sync.WaitGroup{}
			for _, v := range u {
				wg.Add(1)
				newData := v
				newData.Salary *= 15_000
				usersData = append(usersData, newData)
				wg.Done()
			}
			wg.Wait()
			usersCh <- usersData
		}
		close(usersCh)
	}()

	return usersCh
}

func readFileChunkConcurrent(filename string) (<-chan []User, error) {
	usersCh := make(chan []User)
	now := time.Now()
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	users := []User{}

	err = json.Unmarshal(dataByte, &users)
	if err != nil {
		return nil, err
	}

	chunkSize := 125

	go func() {
		for i := 0; i < len(users); i += chunkSize {
			end := i + chunkSize
			if end > len(users) {
				end = len(users)
			}

			usersCh <- users[i:end]
		}
		close(usersCh)
	}()

	log.Println("success read data in", time.Since(now).Seconds(), "s")
	return usersCh, nil
}

func writeToFileConcurrent(dataCh <-chan []User, done chan bool) {
	wg := sync.WaitGroup{}

	for data := range dataCh {
		wg.Add(1)
		go func(data []User) {
			for _, v := range data {
				user, _ := json.Marshal(v)
				err := os.WriteFile("./users/"+v.Name+".json", user, 0666)
				if err != nil {
					log.Println("error when try to write file", err)
				}
			}
			wg.Done()
		}(data)
	}

	go func() {
		wg.Wait()
		done <- true
	}()
}
