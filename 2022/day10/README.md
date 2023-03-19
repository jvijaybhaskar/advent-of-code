
## REQUIREMENT

[Day 10: Cathode-Ray Tube](https://adventofcode.com/2022/day/10)

---
## Part 1 - Thought Process

An addition instruction `addx V` takes two cycles to compelete.
The `noop` instruction takes one cycle and does nothing. 

```
   _     _     _     _ 
_ |1| _ |2| _ |3| _ |4|
<  addx C > < n > < n >

```

The goal is to review the value of a register `x` throughout execution. It is quanitifed by `signal strength` 

`signal strength` is the product of cycle number multipleied by value of register `x`

### Logic

- Initialize cycle `cycle`
- Initialize a slice to model cycle count and value of register `x` at any point of time.

- Preare a list of instructions (`noop` and `addx`)
- Loop through the list
    - If `noop`
        - increment the `cycle` by 1
        - copy over the value of `x` to the cycle index
    - else 
        - increment the `cycle` by 1
        - copy over the value of `x` to the cycle index
        - increment the `cycle` by 1
        - Update the value of register and store it in the slice


- Retrieve the value of `x` during 20th cycle and every 40th cycle and sum them up



