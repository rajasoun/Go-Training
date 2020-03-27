package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"sync"
)

// For Windows
// func getRandomIterCount() int {
// 	randomNumber := rand.Intn(100)

// 	iterCount := 300 + randomNumber

// 	return iterCount
// }

// For Unix
func getRandomIterCount() int {
	randomFile, _ := os.Open("/dev/random")
	randomReader := bufio.NewReader(randomFile)
	randomNumber, _ := rand.Int(randomReader, big.NewInt(100))
	iterCount := int(300 + randomNumber.Int64())

	return iterCount
}

func busyWork(iterCount int) {
	for i := 0; i < iterCount; i++ {
		for j := 0; j < iterCount; j++ {
			mul := []byte(strconv.Itoa(i * j))
			ioutil.WriteFile("/dev/null", mul, 0644)
		}
	}
}

func say(s string, index int) {
	iterCount := getRandomIterCount()

	busyWork(iterCount)

	fmt.Printf("Hello, %s! from %d goroutine; Iter Count: %d\n", s, index+1, iterCount)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 11; i++ {
		wg.Add(1)
		go func(k int) {
			say("world", k)
			wg.Done()
		}(i)
	}

	wg.Wait()
	// fmt.Println("Started all say goroutines!")
	// time.Sleep(10 * time.Second)
}
