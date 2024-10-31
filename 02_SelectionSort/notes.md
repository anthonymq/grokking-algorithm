# Selection Sort
## TLDR
- Memory is like a giant drawer
- when you want to store multiple elements, use an array or a list
- with an array, all your element are stored right next to each other
- arrays allow fast reads
- linked lists allow fast inserts and deletes
- all elements in the array should be the same type (all ints, all doubles...)

## Arrays and Linked List
- Arrays
fixed length
elements in contiguous space in memory
- List
elements anywhere in memory
pointers to previous and next item

|  | Arrays | List |
| ------------- | -------------- | -------------- |
| READING | 0(1) | O(n) |
| INSERTION | 0(n) | O(1) |
| DELETION | 0(n) | O(1) |
O(n) linear time
O(1) constant time
in an array you to move a lot of element when inserting are deleting

### EXERCISES
- Suppose you’re building an app for restaurants to take customer
orders. Your app needs to store a list of orders. Servers keep adding
orders to this list, and chefs take orders off the list and make them.
It’s an order queue: servers add orders to the back of the queue, and
the chef takes the first order off the queue and cooks it.
Would you use an array or a linked list to implement this queue?
(Hint: Linked lists are good for inserts/deletes, and arrays are good
for random access. Which one are you going to be doing here?)
A linked list
- Let’s run a thought experiment. Suppose Facebook keeps a list of
usernames. When someone tries to log in to Facebook, a search is
done for their username. If their name is in the list of usernames,
they can log in. People log in to Facebook pretty often, so there are
a lot of searches through this list of usernames. Suppose Facebook
uses binary search to search the list. Binary search needs random
access—you need to be able to get to the middle of the list of
usernames instantly. Knowing this, would you implement the list
as an array or a linked list?
an array
- 2.4 People sign up for Facebook pretty often, too. Suppose you decided
to use an array to store the list of users. What are the downsides
of an array for inserts? In particular, suppose you’re using binary
search to search for logins. What happens when you add new users
to an array?
if the number of elements is bigger than arrray size, you need to copy every elements in a new array
- 2.5 In reality, Facebook uses neither an array nor a linked list to store
user information. Let’s consider a hybrid data structure: an array
of linked lists. You have an array with 26 slots. Each slot points to a
linked list. For example, the first slot in the array points to a linked
list containing all the usernames starting with a. The second slot
points to a linked list containing all the usernames starting with b,
and so on.
Suppose Adit B signs up for Facebook, and you want to add them
to the list. You go to slot 1 in the array, go to the linked list for slot
1, and add Adit B at the end. Now, suppose you want to search for
Zakhir H. You go to slot 26, which points to a linked list of all the
Z names. Then you search through that list to find Zakhir H.
Compare this hybrid data structure to arrays and linked lists. Is it
slower or faster than each for searching and inserting? You don’t
have to give Big O run times, just whether the new data structure
would be faster or slower.
searching slower than array
inserting faster than array

## Selection sort algo
Ex: Liste de musique avec play count
on veut order by play count
naivement on va devoir iterer plusieurs sur tout les éléments pour trouver le plus écouter
et ajouter à la nouvelle liste
    O(n<sup>2</sup>)

You need a sorting algorithm
- Phone Book
- Travel dates
- Emails (newest to oldest)









