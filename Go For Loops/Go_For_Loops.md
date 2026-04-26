# Go For Loops

The `for` loop is the **only loop available in Go**. It loops through a block of code a specified number of times. Each execution of a loop is called an **iteration**.

---

## Syntax

```go
for statement1; statement2; statement3 {
    // code to be executed for each iteration
}
```

| Statement    | Purpose |
|--------------|---------|
| `statement1` | Initializes the loop counter value |
| `statement2` | Evaluated each iteration — if `true` loop continues, if `false` loop ends |
| `statement3` | Increases the loop counter value |

> **Note:** These statements don't need to be loop arguments, but they must be present in the code in some form.

---

## Basic For Loop Examples

### Example 1 — Print 0 to 4

```go
package main
import ("fmt")

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

**Output:**
```
0
1
2
3
4
```

**Explained:**
- `i := 0` → start the counter at 0
- `i < 5` → keep looping while i is less than 5
- `i++` → increase i by 1 each iteration

---

### Example 2 — Count to 100 by Tens

```go
package main
import ("fmt")

func main() {
    for i := 0; i <= 100; i += 10 {
        fmt.Println(i)
    }
}
```

**Output:**
```
0
10
20
30
40
50
60
70
80
90
100
```

**Explained:**
- `i := 0` → start at 0
- `i <= 100` → keep looping while i is less than or equal to 100
- `i += 10` → increase i by 10 each iteration

---

## The `continue` Statement

The `continue` statement skips one or more iterations, then continues with the next iteration.

### Example 3 — Skip the Value 3

```go
package main
import ("fmt")

func main() {
    for i := 0; i < 5; i++ {
        if i == 3 {
            continue
        }
        fmt.Println(i)
    }
}
```

**Output:**
```
0
1
2
4
```

> When `i == 3`, the `continue` statement causes the loop to skip printing 3 and jump to the next iteration.

---

## The `break` Statement

The `break` statement terminates the loop immediately.

### Example 4 — Break When i Equals 3

```go
package main
import ("fmt")

func main() {
    for i := 0; i < 5; i++ {
        if i == 3 {
            break
        }
        fmt.Println(i)
    }
}
```

**Output:**
```
0
1
2
```

> When `i == 3`, the `break` statement stops the loop entirely. Nothing after 2 is printed.

> **Note:** `continue` and `break` are usually used with **conditions**.

---

## Nested Loops

A loop can be placed inside another loop. The **inner loop** executes fully for each iteration of the **outer loop**.

### Example 5 — Nested Loop with Arrays

```go
package main
import ("fmt")

func main() {
    adj := [2]string{"big", "tasty"}
    fruits := [3]string{"apple", "orange", "banana"}

    for i := 0; i < len(adj); i++ {
        for j := 0; j < len(fruits); j++ {
            fmt.Println(adj[i], fruits[j])
        }
    }
}
```

**Output:**
```
big apple
big orange
big banana
tasty apple
tasty orange
tasty banana
```

> The outer loop runs 2 times (for each adjective), and the inner loop runs 3 times (for each fruit) — giving 2 × 3 = **6 total iterations**.

---

## The `range` Keyword

The `range` keyword makes it easier to iterate over arrays, slices, or maps. It returns both the **index** and the **value**.

### Syntax

```go
for index, value := range array|slice|map {
    // code to be executed for each iteration
}
```

---

### Example 6 — Range with Index and Value

```go
package main
import ("fmt")

func main() {
    fruits := [3]string{"apple", "orange", "banana"}

    for idx, val := range fruits {
        fmt.Printf("%v\t%v\n", idx, val)
    }
}
```

**Output:**
```
0    apple
1    orange
2    banana
```

---

### Example 7 — Range, Omit the Index

Use `_` (blank identifier) to discard the index when it's not needed.

```go
package main
import ("fmt")

func main() {
    fruits := [3]string{"apple", "orange", "banana"}

    for _, val := range fruits {
        fmt.Printf("%v\n", val)
    }
}
```

**Output:**
```
apple
orange
banana
```

---

### Example 8 — Range, Omit the Value

Use `_` to discard the value when only the index is needed.

```go
package main
import ("fmt")

func main() {
    fruits := [3]string{"apple", "orange", "banana"}

    for idx, _ := range fruits {
        fmt.Printf("%v\n", idx)
    }
}
```

**Output:**
```
0
1
2
```

---

## Summary

| Feature          | Description |
|------------------|-------------|
| `for`            | The only loop in Go |
| `break`          | Exits the loop immediately |
| `continue`       | Skips the current iteration |
| Nested loops     | A loop inside another loop |
| `range`          | Iterates over arrays, slices, or maps |
| `_` (blank id)   | Discards index or value you don't need |

---

*Source: [W3Schools – Go For Loops](https://www.w3schools.com/go/go_loops.php)*