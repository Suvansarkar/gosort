package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"sync"
	"time"
)

var messages_sent int = 0
var messages_received int = 0
var comparisions int = 0

func process(id int, value int, wg *sync.WaitGroup, prev *chan int, next *chan int, final *[]int, round int) {
	defer wg.Done()
	var odd bool = id%2 == 1
	// swap function to do
	if round%2 == 0 {
		if odd {
			if next != nil {
				value = send(next, value)
				messages_sent++
			}
		} else {
			if prev != nil {
				value = recieve(prev, value)
				messages_received++
			}
		}
	} else {
		if odd {
			if prev != nil {
				value = recieve(prev, value)
				messages_received++
			}
		} else {
			if next != nil {
				value = send(next, value)
				messages_sent++
			}
		}
	}

	(*final)[id] = value
}

func send(ch *chan int, value int) int {
	*ch <- value
	var recv int = <-*ch
	if recv < value {
		value = recv
	}
	comparisions++
	return value
}

func recieve(ch *chan int, value int) int {
	recv := <-*ch
	*ch <- value
	if recv >= value {
		value = recv
	}
	comparisions++
	return value
}

func randArray(n int) []int {
	var arr = make([]int, n)
	for i := range arr {
		arr[i] = rand.IntN(n + 10)
	}
	return arr
}

func run(N int) {
	var arr []int = randArray(N)
	fmt.Println("N = ", N)
	fmt.Println("-----------------------")
	fmt.Println("Random array generated:", arr)
	var channels = make([]chan int, N-1)
	for i := range channels {
		channels[i] = make(chan int)
	}

	// Only start after all channels are created and arrays have been generated
	start := time.Now()

	var wg sync.WaitGroup
	for j := 0; j < N-1; j++ {
		for i := 0; i < N; i++ {
			wg.Add(1)
			if i == 0 {
				go process(i, arr[i], &wg, nil, &channels[i], &arr, j)
			} else if i == N-1 {
				go process(i, arr[i], &wg, &channels[i-1], nil, &arr, j)
			} else {
				go process(i, arr[i], &wg, &channels[i-1], &channels[i], &arr, j)
			}
		}
		wg.Wait()
	}

	fmt.Println("Array after ", N-1, " passes: ", arr)
	fmt.Println("Messages sent: ", messages_sent)
	fmt.Println("Messages received: ", messages_received)
	fmt.Println("Comparisions: ", comparisions)
	fmt.Println("Execution time: ", time.Since(start))
	fmt.Println("==========================")
}

func main() {
	var args = os.Args
	fmt.Println("==========================")
	if len(args) >= 2 {
		N, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		run(N)
	} else {
		run(10)
		run(20)
		run(30)
		run(50)
	}

}
