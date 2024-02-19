* lock free: gaurantees one thread can make progress
* wait-free: threads will complete operations regardless of other threads

# Peterson's Algorithm
- provide two thread mutual exclusion
- PREVENT DEADLOCK: because one is waiting for other to release resource
    - because victim and flag ensure processes do not enter critical section simultaneously
    - ok: there is a problem with algos that is missing either the flag or victim because if they try to acquire the lock together then deadlock can happen
        - however for peterson's algorithm the victim field ensures that only one thread can be victim and the other thread's victim == i will become false and get the lock (victim cannot be both A and B)
- starvation free: a thread cannot run forever in the lock() method
```cpp
class Peterson implements Lock {
// thread-local index, 0 or 1
private boolean[] flag = new boolean[2]; 
private int victim;
public void lock() {
    int i = ThreadID.get();
    int j = 1 - i;
    flag[i] = true;
    victim = i;
    while (flag[j] && victim == i) {} // wait
}
public void unlock() {
int i = ThreadID.get(); flag[i] = false;
```

* even tho peterson's algo is deadlock free, when program use multiple p locks, deadlocks can still happen because 

## filter lock 
    * generalizes the peterson's lock into n threads
    * creates n-1 waiting rooms called levels (that threads traverse through to get locks)
        - atleast one thread trying to enter level l succeeds
        - if more than one thread is trying to enter at least one is blocked

* starvation free: gaurantees every thread that call lock eventually genter the critical section
* fairness: A calls lock before B so A enter the CS first
    - use doorway section (always complete in bounded num of steps) and waiting section

# Bakery algorithm
- gaurantees first-come-first-served property
- have a label[i] = max(label[0], ...,label[n-1]) + 1;
    - while ((∃k != i)(flag[k] && (label[k],k) << (label[i],i))) {};
    - a thread is like a receptionist
    - a generated is a customer
    - if two threads got the same customer then the receptionist with the smaller id will get the customer
- may have to worry about label id overflow (as it grow without bound)
