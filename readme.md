# Writing a good multi-threaded program

- To get the maximum performance out of a program, we tend to add threads to it, but it that it?
    - Writing concurrent programs is easy, but things become tricky when we have to ensure correctness and optimality.
    - **Ensuring Correctness**: locking and atomic instructions
    - **Ensuring Optimality**: fairness and efficient logic
- **Counting prime number still 100 million**
    - *Sequential Approach* (3 min 49 sec): For each number, check if it is divisible by any of it’s previous values
        - Actual output: Sequential Approach: Total Prime Numbers = 5761455, Time taken = 4m3.487358s
    - *Add Threads (42 seconds)*: 10 threads, each handling an equal range of  ~10 million
        - We did speed up, but was it fair?
            - Smaller numbers can be checked quickly: 1 to square root (n)
            - More prime numbers in the smaller range
            - Some threads finish early and wait for others to complete

        - We can see that each thread is doing a disproportionate amount of work
    - Threading with fairness (35 seconds)
        - 10 threads, each thread picks up the next unprocessed number and checks if it is prime
        - loop until all numbers are processed, each thread
            - global variable **currentNum**
            - increment **currentNum** atomically: **Correctness**
            - check prime
        - So here each thread will keep checking the next number which is available one by one
        - All threads end at nearly the same time and do nearly the same work: **Maximizing optimality**