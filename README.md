# BrickLink Simplex

Ideas how to implement a simplex algorithm to a BrickLink transaction.

## base requirements

1. BrickLink item with a unique id, e. g. a brick 2x4 in color black. 

1. WantList with `i` items I want to buy in `q` quantity each.

itemId | qty
--- | ---
item1 | 1
item2 | 2
item3 | 3

1. Seller, who has `i` items in quantity `q` to sell. Each item is offered for a specific price `p`. A seller could have a minimum order value, meaning an order at seller `s` with `i` items in quantity `q` for price `p` each must be higher or equal to the minimum order value `v`. Each seller has a unique sellerId. 

itemId | qty | price
--- | --- | ---
item1 | 3 | 0.30
item2 | 0 | 0
item3 | 2 | 0.20

1. AvailabilityList