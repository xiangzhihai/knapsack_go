# knapsack_go
This package is the coding practice for knapsack problem implemented in golang

## Complexity

For both Recursive and DP method of solving this problem, time and space complexity would be both O(M * N) where M is number of items and N is the maxWeight.

Basicly, we would like to know when deciding whether to put a specific item into backpack, what is the best result (value) according to current weight limit and items left. 

For recusive method, I created a map of maps, this has more overheap for space comparing two a 2d arry, but essentially both have the same complexity.