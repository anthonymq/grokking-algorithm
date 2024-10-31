# Introduction to algorithms
## TLDR
- BInary Search is a lot faster than simple search
- O(log n ) is faster than O(n), but it gets a lot faster once the list of items you're searching through grows
- Algorithm speed isn't measeured in seconds
- Algorithm timmes are measured in terms of growth of an algorithm
- Algorihm times are written in Big O notation

## Binary search
### What's the maximum number of steps to search in a list with 128 item ?
$$
\log_{2} 128 = 7
2^7 = 128
$$

### What if you double the list ?
$$
log_2 256 = 8
2^8 = 256
$$

## Running time
Simple search linear time : O(n)
binary search logarithmic time : O(log n)
4 000 000 000 => 32 steps with binary search
Generally you want to chose the most efficient algorithm in terms of time and space.

## Big O notation
Tells you how fast is an algorithm

## Algo running time grow at different rates
each check take 1ms
with simple search on 1000000000 = 1 000 000 000ms = 11days
vs binary search log2(1000000000)=30ms
binary search is 33million times faster on 1 billion element
Big O let's you compare the number of operations


> **Big 0 establishes a worst-case runtime**

## common Big O runtimes
- O(log n) => log time. Ex: binary search
- O(n) => linear time. Ex: simple search
- O(n * log n) => fast sorting algorithm like quicksort
- O(n<sup>2</sup>) => slow sorting algorithm like selection sort
- O(!n) => really slow sorting like traveling salesperson

## Exercices
Give the run time for each of these scenarios in terms of Big O notation
- You have a name and you want to find the person's phone number in the phone book
O(log n)
- You have a phone number and you want to find the person's name in the phone book
O(n)
- You want to read the numbers of every person in the phone book
O(n)
- You want to read the numbers of just the As
O(n)
