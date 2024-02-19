# LECTURE 2

## Deadlock

## Proof of Mutual Exclusion (proof of contradiction)
    -  assuming both thread writing to memory X
    -  by reasoning backward both threads need to request to write to memory
    -  but by the flag protocol only one thread can obtain the memory address

 ## Proof of No Deadlock

1. mutual exclusion
2. producer/consumer
3. no starvation

### Readers/Writers
    - easy with mutual exclusion (but requires waiting)
    - RW w/out mutual exclusion
        - thread read one char at a time
        - thread write one char at a time

### Amdahl's Law
w 
    - speedup = 1-thread execution time / n-thread execution time
    - 1 / ((1 - p) + p/n)
        - p is percent of parallelizable portion of the task
        - n is num of thread
    - calculate the upper bound of speedup


### Thread
    - is just a sequence of events
