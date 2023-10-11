package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := map[string]string{
		"Name":    "NooBee",
		"Class":   "Backend Intermediate",
		"Address": "Jakarta",
	}

	wg := sync.WaitGroup{}

	for k, v := range arr {
		wg.Add(1)
		go func(k, v string) {
			fmt.Printf("Key : %s Value : %s \n", k, v)
			wg.Done()
		}(k, v)
	}

	wg.Wait()

	// buatlah sebuah function print untuk nge handle hasil seperti dibawah
	// pastikan menggunakan goroutine, agar urutan hasilnya itu bisa berbeda beda

	// case 1
	// Key : Name, Value : NooBee
	// Key : Class, Value : Backend Intermediate
	// Key : Addres, Value : Jakarta

	// case 2
	// Key : Class, Value : Backend Intermediate
	// Key : Name, Value : NooBee
	// Key
}
