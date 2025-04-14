package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

// ProcessMetrics tracks performance metrics for each process
type ProcessMetrics struct {
	ID           int
	FinalValue   int
	MessagesSent int
	MessagesRecv int
	Comparisons  int
	RoundValues  []int // Value after each round
}

// Process represents a single processing entity in the line network
func Process(id int, initialValue int, n int, leftCh, rightCh chan int,
	barrier *sync.WaitGroup, resultCh chan ProcessMetrics) {
	defer barrier.Done()

	metrics := ProcessMetrics{
		ID:           id,
		FinalValue:   initialValue,
		MessagesSent: 0,
		MessagesRecv: 0,
		Comparisons:  0,
		RoundValues:  make([]int, n+1), // +1 to include initial value
	}

	// Record initial value
	metrics.RoundValues[0] = initialValue
	currentValue := initialValue

	// We need n phases to ensure the array is sorted
	for phase := 0; phase < n; phase++ {
		// Even phase: processes with even ids compare-exchange with right neighbor
		// Odd phase: processes with odd ids compare-exchange with right neighbor
		isEvenPhase := phase%2 == 0

		// Wait for all processes to complete the previous phase
		barrier.Done()
		barrier.Wait()
		barrier.Add(1)

		if isEvenPhase { // Even phase (0, 2, 4, ...)
			if id%2 == 0 && id < n-1 {
				// Even process sends to right and receives from right
				rightCh <- currentValue
				metrics.MessagesSent++

				neighborValue := <-rightCh
				metrics.MessagesRecv++

				// Compare and keep the smaller value
				metrics.Comparisons++
				if neighborValue < currentValue {
					currentValue = neighborValue
				}
			} else if id%2 == 1 && id > 0 {
				// Odd process receives from left and sends to left
				neighborValue := <-leftCh
				metrics.MessagesRecv++

				leftCh <- currentValue
				metrics.MessagesSent++

				// Compare and keep the larger value
				metrics.Comparisons++
				if neighborValue > currentValue {
					currentValue = neighborValue
				}
			}
		} else { // Odd phase (1, 3, 5, ...)
			if id%2 == 1 && id < n-1 {
				// Odd process sends to right and receives from right
				rightCh <- currentValue
				metrics.MessagesSent++

				neighborValue := <-rightCh
				metrics.MessagesRecv++

				// Compare and keep the smaller value
				metrics.Comparisons++
				if neighborValue < currentValue {
					currentValue = neighborValue
				}
			} else if id%2 == 0 && id > 0 {
				// Even process (except 0) receives from left and sends to left
				neighborValue := <-leftCh
				metrics.MessagesRecv++

				leftCh <- currentValue
				metrics.MessagesSent++

				// Compare and keep the larger value
				metrics.Comparisons++
				if neighborValue > currentValue {
					currentValue = neighborValue
				}
			}
		}

		// Record value after this round
		metrics.RoundValues[phase+1] = currentValue
	}

	// Update final value in metrics
	metrics.FinalValue = currentValue

	// Send metrics back
	resultCh <- metrics
}

// OddEvenTranspositionSort performs distributed odd-even transposition sort
func OddEvenTranspositionSort(values []int) ([]ProcessMetrics, time.Duration) {
	n := len(values)
	if n <= 1 {
		if n == 1 {
			return []ProcessMetrics{{ID: 0, FinalValue: values[0]}}, 0
		}
		return []ProcessMetrics{}, 0
	}

	// Create channels for communication between processes
	channels := make([]chan int, n-1)
	for i := range channels {
		channels[i] = make(chan int, 1) // Make buffered channels to avoid deadlocks
	}

	// Create barrier for synchronization
	var barrier sync.WaitGroup
	barrier.Add(n) // Add number of processes to barrier

	// Channel for collecting results
	resultCh := make(chan ProcessMetrics, n)

	// Record start time
	startTime := time.Now()

	// Start processes
	for i := 0; i < n; i++ {
		var leftCh, rightCh chan int

		if i > 0 {
			leftCh = channels[i-1]
		}
		if i < n-1 {
			rightCh = channels[i]
		}

		go Process(i, values[i], n, leftCh, rightCh, &barrier, resultCh)
	}

	// Wait for processes to complete initial setup
	barrier.Wait()
	barrier.Add(n) // Reset barrier for phases

	// Wait for all processes to complete
	barrier.Wait()
	duration := time.Since(startTime)

	// Collect results
	metrics := make([]ProcessMetrics, n)
	for i := 0; i < n; i++ {
		metrics[i] = <-resultCh
	}

	// Sort metrics by ID for consistent output
	sortMetricsByID(metrics)

	return metrics, duration
}

// sortMetricsByID sorts metrics by process ID
func sortMetricsByID(metrics []ProcessMetrics) {
	for i := 0; i < len(metrics); i++ {
		for j := i + 1; j < len(metrics); j++ {
			if metrics[i].ID > metrics[j].ID {
				metrics[i], metrics[j] = metrics[j], metrics[i]
			}
		}
	}
}

// GenerateRandomArray creates a random array of integers
func GenerateRandomArray(size int, maxVal int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(maxVal)
	}
	return arr
}

// PrintMetricsTable prints process metrics in a table format
func PrintMetricsTable(metrics []ProcessMetrics, duration time.Duration) {
	// Initialize tabwriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)

	// Print table header
	fmt.Fprintln(w, "ID\tFinal Value\tMessages Sent\tMessages Recv\tComparisons\tExec Time (ms)\tRound Values")
	fmt.Fprintln(w, "--\t-----------\t-------------\t------------\t-----------\t--------------\t------------")

	// Print process metrics
	for _, m := range metrics {
		// Format round values - limiting to first 5 rounds for clarity
		roundVals := ""
		maxRounds := 5
		if len(m.RoundValues) < maxRounds {
			maxRounds = len(m.RoundValues)
		}
		for i := 0; i < maxRounds; i++ {
			if i > 0 {
				roundVals += ", "
			}
			roundVals += fmt.Sprintf("%d", m.RoundValues[i])
		}
		if len(m.RoundValues) > maxRounds {
			roundVals += ", ..."
		}

		fmt.Fprintf(w, "%d\t%d\t%d\t%d\t%d\t%.2f\t[%s]\n",
			m.ID, m.FinalValue, m.MessagesSent, m.MessagesRecv,
			m.Comparisons, float64(duration.Microseconds())/1000.0, roundVals)
	}

	w.Flush()

	// Print summary statistics
	totalMsgSent := 0
	totalMsgRecv := 0
	totalComparisons := 0
	for _, m := range metrics {
		totalMsgSent += m.MessagesSent
		totalMsgRecv += m.MessagesRecv
		totalComparisons += m.Comparisons
	}

	fmt.Printf("\nSummary Statistics (n=%d):\n", len(metrics))
	fmt.Printf("Total Messages Sent: %d\n", totalMsgSent)
	fmt.Printf("Total Messages Received: %d\n", totalMsgRecv)
	fmt.Printf("Total Comparisons: %d\n", totalComparisons)
	fmt.Printf("Total Execution Time: %.3f milliseconds\n", float64(duration.Microseconds())/1000.0)
}

// VerifySorted checks if the array is properly sorted
func VerifySorted(metrics []ProcessMetrics) bool {
	for i := 1; i < len(metrics); i++ {
		if metrics[i-1].FinalValue > metrics[i].FinalValue {
			return false
		}
	}
	return true
}

func claude() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Run for different sizes
	sizes := []int{10, 20, 30, 50}

	for _, size := range sizes {
		fmt.Printf("\n=== Running Odd-Even Transposition Sort with %d processes ===\n\n", size)

		// Generate random array
		values := GenerateRandomArray(size, 1000)

		// Print first few values
		fmt.Printf("Initial values (first 10): ")
		for i := 0; i < min(10, size); i++ {
			fmt.Printf("%d ", values[i])
		}
		if size > 10 {
			fmt.Print("...")
		}
		fmt.Println()

		// Run the algorithm and measure time
		metrics, duration := OddEvenTranspositionSort(values)

		// Print results
		PrintMetricsTable(metrics, duration)

		// Verify the array is sorted
		sorted := VerifySorted(metrics)
		fmt.Printf("\nResult is correctly sorted: %t\n", sorted)
	}
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
