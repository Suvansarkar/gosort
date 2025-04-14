package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"sync"
)

func process(id int, value int, wg *sync.WaitGroup, prev *chan int, next *chan int, final *[]int, round int) {
	defer wg.Done()
	var odd bool = id%2 == 1
	// swap function to do
	if round%2 == 0 {
		if odd {
			if next != nil {
				value = send(next, value)
			}
		} else {
			if prev != nil {
				value = recieve(prev, value)
			}
		}
	} else {
		if odd {
			if prev != nil {
				value = recieve(prev, value)
			}
		} else {
			if next != nil {
				value = send(next, value)
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
	return value
}

func recieve(ch *chan int, value int) int {
	recv := <-*ch
	*ch <- value
	if recv >= value {
		value = recv
	}
	return value
}

func randArray(n int) []int {
	var arr = make([]int, n)
	for i := range arr {
		arr[i] = rand.IntN(n + 10)
	}
	return arr
}

func main() {
	var args = os.Args
	var N int
	if len(args) >= 2 {
		var err error
		N, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	} else {
		N = 10
	}

	var arr []int = randArray(N)
	fmt.Println("Random array generated:", arr)

	var wg sync.WaitGroup
	for j := 0; j < N-1; j++ {
		var channels = make([]chan int, N-1)
		for i := 0; i < N; i++ {
			wg.Add(1)
			if i != N-1 {
				channels[i] = make(chan int)
			}
			if i == 0 {
				go process(i, arr[i], &wg, nil, &channels[i], &arr, j)
			} else if i == N-1 {
				go process(i, arr[i], &wg, &channels[i-1], nil, &arr, j)
			} else {
				go process(i, arr[i], &wg, &channels[i-1], &channels[i], &arr, j)
			}
		}
		wg.Wait()
		fmt.Println("Pass: ", j+1)
	}

	fmt.Println("Array after ", N-1, " passes: ", arr)
}
