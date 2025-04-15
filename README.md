# gosort
A distributed system project where I implement three different models of sorting algorithm - Odd Even transposition sort, Sasaki's time optimal sort and alternative time optimal algorithm for line networks 

## Comparisions

| Algorithm/Implementation | Time Complexity | Average Runtime (ms) |
|-------------------------|-----------------|---------------------|
| Odd Even transposition  | (n-1) rounds -> O(n^2)| 45                  |
| Sasaki's time optimal   | (n-1) rounds -> O(n^2)| 78                  |
| Alternative time optimal| (n-1) rounds -> O(n^2)| 150                 |

## Odd-Even Transposition Algorithm for distributed sorting on a line networkOdd Even Transposition Sort

On parallel processors, with one value per processor and only local left–right neighbor connections, the processors all concurrently do a compare–exchange operation with their neighbors, alternating between odd–even and even–odd pairings. This algorithm was originally presented, and shown to be efficient on such processors, by Habermann in 1972.

The algorithm extends efficiently to the case of multiple items per processor. In the Baudet–Stevenson odd–even merge-splitting algorithm, each processor sorts its own sublist at each step, using any efficient sort algorithm, and then performs a merge splitting, or transposition–merge, operation with its neighbor, with neighbor pairing alternating between odd–even and even–odd on each step

> Source: wikipedia - https://en.wikipedia.org/wiki/Odd%E2%80%93even_sort

### How to run

```bash
    go run oddeven/oddeven.go <number of processes>
```
> if number of processes are not given a default values of `10`,`20`, `30` and `50` will be taken.

### Sample Output

![Odd Even algorithm output](/screenshots/oddeven.png)

## Sasaki's time-optimal algorithm for distributed sorting on a line networkSasaki's time optimal sort

"We have achieved a strict lower time bound of n − 1 for distributed sorting on a line network, where n is the number of processes. The lower time bound has traditionally been considered to be n because it is proved based on the number of disjoint comparison-exchange operations in parallel sorting on a linear array. Our result has overthrown the traditional common belief.  2001 Elsevier Science B.V. All rights reserved."

> Source: https://www.sciencedirect.com/science/article/abs/pii/S0020019001003076

### How to run

```bash
    go run sasaki/sasaki.go <number of processes>
```
> if number of processes are not given a default values of `10`,`20`, `30` and `50` will be taken.

### Sample Output

![Sasaki's algorithm output](/screenshots/sasaki.png)

## An alternative time-optimal algorithm for distributed sorting on a line network

An alternative approach where each node is labelled from 0 through 2, the middle node receives the data from the neighbouring nodes, compares the value and sends the correct value back to the respective node. The labels are then incremented by 1, where 2 loops back to 0. (n-1) rounds algorithm.

### How to run

```bash
    go run alternative/alternative.go <number of processes>
```

> if number of processes are not given a default values of `10`,`20`, `30` and `50` will be taken.

### Sample Output

![Alternative algorithm output](/screenshots/alternative.png)

