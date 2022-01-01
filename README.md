# Goterator

 - Goterator is util library that Supports agreegate & transform functions Go, including:
   - for-each
   - find
   - exist
   - reduce
   - filter
   - map
 - API and functions are inspired from Rust and Javascript.

# Requirement

 - Go 1.18

# Installation

This assumes you already have a working Go environment.

Use Go get to pull goterator for using

```
go get github.com/ledongthuc/goterators
```

# Usage

Import the package goterators into your project before using.

```
import "github.com/ledongthuc/goterators"
```

# Functions

## For-each

 - For-each function act the same `for` in Go. Just another option to loop through items in a list.

```go
ForEach(list1, func(item int) {
  // handle each item in the list
})

ForEach(list2, func(item string) {
  // handle each item in the list
})

ForEach(list3, func(item MyStruct) {
  // handle each item in the list
})
```

## Find

 - Find function return first element of the list that meets function condition. In case no element meet the condition function, return the error "Not Found".

```go
matchedInt, err := Find(list, func(item int) bool {
  return item == 1
})

matchedString, err := Find(list, func(item string) bool {
  return item == "searching text"
})

matchedStruct, err := Find(list, func(item MyStruct) bool {
  return item == searchingStruct
})
```

## Exist

 - Exist check an existence of element in the list

```go
matchedInt, err := Exist(list, 1)

matchedString, err := Exist(list, "searching string")

matchedStruct, err := Exist(list, SearchingStruct)
```

## Reduce

 - Similar to Fold Left, Reduce function run the reducer function on each element of array. In order, the reduce function passes in the return value from calculation on the preceding element.  The final result of running the reducer across all elements of the array is the return value of final reducer on last element.
 - Reduce function has 3 parameters:
   - list: source list we want to process.
   - initial value: the previous value that's used in reducer call of first element. At this time, previous = initial value, current = first element of list.
   - reducer function: the function will run on all elements of source list.

```go
// Sum
total := Reduce(list, 0, func(previous int, current int, index int, list []int) int {
	return previous + current
})

// Mapping ints to floats
items := Reduce(testSource, []float64{}, func(previous []float64, current int, index int, list []int) []float64 {
	return append(previous, float64(current))
})
```

## Filter

 - Filter function filters items that meets function condition

```go
filteredItems, err := Filter(list, func(item int) bool {
  return item % 2 == 0
})

filteredItems, err := Filter(list, func(item string) bool {
  return item.Contains("ValidWord")
})

filteredItems, err := Filter(list, func(item MyStruct) bool {
  return item.Valid()
})
```

## Map

 - Map function convert items in the list to output list

```go
mappedItems := Map(testSource, func(item int64) float64 {
  return float64(item)
})

prices := Map(testSource, func(item Order) Price {
  return item.GetPrice()
})
```

# License

MIT

# Contribution

All your contributions to project and make it better, they are welcome. Feel free to reach me https://thuc.space or create an issue to start it.

# Thanks! 🙌

 - Viet Nam We Build group https://webuild.community
