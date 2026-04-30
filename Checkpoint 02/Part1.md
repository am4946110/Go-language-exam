# 📋 Quick Reference — Checkpoint 02 Topics

## 1. `Fibonacci(index int) int`

- Return `-1` for negative index
- Base cases: `F(0) = 0`, `F(1) = 1`
- Recursive: `F(n) = F(n-1) + F(n-2)`
- **No `for` loops allowed**

```go
func Fibonacci(index int) int {
    if index < 0 { return -1 }
    if index == 0 { return 0 }
    if index == 1 { return 1 }
    return Fibonacci(index-1) + Fibonacci(index-2)
}
```

---

## 2. `IsPrime(nb int) bool`

- `< 2` → false, `== 2` → true, even → false
- Check odd divisors up to `√nb` (i.e. `i*i <= nb`, `i += 2`)
- **O(√n)** — avoids timeouts

```go
func IsPrime(nb int) bool {
    if nb < 2 { return false }
    if nb == 2 { return true }
    if nb%2 == 0 { return false }
    for i := 3; i*i <= nb; i += 2 {
        if nb%i == 0 { return false }
    }
    return true
}
```

---

## 3. `FindNextPrime(nb int) int`

- `< 2` → return `2`
- If `nb` is even, bump to odd; then loop `+= 2` checking `IsPrime`
- Skip even numbers to halve the work

```go
func FindNextPrime(nb int) int {
    if nb < 2 { return 2 }
    if nb%2 == 0 { nb++ } else { nb += 2 }
    for {
        if IsPrime(nb) { return nb }
        nb += 2
    }
}
```

> ⚠️ Check if the starting `nb` itself is prime first before bumping!

---

## 4. `IterativeFactorial(nb int) int`

- Negative → `0`, `0` → `1`
- Overflow check before multiply: `result > 2147483647/i` → return `0`
- Loop from `i = 2` to `nb`

```go
func IterativeFactorial(nb int) int {
    if nb < 0 { return 0 }
    if nb == 0 { return 1 }
    result := 1
    for i := 2; i <= nb; i++ {
        if result > 2147483647/i { return 0 }
        result *= i
    }
    return result
}
```

---

## 5. `RecursiveFactorial(nb int) int`

- Same logic but recursive — **no `for`**
- Add early guard: `nb > 20` → return `0` (21! overflows int)
- Propagate `0` if recursive call returned `0` and `nb > 1`

```go
func RecursiveFactorial(nb int) int {
    if nb < 0 { return 0 }
    if nb == 0 { return 1 }
    if nb > 20 { return 0 }
    result := RecursiveFactorial(nb - 1)
    if result == 0 && nb > 1 { return 0 }
    if result > 2147483647/nb { return 0 }
    return nb * result
}
```

---

## 6. `IterativePower(nb, power int) int`

- Negative power → `0`, `power == 0` → `1`
- Loop `power` times, multiplying result by `nb`
- No overflow handling needed

```go
func IterativePower(nb int, power int) int {
    if power < 0 { return 0 }
    if power == 0 { return 1 }
    result := 1
    for i := 0; i < power; i++ {
        result *= nb
    }
    return result
}
```

---

## 7. `RecursivePower(nb, power int) int`

- Same as above but recursive — **no `for`**
- Base: `power == 0` → `1`
- Recursive: `nb * RecursivePower(nb, power-1)`

```go
func RecursivePower(nb int, power int) int {
    if power < 0 { return 0 }
    if power == 0 { return 1 }
    return nb * RecursivePower(nb, power-1)
}
```

---

## 8. `Sqrt(nb int) int`

- Negative → `0`, `0` → `0`, `1` → `1`
- Linear search: loop while `i*i <= nb`, return `i` if `i*i == nb`
- Or binary search for O(log n)

```go
func Sqrt(nb int) int {
    if nb < 0 { return 0 }
    for i := 0; i*i <= nb; i++ {
        if i*i == nb { return i }
    }
    return 0
}
```

---

## ⚠️ Most Common Mistakes

| Mistake                                  | Fix                                                |
| ---------------------------------------- | -------------------------------------------------- |
| Forgetting `nb == 2` in `IsPrime`        | Even check would wrongly return false for 2        |
| Missing base case in recursion           | Always handle `0`, `1`, and negatives first        |
| Overflow in factorial                    | Check `result > MAX_INT/i` **before** multiplying  |
| `FindNextPrime` not checking `nb` itself | Check if `nb` is already prime before bumping      |
