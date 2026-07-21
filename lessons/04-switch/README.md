# Chapter 04 - Switch

## 📖 What is Switch?

A `switch` statement is used to execute different blocks of code based on the value of an expression.

It is often cleaner and easier to read than writing many `if...else if` statements.

---

## 💡 Syntax

```go
switch expression {
case value1:
    // code
case value2:
    // code
default:
    // code
}
```

---

## 🧠 What I Learned

- Using `switch` statements
- Using `case` and `default`
- Comparing values with `switch`
- Replacing multiple `if...else if` statements
- Writing cleaner and more readable code

---

## 💻 Example

See `main.go`.

---

## 📂 Exercises

- Practice exercises are available in `exercises.md`.
- Reference solutions are available in the `solutions/` directory.

---

## ❌ Common Mistakes

Using the wrong data type in `case`.

Incorrect:

```go
switch service {
case nginx:
}
```

Correct:

```go
switch service {
case "nginx":
}
```

---

## 💭 Personal Notes

- Go automatically exits the `switch` after a matching case (no `break` is needed).
- The type of each `case` value must match the type of the `switch` expression.
  - Example: `switch service { case "nginx": }`

---

## 📌 Summary

A `switch` statement makes code cleaner when checking multiple possible values.
It is a great alternative to long `if...else if` chains.
