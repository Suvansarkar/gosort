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

if number of processes are not given a default value of 10 will be taken.

## Sasaki's time-optimal algorithm for distributed sorting on a line networkSasaki's time optimal sort

## An alternative time-optimal algorithm for distributed sorting on a line network
