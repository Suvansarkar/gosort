package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"sync"
	"time"
)

// | ID  | Final Value | Messages Sent | Messages Recv | Comparisons | Exec Time (ms) | Round Values |

var messages_sent int = 0
var messages_received int = 0
var comparisions int = 0

func process(id int, value values, wg *sync.WaitGroup, prev *chan int, next *chan int, final *[]values, round int) {
	defer wg.Done()

	// if first element, only check with last. Same with last element
	// Place data on right channels first, then check left channels, then send on left channels and check right channels
	if id == 0 {
		*next <- value.right
		*next <- value.marked
		recv := <-*next
		marked := <-*next
		if recv < value.right {
			value.right = recv
			value.marked = marked
		}
	} else {
		if id%2 == 0 {

		} else {
			*next <- value.right
			*next <- value.marked
			recv1 := <-*prev
			marked1 := <-*prev
			*prev <- value.left
			*prev <- value.marked
			recv2 := <-*next
			marked2 := <-*next
			if recv1 > value.right {
				value.right = recv1
				value.marked = marked1
			}
			if recv2 < value.right {
				value.right = recv2
				value.marked = marked2
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

func randArray(n int) []values {
	var arr = make([]values, n)
	for i := range arr {
		var randInt int = rand.IntN(n + 10)
		if i == 0 {
			arr[i].right = randInt
			arr[i].marked = 1
			arr[i].a = -1
		} else if i == n-1 {
			arr[i].left = randInt
			arr[i].marked = -1
			arr[i].a = 1
		} else {
			arr[i].left = randInt
			arr[i].right = randInt
		}
	}
	return arr
}

type values struct {
	right  int
	left   int
	marked int // 0 = not marked, -1 = left, 1 = right
	a      int // I don't remember what a is but it is important. trust.
}

func main() {
	start := time.Now()

	var args = os.Args
	var N int = 10
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

	var arr []values = randArray(N)
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
	}

	fmt.Println("Array after ", N-1, " passes: ", arr)
	fmt.Println("Messages sent: ", messages_sent)
	fmt.Println("Messages received: ", messages_received)
	fmt.Println("Comparisions: ", comparisions)
	fmt.Println("Execution time: ", time.Since(start))
}
