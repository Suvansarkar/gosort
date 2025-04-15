# gosort
A distributed system project where I implement three different models of sorting algorithm - Odd Even transposition sort, Sasaki's time optimal sort and alternative time optimal algorithm for line networks 

## Odd-Even Transposition Algorithm for distributed sorting on a line networkOdd Even Transposition Sort

On parallel processors, with one value per processor and only local left–right neighbor connections, the processors all concurrently do a compare–exchange operation with their neighbors, alternating between odd–even and even–odd pairings. This algorithm was originally presented, and shown to be efficient on such processors, by Habermann in 1972.

The algorithm extends efficiently to the case of multiple items per processor. In the Baudet–Stevenson odd–even merge-splitting algorithm, each processor sorts its own sublist at each step, using any efficient sort algorithm, and then performs a merge splitting, or transposition–merge, operation with its neighbor, with neighbor pairing alternating between odd–even and even–odd on each step

> Source: wikipedia - https://en.wikipedia.org/wiki/Odd%E2%80%93even_sort

## Comparisions

| Algorithm/Implementation | Time Complexity | Average Runtime (ms) |
|-------------------------|-----------------|---------------------|
| Odd Even transposition  | (n-1) rounds -> O(n^2)| 45                  |
| Sasaki's time optimal   | (n-1) rounds -> O(n^2)| 78                  |
| Alternative time optimal| (n-1) rounds -> O(n^2)| 150                 |

### How to run

```bash
    go run oddeven/oddeven.go <number of processes>
```
> if number of processes are not given a default values of `10`,`20`, `30` and `50` will be taken.

### Sample Output

==========================
N =  10
-----------------------
Random array generated: [14 0 16 3 10 6 2 15 4 19]
Array after  9  passes:  [0 2 3 4 6 10 14 15 16 19]
Messages sent:  40
Messages received:  40
Comparisions:  80
Execution time:  88.925µs
==========================
N =  20
-----------------------
Random array generated: [25 27 10 16 14 15 11 15 22 15 9 22 5 20 7 9 0 10 5 9]
Array after  19  passes:  [0 5 5 7 9 9 9 10 10 11 14 15 15 15 16 20 22 22 25 27]
Messages sent:  220
Messages received:  220
Comparisions:  440
Execution time:  157.77µs
==========================
N =  30
-----------------------
Random array generated: [22 17 7 34 10 7 12 26 5 32 13 25 14 31 31 36 39 22 24 6 34 39 20 6 14 23 2 38 4 3]
Array after  29  passes:  [2 3 4 5 6 6 7 7 10 12 13 14 14 17 20 22 22 23 24 25 26 31 31 32 34 34 36 38 39 39]
Messages sent:  639
Messages received:  640
Comparisions:  1279
Execution time:  492.439µs
==========================
N =  50
-----------------------
Random array generated: [40 26 26 26 28 5 48 20 55 46 8 1 7 21 38 17 59 7 26 39 9 41 6 13 36 2 44 42 4 6 10 50 53 18 6 48 30 40 14 50 25 24 42 43 19 36 29 30 56 14]
Array after  49  passes:  [1 2 4 5 6 6 6 7 7 8 9 10 13 14 14 17 18 19 20 21 24 25 26 26 26 26 28 29 30 30 36 36 38 39 40 40 41 42 42 43 44 46 48 48 50 50 53 55 56 59]
Messages sent:  1836
Messages received:  1839
Comparisions:  3674
Execution time:  1.114159ms
==========================

## Sasaki's time-optimal algorithm for distributed sorting on a line networkSasaki's time optimal sort

"We have achieved a strict lower time bound of n − 1 for distributed sorting on a line network, where n is the number of processes. The lower time bound has traditionally been considered to be n because it is proved based on the number of disjoint comparison-exchange operations in parallel sorting on a linear array. Our result has overthrown the traditional common belief.  2001 Elsevier Science B.V. All rights reserved."

> Source: https://www.sciencedirect.com/science/article/abs/pii/S0020019001003076

### How to run

```bash
    go run sasaki/sasaki.go <number of processes>
```
> if number of processes are not given a default values of `10`,`20`, `30` and `50` will be taken.

### Sample Output

==========================
Random array generated:
[ [ 0 65 ][ 108 108 ][ 15 15 ][ 46 46 ][ 82 82 ][ 89 89 ][ 40 40 ][ 3 3 ][ 103 103 ][ 22 0 ]]
Array after  9  passes:
[ 3 15 22 40 46 65 82 89 103 108 ]
Messages sent:  162
Messages received:  162
Comparisions:  123
Execution time:  111.682µs
==========================
Random array generated:
[ [ 0 74 ][ 59 59 ][ 107 107 ][ 33 33 ][ 73 73 ][ 27 27 ][ 64 64 ][ 109 109 ][ 21 21 ][ 103 103 ][ 104 104 ][ 8 8 ][ 55 55 ][ 95 95 ][ 44 44 ][ 32 32 ][ 11 11 ][ 0 0 ][ 2 2 ][ 73 0 ]]
Array after  19  passes:
[ 0 2 8 11 21 27 32 33 44 55 59 64 73 73 74 95 103 104 107 109 ]
Messages sent:  884
Messages received:  884
Comparisions:  821
Execution time:  315.377µs
==========================
Random array generated:
[ [ 0 87 ][ 44 44 ][ 88 88 ][ 76 76 ][ 21 21 ][ 56 56 ][ 82 82 ][ 79 79 ][ 104 104 ][ 7 7 ][ 120 120 ][ 94 94 ][ 56 56 ][ 65 65 ][ 34 34 ][ 125 125 ][ 114 114 ][ 37 37 ][ 96 96 ][ 106 106 ][ 81 81 ][ 63 63 ][ 86 86 ][ 16 16 ][ 82 82 ][ 21 21 ][ 124 124 ][ 66 66 ][ 64 64 ][ 4 0 ]]
Array after  29  passes:
[ 4 7 16 21 21 34 37 44 56 56 63 64 65 66 76 79 81 82 82 86 87 88 94 96 104 106 114 120 124 125 ]
Messages sent:  2566
Messages received:  2566
Comparisions:  2046
Execution time:  735.769µs
==========================
Random array generated:
[ [ 0 13 ][ 48 48 ][ 64 64 ][ 5 5 ][ 17 17 ][ 66 66 ][ 103 103 ][ 124 124 ][ 44 44 ][ 121 121 ][ 116 116 ][ 63 63 ][ 73 73 ][ 51 51 ][ 110 110 ][ 143 143 ][ 125 125 ][ 83 83 ][ 138 138 ][ 37 37 ][ 83 83 ][ 67 67 ][ 0 0 ][ 143 143 ][ 12 12 ][ 56 56 ][ 50 50 ][ 89 89 ][ 11 11 ][ 32 32 ][ 79 79 ][ 56 56 ][ 89 89 ][ 1 1 ][ 99 99 ][ 102 102 ][ 90 90 ][ 46 46 ][ 42 42 ][ 73 73 ][ 140 140 ][ 2 2 ][ 148 148 ][ 56 56 ][ 113 113 ][ 137 137 ][ 10 10 ][ 148 148 ][ 36 36 ][ 19 0 ]]
Array after  49  passes:
[ 0 1 2 5 10 11 12 13 17 19 32 36 37 42 44 46 48 50 51 56 56 56 63 64 66 67 73 73 79 83 83 89 89 90 99 102 103 110 113 116 121 124 125 137 138 140 143 143 148 148 ]
Messages sent:  7368
Messages received:  7368
Comparisions:  5427
Execution time:  1.290813ms
==========================

## An alternative time-optimal algorithm for distributed sorting on a line network

An alternative approach where each node is labelled from 0 through 2, the middle node receives the data from the neighbouring nodes, compares the value and sends the correct value back to the respective node. The labels are then incremented by 1, where 2 loops back to 0. (n-1) rounds algorithm.

### How to run

```bash
    go run alternative/alternative.go <number of processes>
```

> if number of processes are not given a default values of `10`,`20`, `30` and `50` will be taken.

### Sample Output

==========================
N =  10
-----------------------
Random array generated: [4 4 13 10 0 8 7 7 8 10]
Array after  9  passes:  [0 4 4 7 7 8 8 10 10 13]
Messages sent:  108
Messages received:  108
Comparisions:  84
Execution time:  50.35µs
==========================
N =  20
-----------------------
Random array generated: [8 9 23 9 13 27 0 11 7 0 22 27 27 21 24 15 0 20 12 15]
Array after  19  passes:  [0 0 0 7 8 9 9 11 12 13 15 15 20 21 22 23 24 27 27 27]
Messages sent:  590
Messages received:  590
Comparisions:  452
Execution time:  257.683µs
==========================
N =  30
-----------------------
Random array generated: [23 31 14 30 12 19 20 14 37 1 21 11 38 23 4 2 4 28 5 13 14 2 21 22 1 2 30 5 25 21]
Array after  29  passes:  [1 1 2 2 2 4 4 5 5 11 12 13 14 14 14 19 20 21 21 21 22 23 23 25 28 30 30 31 37 38]
Messages sent:  1712
Messages received:  1712
Comparisions:  1303
Execution time:  660.229µs
==========================
N =  50
-----------------------
Random array generated: [23 2 48 3 20 10 33 33 33 9 48 10 20 54 6 33 19 28 8 0 15 52 44 24 32 20 30 4 41 14 3 58 34 57 4 47 32 41 43 8 33 57 54 53 35 37 57 45 39 29]
Array after  49  passes:  [0 2 3 3 4 4 6 8 8 9 10 10 14 15 19 20 20 20 23 24 28 29 30 32 32 33 33 33 33 33 34 35 37 39 41 41 43 44 45 47 48 48 52 53 54 54 57 57 57 58]
Messages sent:  4908
Messages received:  4904
Comparisions:  3718
Execution time:  1.132313ms
==========================

