package main

func main() {
}

type Item struct {
    Id string
    statusNew bool

    Qty int
    Price float64
}

type WantList map[string]int

type Seller struct {
    Id string
    Items []Item
    MinOrder float64
}