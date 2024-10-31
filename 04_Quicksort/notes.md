# Quicksort


Divide and Conquer -> well known recursive technique
Quicksort is much faster than selection sort

## TLDR
- D&C works by breaking a problem down into smaller and smaller pieces. If you're
using D&C on a list, the base cased is probably an empty array or an array with one element.
- If you're implement qicksort, choose a random element as the pivot. The average
runtime of quicksort is O(n log n)
- The constant in Big O notation can matter sometimes. That's why quicksort is faster than merge sort
- The constant almost never matters for simple search versus binary search,
becauseO(log n) is so much faster than O(n) when your list gets big

## Exercices
in quickort.go
- Remember binary search from chapter 1? It’s a divide-and-conquer
algorithm, too. Can you come up with the base case and recursive
case for binary search?

Base case: an array with one item, if the search item is in the array, you found it! Otherwise it's not in the array
recursive case: split the array in half, and call binarySearch on the other half

## Quicksort
Sorting algorithm frequently used
Really fast
simplest cases: [] and [23]
No need to sort arrays like these
[8,3] just swap the two items

### What about an array of three element ?
33,15,1O
Let's brake this down until the base case
- *33* is the pivot
now find the elements smaller the the pivot and the elements larger than the pivot
15,1O *33* []
This is called partitioning, now you have:
- a sub array of all the numbers less than the pivot
- the pivot
- a sub array of all the numbers greater than the pivot
The two arrays are not sorted, they are just partitioned
If they were sorted you could combine them easily : left array + pivot + right array
[10,15]+[33]+[]
quicksort([15,1O])+[33]+quicksort([])
-> [10,15,33]

This would work with any pivot

### an array with 4 elements ?
[ 33,10,15,7 ]
- pick 33 as pivot
[10,15,7] +[*33*]+[]
[7]+[*10*]+[15]

## Inductive proof
inductive proof  has two steps:
- the base case
- the inductive case
for quicksort:
algo works for base case on arrays of size 0 and 1
on the inductive case, works for an array of size 1, it will work for
an array of size 2. If it works for arrays of size 2, it will work for
arrays of size 3...

## Big O notation revisited
Quicksort worst case O(n<sup>2</sup>)
Average case O( n log n)

### Merge sort vs quicksort
Big O notation different than execution time
O(n) == c*n
where c is "fixed amount of time"
c might be 10ms or 10s

Ex:
10ms * n              1sec * log n
simple search           binary search
In a list of 4 billion elements
simple search = 10ms * 4billion = 463 days
binary search = 1s * 32 = 32 seconds

Quicksort has a smaller constant than merge sort, so in practice quyicksort is faster

### Average case vs worst case
performance of quicksort heacily depends on the pivo you choose.
Suppoose you always choose the first element as the pivot and you call quicksort with an array that is already sort.
Qicksort doesn't check to see whether the input array is already sorted. so it will still try to sort it.

You are not splitting the array into two halves. One sub array is always empty
Ex: for 8 element, call stack height is 8

By splitting the array in halves the call stack height is now 4 for 8 element

## Exercices
How long would each of these operations take in Big O notation?
- Printing the value of each element in an array.
O(n)
- Doubling the value of each element in an array.
O(n)
- Doubling the value of just the first element in an array.
O(1)
- Creating a multiplication table with all the elements in the array. So
if your array is [2, 3, 7, 8, 10], you first multiply every element by 2,
then multiply every element by 3, then by 7, and so on.
O(n <sup>2)
