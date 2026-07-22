# Chapter 05 - For Loop

## 📖 What is a For Loop?

A `for` loop is used to repeat a block of code multiple times.

Instead of writing the same code again and again, a loop performs the repetition automatically.

---

## 💡 Syntax

```go
for initialization; condition; update {
    // code
}
```

---

## 🧠 What I Learned

- Using the `for` loop
- Initializing a loop variable
- Writing loop conditions
- Updating the loop variable with `i++`
- Repeating tasks efficiently

---

## 💻 Example

See `main.go`.

---

## 📂 Exercises

- Practice exercises are available in `exercises.md`.
- Reference solutions are available in the `solutions/` directory.

---

## ❌ Common Mistakes

Forgetting to update the loop variable.

Incorrect:

```go
for i := 1; i <= 5; {
    fmt.Println(i)
}
```

This creates an infinite loop because `i` never changes.

---

## 💭 Personal Notes

A `for` loop repeats the same task until its condition becomes false.

- A `for` loop has three parts: start, condition, and update.
- Declare the loop variable inside the `for` statement when it is only used in the loop.
  - Example: `for i := 1; i <= 5; i++ {}`

---

## 📌 Summary

A `for` loop allows Go programs to repeat the same task efficiently.
It is one of the most frequently used features in Go programming.