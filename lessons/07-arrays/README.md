# Chapter 07 - Arrays

## What is an Array?

An array is a fixed-size collection of elements of the same type. Each element is accessed using a **0-based index**.

## Syntax

```go
// Fixed-size array
var scores [3]int

// Array with inferred size
colors := [...]string{"Red", "Green", "Blue"}
```

## What I Learned

* Declare fixed-size arrays with `[size]type`
* Get the number of elements with **`len()`**
* Initialize arrays with values using **`[...]`**
* Iterate over arrays using **`for ... range`**
* Work with multi-dimensional (2D) arrays

## Example

See **`main.go`**.

## Exercises

* Exercises: **`exercises.md`**
* Solutions: **`solutions/`**

## Common Mistake

Accessing an index outside the array bounds causes a runtime panic.

```go
var scores [3]int

scores[3] = 100 // ❌ Index out of range
scores[2] = 100 // ✅ Last element
```

For an array of length `3`, valid indexes are:

```text
0, 1, 2
```

## Personal Notes

Go arrays have a **fixed size**, and the size is part of the array's type.

For example:

```go
[3]int
[4]int
```

These are two different types.

Because arrays cannot grow or shrink, **Slices** are used much more often in Go. Slices are dynamic and more flexible.

## Summary

Arrays store multiple values of the same type in a **fixed-size, zero-indexed** collection. They can be easily traversed using `range`.

Understanding arrays is an important step before learning **Slices**.
