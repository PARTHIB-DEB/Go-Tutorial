# Mastering `break` and `continue` in Go

In Go, `break` is used to **exit** a loop immediately, while `continue` is used to **skip** the current iteration and proceed to the next one. Understanding their behavior is essential for controlling loop execution.

## Questions

### 1️⃣ Breaking Out of a Loop  
Write a program that prints numbers from 1 to 10 but stops printing when it reaches 6 using `break`.  

### 2️⃣ Skipping Even Numbers with `continue`  
Modify the given program to print only the odd numbers from a list while skipping the even numbers using `continue`.  

### 3️⃣ Controlling Nested Loops with `break`  
Create a program with a nested loop (one `for` loop inside another).  
Use `break` inside the inner loop and analyze how it affects the outer loop.  

### 4️⃣ Using `continue` to Filter Output  
Write a loop that iterates from 1 to 20. Use `continue` to **skip numbers divisible by 3** and print the rest.  

### 5️⃣ Breaking an Infinite Loop  
Write an infinite `for` loop and use `break` to exit when a user enters a specific input (e.g., `"exit"`).  

### 6️⃣ Understanding `continue` with a Condition  
Modify the given program to **skip printing numbers greater than 15** using `continue`.  

### 7️⃣ Breaking a Loop at a Specific Index  
Modify the given program to exit the loop using `break` when the `i` index reaches **3**.  

### 8️⃣ Using `continue` to Retry an Operation  
Write a program that simulates a retry mechanism:  
- If a generated random number is less than 5, use `continue` to retry the operation.  
- Otherwise, print `"Success"` and exit the loop.  

### 9️⃣ Unexpected Behavior with `continue` and Appending  
Analyze the given program:  
- Why does the loop run longer than expected?  
- How does appending affect `mylist`?  
- What role does `continue` play in extending the loop?  

### 🔟 Combining `break` and `continue`  
Modify the program so that:  
- `continue` is used to skip numbers divisible by 4.  
- `break` is used to exit when the sum of printed numbers exceeds 50.  