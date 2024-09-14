# skiplist


```golang
$ ./skiplist 
Add 4
Adding 4 to 1 levels from floor
Skip List:
L00 -> 04 -> nil 

Add 7
Adding 7 to 2 levels from floor
Skip List:
L01 -------> 07 -> nil 
L00 -> 04 -> 07 -> nil 

Add 15
Adding 15 to 1 levels from floor
Skip List:
L01 -------> 07 -> nil 
L00 -> 04 -> 07 -> 15 -> nil 

Add 20
Adding 20 to 2 levels from floor
Skip List:
L01 -------> 07 -------> 20 -> nil 
L00 -> 04 -> 07 -> 15 -> 20 -> nil 

Add 12
Adding 12 to 5 levels from floor
Skip List:
L04 -------------> 12 -> nil 
L03 -------------> 12 -> nil 
L02 -------------> 12 -> nil 
L01 -------> 07 -> 12 -------> 20 -> nil 
L00 -> 04 -> 07 -> 12 -> 15 -> 20 -> nil 

Find 15, found?  true  value:  &{15 [0xc00009e080]}  level:  0
Find 10, found?  false  value:  <nil>
Find 12, found?  true  value:  &{12 [0xc00009e060 0xc00009e080 <nil> <nil> <nil>]}

Delete 12

Skip List:
L01 -------> 07 -------> 20 -> nil 
L00 -> 04 -> 07 -> 15 -> 20 -> nil 

Find 12 (again), found?  false  value:  <nil>
```
