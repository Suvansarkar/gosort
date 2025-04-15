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

type Element struct {
	Value    int
	IsUnique bool // Mark for tracking minimum and maximum elements
}

type Process struct {
	ID   int
	Area int     // Used to decide which value to select at the end
	VL   Element // Left value
	VR   Element // Right value
}

func process(id int, value *Process, wg *sync.WaitGroup, prev *chan Element, next *chan Element, rounds int) {
	defer wg.Done()

	// First complete all communications
	// Then lets decide what to do with it.

	// All send right, receive left, send left, receive right

	var recvL, recvR Element
	if next != nil {
		*next <- value.VR
		messages_sent++
	}
	if prev != nil {
		recvL = <-*prev
		messages_received++
	}
	if prev != nil {
		*prev <- value.VL
		messages_sent++
	}
	if next != nil {
		recvR = <-*next
		messages_received++
	}

	if recvL.Value > value.VL.Value && id != 0 {
		comparisions++
		value.VL.Value = recvL.Value
		if value.VL.IsUnique {
			value.Area += 1
		}
		if recvL.IsUnique {
			value.VL.IsUnique = recvL.IsUnique
			value.Area -= 1
		}
	}

	if recvR.Value < value.VR.Value && id != rounds-1 {
		comparisions++
		value.VR.Value = recvR.Value
		value.VR.IsUnique = recvR.IsUnique
	}

	if value.VR.Value < value.VL.Value && id != rounds-1 {
		comparisions++
		tempInt := value.VR.Value
		tempBool := value.VR.IsUnique

		value.VR.Value = value.VL.Value
		value.VR.IsUnique = value.VL.IsUnique

		value.VL.Value = tempInt
		value.VL.IsUnique = tempBool
	}
	// (*final)[id] = value
}

func randArray(n int) []Process {
	var arr = make([]Process, n)
	for i := range arr {
		var randInt int = rand.IntN(n + 100)
		if i == 0 {
			arr[i] = Process{
				ID:   i,
				Area: -1,
				VR: Element{
					Value:    randInt,
					IsUnique: true,
				},
			}
		} else if i == n-1 {
			arr[i] = Process{
				ID:   i,
				Area: 0,
				VL: Element{
					Value:    randInt,
					IsUnique: true,
				},
			}
		} else {
			arr[i] = Process{
				ID:   i,
				Area: 0,
				VL: Element{
					Value:    randInt,
					IsUnique: false,
				},
				VR: Element{
					Value:    randInt,
					IsUnique: false,
				},
			}
		}
	}
	return arr
}

func run(N int) {
	var arr []Process = randArray(N)
	fmt.Println("Random array generated:")
	fmt.Print("[ ")
	for i := range arr {
		fmt.Print("[ ", arr[i].VL.Value, arr[i].VR.Value, " ]")
	}
	fmt.Print("] \n")

	// Only start after all channels are created and arrays have been generated
	start := time.Now()

	var wg sync.WaitGroup
	for j := 0; j < N-1; j++ {
		var channels = make([]chan Element, N-1)
		for i := 0; i < N; i++ {
			wg.Add(1)
			if i != N-1 {
				channels[i] = make(chan Element)
			}
			if i == 0 {
				go process(i, &arr[i], &wg, nil, &channels[i], N)
			} else if i == N-1 {
				go process(i, &arr[i], &wg, &channels[i-1], nil, N)
			} else {
				go process(i, &arr[i], &wg, &channels[i-1], &channels[i], N)
			}
		}
		wg.Wait()
	}

	fmt.Println("Array after ", N-1, " passes: ")
	fmt.Print("[ ")
	for i := range arr {
		if arr[i].Area == -1 {
			fmt.Print(arr[i].VR.Value, " ")
		} else {
			fmt.Print(arr[i].VL.Value, " ")
		}
	}
	fmt.Print("] \n")

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
	}
}
