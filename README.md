# Goterators
[![Built with WeBuild](https://raw.githubusercontent.com/webuild-community/badge/master/svg/WeBuild.svg)](https://webuild.community) [![Go Reference](https://pkg.go.dev/badge/github.com/ledongthuc/goterators.svg)](https://pkg.go.dev/github.com/ledongthuc/goterators)

![goterators-Thumbnail](https://user-images.githubusercontent.com/1828895/147876484-5bc7cfd0-5f14-4889-a3f0-64cb307b7765.png)

 - Goterators is util library that support aggregate & transforms functions list in Go, including:
   - [for-each](#for-each)
   - [find](#find)
   - [exist](#exist)
   - [reduce](#reduce)
   - [reduce right](#reduce-right)
   - [filter](#filter)
   - [map](#map)
   - [every](#every)
   - [some](#some)
   - [flat](#flat)

 - The API and functions are inspired by Rust and Javascript.

# Requirement

 - Go 1.18

# Installation

This assumes you already have a working Go environment.

Use Go get to pull goterators for using

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

![goterators-ForEach](https://user-images.githubusercontent.com/1828895/147876359-432c3122-25f3-492e-844d-6172eafe92a6.png)

 - For-each function does the same `for` in Go. Just another option to loop through items in a list.

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

![goterators-Find](https://user-images.githubusercontent.com/1828895/147876363-245705c5-2aa8-4135-8d29-5cbaca173529.png)

 - Find function returns the first element and its index in the list that meets the functional condition. If no element meet the condition function, return the error "Not Found".

```go
matchedInt, index, err := Find(list, func(item int) bool {
  return item == 1
})

matchedString, index, err := Find(list, func(item string) bool {
  return item == "searching text"
})

matchedStruct, index, err := Find(list, func(item MyStruct) bool {
  return item == searchingStruct
})
```

## Exist

![goterators-Exist](https://user-images.githubusercontent.com/1828895/147876367-c0c7fd50-1888-4152-a7c8-5960ca26b6d9.png)

 - Exist function checks the existence of an element in the list.

```go
matchedInt, err := Exist(list, 1)

matchedString, err := Exist(list, "searching string")

matchedStruct, err := Exist(list, SearchingStruct)
```

## Reduce

![goterators-Reduce](https://user-images.githubusercontent.com/1828895/147876373-4cb1f784-b9d4-4b95-a4f9-30709ba3690d.png)

 - Similar to Fold Left, Reduce function runs the reducer function on each element of the array. In order, the reduce function passes in the return value from the calculation on the preceding element. The final result of running the reducer across all elements of the array is the return value of the final reducer on the last element.
 - Reduce function has 3 parameters:
   - list: source list we want to process.
   - initial value: the previous value that's used in the reducer call of the first element. At this time, previous = initial value, current = first element of the list.
   - reducer function: the function will run on all elements of the source list.

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

## Reduce right

![goterators-Reduce right](https://user-images.githubusercontent.com/1828895/147876376-1f168d48-3ba5-4d44-aa0a-4e1c36505be5.png)

 - Similar to Fold Right, Reduce right function run the reducer function on each element of the array, from last to the first element. In order, the reduce function passes in the return value from the calculation on the preceding element. The final result of running the reducer across all elements of the array is the return value of the final reducer on the first element.
 - Reduce function has 3 parameters:
   - list: source list we want to process.
   - initial value: the previous value that's used in the reducer call of the last element. At this time, previous = initial value, current = last element of list.
   - reducer function: the function will run on all elements of the source list.

```go
// Reverse
reversedList := Reduce(list, []string{}, func(previous []string, current string, index int, list []string) []string {
  return append(list, current)
})
```

## Filter

![goterators-Filter](https://user-images.githubusercontent.com/1828895/147876377-6df6ea14-6fb7-478c-9671-e6ca7c6e2a97.png)

 - Filter function return items that pass the filter function.

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

![goterators-Map](https://user-images.githubusercontent.com/1828895/147876383-5d701c6e-fb65-442f-b5ed-e97d30c23115.png)

 - Map function converts items in the list to the output list.

```go
mappedItems := Map(testSource, func(item int64) float64 {
  return float64(item)
})

prices := Map(testSource, func(item Order) Price {
  return item.GetPrice()
})
```

## Every

![goterators-Every](https://user-images.githubusercontent.com/1828895/147876387-520ee3b5-2846-4052-ad35-dd57d8741bd1.png)

 - Every function checks all elements in the list with condition function. If it's yes return true; otherwise, return false.

```go
valid := Every(list, func(item int) bool { item % 2 == 0 }) 

valid := Every(list, func(item string) bool { len(item) < 20 }) 
```

## Some

![goterators-Some](https://user-images.githubusercontent.com/1828895/147876396-bbf186e9-443a-4d66-85fe-644f72746b43.png)

 - Some function check at least one element in the list meet the condition; return true, or return false if all elements don't meet the condition.

```go
valid := Some(list, func(item int) bool { item % 2 == 0 }) 

valid := Some(list, func(item string) bool { len(item) < 20 }) 
```

## Group

![goterators-Group](https://user-images.githubusercontent.com/1828895/147878206-bef39880-96db-4269-b54e-2dcbb06f6bac.png)

 - Group groups elements into the nested level. Use a build-group function to define group type.

```
groups := Group(list, func(item Product) groupKey {
   return item.ComposeGroupKey()
}) // Group contains [ [ Product1, Product2, Product3 ], [ Product4, Product5 ] ]
```

## Flat

![goterators-Flat](https://user-images.githubusercontent.com/1828895/147876403-25e84044-d761-45b7-b126-6ad8a7c5a4d1.png)

 - Flat returns a new array with all sub-array elements concatenated with 1 level depth.

```go
output := Flat([][]int{{1,2}, {3}}) // output = {1,2,3}
```

# License

MIT

# Contribution

All your contributions to project and make it better, they are welcome. Feel free to reach me https://thuc.space or create an issue to start it.

# Thanks! 🙌

 - Viet Nam We Build group https://webuild.community

[![Stargazers repo roster for @ledongthuc/goterators](https://reporoster.com/stars/ledongthuc/goterators)](https://github.com/ledongthuc/goterators/stargazers)
