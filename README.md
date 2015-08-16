# monty
An experiment investigating the performance gains obtained by using golang's concurrency to solve for pi.

Based off [this](http://www.soroushjp.com/2015/02/07/go-concurrency-is-not-parallelism-real-world-lessons-with-monte-carlo-simulations/) interesting article by Soroush Pour

Results on an 8 core (2.4Ghz) box:
```
go install; time monty
```
Single Routine
```
3.141585

real    2m54.712s
user    2m54.848s
sys     0m0.084s
```
----------------------------------------
```
Multiple Routines

3.141588

real    0m12.695s
user    1m41.116s
sys     0m0.008s
```
