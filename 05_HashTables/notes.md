# Hash Tables

| # of items in the book | O(n) | O(log n) |
| --------------- | --------------- | --------------- |
| 100 | 10s | 1s   <- log2 100 = 7lines |
| 1000 | 1.66min | 1s  <- log2 1OOO = 10lines |
| 10000 | 16.6min | 2s  <- log2 1OOOO = 14 lines |

Hash Tables gives the response in O(1) time, no matter how big the book is
Faster than binary search
## TLDR
hashes are good for:
- Modeling relationships from one thing to another thing
- Filtering out duplicates
- Caching/memorizing data instead of making your server do  work

## Hash Functions
A hash function is a function where you put in a string and you get back a number

NAMASTE -> f(x)=> 7
HOLA -> f(x)=> 4
HELLO -> f(x)=> 2

A hash function "Maps strings to numbers"
- Needs to be consistent. Should always return the same response
- Should map different to different numbers

- the hash function consistently maps a name to the same index.
- the hash function maps different sstrings to differrent indexes
- the hash function knows how big your array is and only returns valid indices. If you array is 5 items, the hash function doesn't return 100.

A hash table is a data structure with extra logic behind it.
Probably the most useful complex data tructure you'll learn

## Exercices
Which of these hash functions are consistent?
- f(x) = 1 Returns “1” for all input
yes
- f(x) = rand() Returns a random number every time
no
- f(x) = next_empty_slot()
no
- f(x) = len(x) Uses the length of the string as the index
yes


## Usecases
hashtables for lookups
- Phonebook
- DNS domain name -> IP
- Preventing double entries
    - Voting booth
- Use as cache

## Collisions
Most languages have hash tables. You don't need to know how to write you own
map of linked lists to store every items
- Your hash function is really important. Should map keys evenly all over the hash
- If those linked lists get long, it slows down you hash table a lot.
A good hash function will give you very few collisions

## Perfomance
| x | Average case | Worst case |
| --------------- | --------------- | --------------- |
| search | O(1) | O(n) |
| Insert | O(1) | O(n) |
| delete | O(1) | O(n) |

Search is not instant. But it doesn't matter the size of the hash table
1 element or 1 billion would take the same time for search
