package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strconv"
	"sync"
	"time"
)

var messages_sent int = 0
var messages_received int = 0
var comparisions int = 0

func process(id int, value int, wg *sync.WaitGroup, prev *chan int, next *chan int, final *[]int, turn int) {
	defer wg.Done()
	// if turn == 0 send right, if turn == 1 receive left right, calculate and send left rightm if right send left

	var final_value int = value

	if turn == 0 {
		if next != nil {
			// send(next, value)
			*next <- value
			messages_sent++
			final_value = <-*next
			messages_received++
		}
	} else if turn == 2 {
		if prev != nil {
			// send(prev, value)
			*prev <- value
			messages_sent++
			final_value = <-*prev
			messages_received++
		}
	} else {
		if prev != nil && next != nil {
			prev_value := <-*prev
			next_value := <-*next
			messages_received += 2

			ints := []int{prev_value, value, next_value}
			slices.Sort(ints)
			comparisions += 3

			final_value = ints[1]
			*prev <- ints[0]
			*next <- ints[2]
			messages_sent += 2
		} else if prev != nil {
			prev_value := <-*prev
			messages_received++
			ints := []int{prev_value, value}
			slices.Sort(ints)
			comparisions += 2
			final_value = ints[1]
			*prev <- ints[0]
			messages_sent++
		} else if next != nil {
			next_value := <-*next
			messages_received++
			ints := []int{next_value, value}
			slices.Sort(ints)
			comparisions += 2
			final_value = ints[0]
			*next <- ints[1]
			messages_sent++
		}
	}

	(*final)[id] = final_value
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
				go process(i, arr[i], &wg, nil, &channels[i], &arr, (i+j)%3)
			} else if i == N-1 {
				go process(i, arr[i], &wg, &channels[i-1], nil, &arr, (i+j)%3)
			} else {
				go process(i, arr[i], &wg, &channels[i-1], &channels[i], &arr, (i+j)%3)
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
	fmt.Println("==========================")
	var args = os.Args
	if len(args) >= 2 {
		var err error
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
	}

}
