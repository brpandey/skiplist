# skiplist

<p float="center">
  <img src='images/sign_express_local_platform.jpg' width='800' height='225'/>
</p>

<p float="center">
  <img src='images/flushing.png' width='800' height='425'/>
</p>

<p float="center">
    <img src='images/subway.png' width='450' height='325'/>
    <img src='images/brooklyn.jpg' width='250' height='325'/>
</p>

[Express Or Local](http://warofyesterday.blogspot.com/2010/03/subway-map-ii.html)


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
