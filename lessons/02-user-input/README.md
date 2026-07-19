# Chapter 02 - User Input

## 📖 What is User Input?

User input allows a program to receive data from the user while it is running.

Instead of hardcoding values, the program waits for the user to enter information such as names, numbers, or other data.

---

## 💡 Syntax

```go
var serverName string

fmt.Print("Enter server name: ")
fmt.Scanln(&serverName)
```

---

## 🧠 What I Learned

- How to read input from the user
- Using `fmt.Scanln()`
- Why `&` is required when reading input
- Storing user input in variables

---

## 💻 Example

See `main.go`.

---

## 📂 Exercises

- Practice exercises are available in `exercises.md`.
- Reference solutions are available in the `solutions/` directory.

---

## ❌ Common Mistakes

Forgetting to use `&` with `fmt.Scanln()`.

---

## 💭 Personal Notes

- `fmt.Scanln()` reads input from the user.
- The default value of an `int` is `0`.
- The default value of a `string` is `""` (an empty string).

---

## 📌 Summary

User input makes programs interactive.

Using `fmt.Scanln()` together with variables allows a Go program to receive and store data entered by the user.