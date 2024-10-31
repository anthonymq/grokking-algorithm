# Recursion
"Loops may achieve a performance gain for your programmer. Choose whick is more omprtant in your situation!"
Leigh Caldwell
Recursive functions calls itself. -> may end up in an infinite loop.

## TLDR
- Recursion is when a function alls itself
- Every recursive function has two case: the base case and the recursive case
- a stack has two operations: push and pop
- All function calls go onto the call stack
- the call stacck can get very large, whick takes up a lot of memory

## Base case and recursive case
You have to tell a recursive function when to stop.
Two parts: the base and the recursive case
Stops the infinite loop

## The stack
the call stack
Two actions :
- push (insert at the top)
- pop (remove and read at the top)

When a you call a function for another function, the calling function is paused in a partially completed state.
All the values of the variables for that function ar still stored in memory.

### Exercices
Suppose you accidentally write a recursive function that runs
forever. As you saw, your computer allocates memory on the
stack for each function call. What happens to the stack when your
recursive function runs forever?
The call strack grows forever. Each program has a limited space on the call stack.
When out of space -> exit with stack overflow error
