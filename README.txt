# gosort
A distributed system project where I implement three different models of sorting algorithm - Odd Even transposition sort, Sasaki's time optimal sort and alternative time optimal algorithm for line networks 

## Odd-Even Transposition Algorithm for distributed sorting on a line networkOdd Even Transposition Sort

On parallel processors, with one value per processor and only local left–right neighbor connections, the processors all concurrently do a compare–exchange operation with their neighbors, alternating between odd–even and even–odd pairings. This algorithm was originally presented, and shown to be efficient on such processors, by Habermann in 1972.

The algorithm extends efficiently to the case of multiple items per processor. In the Baudet–Stevenson odd–even merge-splitting algorithm, each processor sorts its own sublist at each step, using any efficient sort algorithm, and then performs a merge splitting, or transposition–merge, operation with its neighbor, with neighbor pairing alternating between odd–even and even–odd on each step

> Source: wikipedia - https://en.wikipedia.org/wiki/Odd%E2%80%93even_sort

### How to run

```bash
    go run oddeven/oddeven.go <number of processes>
```
> if number of processes are not given a default values of `10`,`20` and `30` will be taken.

### Sample Output

==========================
N =  10
-----------------------
Random array generated: [6 19 19 7 7 6 2 7 13 19]
Array after  9  passes:  [2 6 6 7 7 7 13 19 19 19]
Messages sent:  40
Messages received:  40
Comparisions:  80
Execution time:  64.786µs
==========================
N =  20
-----------------------
Random array generated: [4 11 28 22 8 3 3 21 17 11 27 18 27 26 5 18 12 22 14 4]
Array after  19  passes:  [3 3 4 4 5 8 11 11 12 14 17 18 18 21 22 22 26 27 27 28]
Messages sent:  220
Messages received:  219
Comparisions:  439
Execution time:  168.105µs
==========================
N =  30
-----------------------
Random array generated: [2 25 21 24 14 10 21 5 29 39 26 35 5 29 18 34 8 30 16 17 27 4 1 0 39 3 36 17 7 18]
Array after  29  passes:  [0 1 2 3 4 5 5 7 8 10 14 16 17 17 18 18 21 21 24 25 26 27 29 29 30 34 35 36 39 39]
Messages sent:  640
Messages received:  639
Comparisions:  1279
Execution time:  559.394µs
==========================

## Sasaki's time-optimal algorithm for distributed sorting on a line networkSasaki's time optimal sort

"We have achieved a strict lower time bound of n − 1 for distributed sorting on a line network, where n is the number of processes. The lower time bound has traditionally been considered to be n because it is proved based on the number of disjoint comparison-exchange operations in parallel sorting on a linear array. Our result has overthrown the traditional common belief.  2001 Elsevier Science B.V. All rights reserved."

> Source: https://www.sciencedirect.com/science/article/abs/pii/S0020019001003076

### How to run

```bash
    go run sasaki/sasaki.go <number of processes>
```
> if number of processes are not given a default values of `10`,`20` and `30` will be taken.

### Sample Output

==========================
Random array generated:
[ [ 0 62 ][ 87 87 ][ 25 25 ][ 24 24 ][ 37 37 ][ 66 66 ][ 14 14 ][ 26 26 ][ 93 93 ][ 63 0 ]]
Array after  9  passes:
[ 14 24 25 26 37 62 63 66 87 93 ]
Messages sent:  162
Messages received:  162
Comparisions:  102
Execution time:  112.743µs
==========================
Random array generated:
[ [ 0 72 ][ 9 9 ][ 61 61 ][ 44 44 ][ 50 50 ][ 45 45 ][ 5 5 ][ 109 109 ][ 75 75 ][ 59 59 ][ 74 74 ][ 74 74 ][ 106 106 ][ 89 89 ][ 45 45 ][ 85 85 ][ 5 5 ][ 4 4 ][ 11 11 ][ 50 0 ]]
Array after  19  passes:
[ 4 5 5 9 11 44 45 45 50 50 59 61 72 74 74 75 85 89 106 109 ]
Messages sent:  884
Messages received:  884
Comparisions:  615
Execution time:  474.447µs
==========================
Random array generated:
[ [ 0 95 ][ 76 76 ][ 94 94 ][ 36 36 ][ 54 54 ][ 48 48 ][ 37 37 ][ 123 123 ][ 6 6 ][ 42 42 ][ 91 91 ][ 86 86 ][ 32 32 ][ 101 101 ][ 119 119 ][ 41 41 ][ 95 95 ][ 10 10 ][ 57 57 ][ 108 108 ][ 12 12 ][ 108 108 ][ 19 19 ][ 109 109 ][ 52 52 ][ 46 46 ][ 104 104 ][ 55 55 ][ 2 2 ][ 81 0 ]]
Array after  29  passes:
[ 2 6 10 12 19 32 36 37 41 42 46 48 52 54 55 57 76 81 86 91 94 95 95 101 104 108 108 109 119 123 ]
Messages sent:  2566
Messages received:  2566
Comparisions:  1850
Execution time:  712.647µs
==========================

## An alternative time-optimal algorithm for distributed sorting on a line network

An alternative approach where each node is labelled from 0 through 2, the middle node receives the data from the neighbouring nodes, compares the value and sends the correct value back to the respective node. The labels are then incremented by 1, where 2 loops back to 0. (n-1) rounds algorithm.

### How to run

```bash
    go run alternative/alternative.go <number of processes>
```

> if number of processes are not given a default values of `10`,`20` and `30` will be taken.

### Sample Output

==========================
N =  10
-----------------------
Random array generated: [18 16 2 6 8 7 14 5 0 4]
Array after  9  passes:  [0 2 4 5 6 7 8 14 16 18]
Messages sent:  108
Messages received:  108
Comparisions:  84
Execution time:  115.875µs
==========================
N =  20
-----------------------
Random array generated: [12 5 6 6 23 11 28 2 8 3 0 12 28 6 3 5 18 8 17 16]
Array after  19  passes:  [0 2 3 3 5 5 6 6 6 8 8 11 12 12 16 17 18 23 28 28]
Messages sent:  590
Messages received:  590
Comparisions:  452
Execution time:  248.545µs
==========================
N =  30
-----------------------
Random array generated: [15 8 20 23 13 30 23 8 37 12 24 7 13 31 38 31 34 36 24 13 3 28 0 29 35 11 7 10 27 4]
Array after  29  passes:  [0 3 4 7 7 8 8 10 11 12 13 13 13 15 20 23 23 24 24 27 28 29 30 31 31 34 35 36 37 38]
Messages sent:  1712
Messages received:  1712
Comparisions:  1300
Execution time:  459.461µs
==========================

