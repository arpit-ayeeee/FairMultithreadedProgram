## Output:
- Thread 0 [3, 10000003] completed in 15.3733946s
- Thread 1 [10000003, 20000003] completed in 28.9208539s

- Thread 2 [20000003, 30000003] completed in 34.647904s
- Thread 3 [30000003, 40000003] completed in 39.8680671s

- Thread 4 [40000003, 50000003] completed in 43.4786427s

- Thread 5 [50000003, 60000003] completed in 46.131587s
- Thread 6 [60000003, 70000003] completed in 48.1935657s

- Thread 7 [70000003, 80000003] completed in 50.346654s
- Thread 8 [80000003, 90000003] completed in 52.0361813s

- Thread 9 [90000003, 100000000] completed in 54.5982922s
- Checking till 100000000 found 5761457 prime numbers took 54.5982922s

## This is mentioned as an unfair approach, because the inital batches allocated to the thread pool completed in much less time than the latter batched, we need ot find a way to use all the thread in their max potential