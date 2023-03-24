## REQUIREMENTS

[Day 11: Monkey in the Middle] https://adventofcode.com/2022/day/11

---
## Part 1 - Thought Process

### Logic

Create a data strcucture to represent a monkey and its properties. 
It should hold  the following info:
- Monkey id 
- Starting Items 
- TestCondition
- Action array
- InspectedItemCount

Create a map of monkeys and loop through it `x` number of times.

Identify the top two moneys who have dealt with maximum items.

---

## Part 2 - Thought Process

The key challenge in this part is to find another way to keep the worry levels managable as the number of rounds increase.

In this part instead of dividing, modulo arithmatic is used to manage the worry levels. The bigger challenge is finding a suitable number to compute the modulo.
This is performed by multiplying the divisibility test numbers provided in the data. 

---
@TODO
Write a function to sort the list of monkey by the most inspected items
