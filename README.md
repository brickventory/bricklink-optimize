# BrickLink Simplex

Ideas how to implement a simplex algorithm to a BrickLink transaction.

## The Problem

How to find the best combination of shops at [BrickLink](bricklink.com) with the lowest price for a list of wanted bricks. 

## Background



## base requirements

1. BrickLink item with a unique id, e. g. a brick 2x4 in color black. 

1. WantList with a number of `I` items I want to buy in `q` quantity each.

itemId | qty
--- | ---
item1 | 1
item2 | 2
item3 | 3

1. A number of `S` Seller, who offer items from my WantList. 

1. Each Seller `s` has `i` items in quantity `q` to sell. Each item is offered for a specific price `p`. A seller could have a minimum order value `v`, meaning an order at seller `s` with `i` items in quantity `q` for price `p` must be higher or equal in value to the minimum order value `v`.

itemId | qty | price
--- | --- | ---
item1 | 3 | 0.30
item2 | 0 | 0
item3 | 2 | 0.20

1. AvailabilityList