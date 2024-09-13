# skiplist


```golang
$ ./skiplist 
Add 4
Adding 4 to 1 levels from floor
Skip List:
L0 4 --->  nil 

Add 7
Adding 7 to 3 levels from floor
Skip List:
L2 7 --->  nil 
L1 7 --->  nil 
L0 4 ---> 7 --->  nil 

Add 15
Adding 15 to 2 levels from floor
Skip List:
L2 7 --->  nil 
L1 7 ---> 15 --->  nil 
L0 4 ---> 7 ---> 15 --->  nil 

Add 20
Adding 20 to 1 levels from floor
Skip List:
L2 7 --->  nil 
L1 7 ---> 15 --->  nil 
L0 4 ---> 7 ---> 15 ---> 20 --->  nil 

Add 12
Adding 12 to 4 levels from floor
Skip List:
L3 12 --->  nil 
L2 7 ---> 12 --->  nil 
L1 7 ---> 12 ---> 15 --->  nil 
L0 4 ---> 7 ---> 12 ---> 15 ---> 20 --->  nil 

Find 15, found?  true  value:  &{15 [0xc00009e080 <nil>]}  level:  1
Find 10, found?  false  value:  <nil>
Find 12, found?  true  value:  &{12 [0xc00009e060 0xc00009e060 <nil> <nil>]}

Delete 12

Skip List:
L2 7 --->  nil 
L1 7 ---> 15 --->  nil 
L0 4 ---> 7 ---> 15 ---> 20 --->  nil 

Find 12 (again), found?  false  value:  <nil>

```
