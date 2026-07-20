# Chapter 03 - If Else

## 📖 What is If Else?

`if` statements allow a program to make decisions based on conditions.

Instead of always executing the same code, a program can choose different actions depending on the input or data.

---

## 💡 Syntax

```go
if condition {
    // code
} else if anotherCondition {
    // code
} else {
    // code
}
```

---

## 🧠 What I Learned

- Using `if` to make decisions
- Using `else if` for multiple conditions
- Using `else` for the remaining cases
- Comparing values with comparison operators
- Writing boolean expressions with `&&`

---

## 💻 Example

See `main.go`.

---

## 📂 Exercises

- Practice exercises are available in `exercises.md`.
- Reference solutions are available in the `solutions/` directory.

---

## ❌ Common Mistakes

Using `=` instead of `==` when comparing values.

Incorrect:

```go
if role = "admin" {
}
```

Correct:

```go
if role == "admin" {
}
```

---

## 💭 Personal Notes

The program checks each condition from top to bottom.
When the first condition is true, it executes that block and skips the rest.

- `==` compares values.
- `&&` means **and**.
- `else` handles all remaining cases.
- Avoid repeating conditions that have already been checked.
  - Example: `port <= 65535` instead of `port >= 1024 && port <= 65535`.

---

## 📌 Summary

`if`, `else if`, and `else` allow Go programs to make decisions based on different conditions.
