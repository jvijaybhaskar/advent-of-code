
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



---
## Part 2 - Thought Process

The CRT screen is a 40x6 grid. The left most positions of a row is indexed by 1 and right most by 40.

Weather to display a particular pixel with `#` or `.` is determined by the overlaying sprite position `###` and the current cycle count. 


```
Cycle   1 -> ######################################## <- Cycle  40
Cycle  41 -> ######################################## <- Cycle  80
Cycle  81 -> ######################################## <- Cycle 120
Cycle 121 -> ######################################## <- Cycle 160
Cycle 161 -> ######################################## <- Cycle 200
Cycle 201 -> ######################################## <- Cycle 240
```

### Logic

- Loop through the value of a register (represents the sprite position) computed in the previous part during each CPU cycle.
    - If there is a an overlap of `sprite position` and `cycle count`
        - `#` is displayed on CRT
    - else 
        - `.` is displayed

    - Offset and print on next line in case the `cycle count` overflow the CRT pixel capacity in a row (40 in this case)


