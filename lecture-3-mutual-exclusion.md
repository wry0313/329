# LECTURE 3

## Mutual Exclusion
- property where a block of code that can only be executed by one thread at a time
- 

________________________________

## Partial Orders
- irreflexive: Never true that A -> A
- Antisymmetric: if A -> B then B -> not true
- transitive: if 

## Total Orders

## Locks

## Deadlock-Free
- System as a whole makes progress

## Bakery Algorithm
```C
 public void lock() {
 flag[i] = true;
 label[i] = max(label[0], ..., label[n-1])+1
 while (for some k:
     flag[k] && (label[i], 1) > (label[k], k));
        // someone is ineterested whose label i is earlier in lex order
 }

 public void unlock() {
    flag[i] = false;
}
```
- one deadlock: because only one thread with earliest label
