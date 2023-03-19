
## REQUIREMENT

[--- Day 9: Rope Bridge ---](https://adventofcode.com/2022/day/9)




## Part 1 - Thought process

To map the movement of head and tail on a 2d pattern, a direct/indirect model for a 2x2 grid pattern is required.

Note that the way tail moves to align itself with head needs to incorporated. 
For example, if the head and tail are not on same row or column, the tail position will have to be moved diagonaly until it either touches the head or falls on the smae row/column after which it cane moved horizontally or vertically to align with head. 


### Data strcuture design


type knotPosition struct  {
    x int
    y int
    knotPart string
    previousPosition []Position
}


type Position struct {
    x int
    y int
}



### Logic for moving the tail

They core logic is finding the relative position of tail with rest to head.

This helps determining the direction and limits while moving the tail towards head. 


If the tail and head are aligned vertically or horizontally:
    Move vertically OR 
    Move horizontally
    
To align diagonally:
    Check if the tail will align with head if moved one step diagonally, if not move one step diagonally



----

## Part 2 - Thought process

There are 9 knots to track now on a two dimentional space. 
Each move by the head has a repurcusion on the knots that are following. 

So do we individually compute position of all knots as the head traverses the 2d space. 

One approach is as follows:

For every step the `Head` makes to complete its `large` move
    Loop through an array of `knots` 
        Determine the relative position of a knot with it predessor
        If a move is required
            Move the current knot diagonally OR vertically OR horizontally
                Record the past positions of the current know for mapping

        
