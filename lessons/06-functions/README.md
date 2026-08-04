# Chapter 06 - Functions

## What is a Function?
A function is a reusable block of code that performs a specific task.
Instead of repeating the same logic, we wrap it inside a function and call it whenever needed.

## Syntax

    func functionName(parameter type) returnType {
        // code
        return value
    }

## What I Learned
- Defining functions with parameters and return types
- Returning multiple values from a function
- The `error` type and how Go handles errors
- Using `nil` to represent "no error"
- Shorthand syntax for same-type parameters (a, b int)

## Example
See `main.go`.

## Exercises
- Practice exercises are available in `exercises.md`.
- Reference solutions are available in the `solutions/` directory.

## Common Mistakes
Forgetting to check the error before using the result.

Incorrect:

    result, err := divide(10, 0)
    fmt.Println(result)

This ignores `err` and may use an invalid result.

Correct:

    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println(result)
    }

## Personal Notes
In Go, errors are returned as values instead of being thrown like exceptions.
A function that can fail usually returns `(result, error)`, and the caller
must check if `err != nil` before trusting the result.

## Summary
Functions let us organize code into reusable pieces. In Go, functions can
return multiple values, which is commonly used for returning both a result
and an error.