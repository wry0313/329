# Concurrent Computation

- sequential consistency:
    - a concurrent object method call in reality takes times and multiple threads calling can be overlapping. 
    - program order: order in which a single thread calls  

- sequential consistency defined by these two: correctness condition
- method call should appear one at a time in sequential order
- method calls should appear to take effect in program order (calls by diff threads are unrelated by program order)

- real time order
    - When one operation completes before another begins, we say that the first operation precedes the second in the real-time order
    
- linearizability
    - stronger constraint than sequential consistency
    - each method call should appear to take effect instaneously at some moment between invocation and response 
    - real time order of method call must be reserved.
    - linearization point: the instant the method takes effect which is when the other threads see the effect of the execution
    - to be linearization you just have to be able to find a point in the interval where the enq and deq output are agreed. 

sequentially consistent means that if the operation is on diff thread then order doesn't have to be respected
