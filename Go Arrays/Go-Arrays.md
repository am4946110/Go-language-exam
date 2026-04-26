# Go Arrays

Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

---

## Declaring an Array

In Go, there are two ways to declare an array:

### 1. With the `var` keyword

```go
var array_name = [length]datatype{values}  // length defined

var array_name = [...]datatype{values}      // length inferred
```

### 2. With the `:=` sign

```go
array_name := [length]datatype{values}     // length defined

array_name := [...]datatype{values}        // length inferred
```

> **Note:** The length specifies the number of elements to store in the array. In Go, arrays have a **fixed length**. The length is either defined by a number or **inferred** (the compiler decides based on the number of values provided).

---

## Array Examples

### Defined Length

```go
package main
import ("fmt")

func main() {
  var arr1 = [3]int{1, 2, 3}
  arr2 := [5]int{4, 5, 6, 7, 8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}
```

**Output:**
```
[1 2 3]
[4 5 6 7 8]
```

### Inferred Length

```go
package main
import ("fmt")

func main() {
  var arr1 = [...]int{1, 2, 3}
  arr2 := [...]int{4, 5, 6, 7, 8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}
```

**Output:**
```
[1 2 3]
[4 5 6 7 8]
```

### Array of Strings

```go
package main
import ("fmt")

func main() {
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Print(cars)
}
```

**Output:**
```
[Volvo BMW Ford Mazda]
```

---

## Accessing Elements

You access a specific array element by referring to its **index number**. In Go, array indexes start at **0**.

| Index | Element |
|-------|---------|
| `[0]` | First element |
| `[1]` | Second element |
| `[2]` | Third element |
| ...   | ... |

```go
package main
import ("fmt")

func main() {
  prices := [3]int{10, 20, 30}

  fmt.Println(prices[0])  // 10
  fmt.Println(prices[2])  // 30
}
```

**Output:**
```
10
30
```

---

## Changing Elements

You can change the value of a specific element by referring to its index number.

```go
package main
import ("fmt")

func main() {
  prices := [3]int{10, 20, 30}

  prices[2] = 50
  fmt.Println(prices)
}
```

**Output:**
```
[10 20 50]
```

---

## Array Initialization

If an array or one of its elements has not been initialized, it is assigned the **default value** of its type.

| Type   | Default Value |
|--------|--------------|
| `int`  | `0`          |
| `string` | `""`       |

```go
package main
import ("fmt")

func main() {
  arr1 := [5]int{}          // not initialized
  arr2 := [5]int{1, 2}      // partially initialized
  arr3 := [5]int{1, 2, 3, 4, 5} // fully initialized

  fmt.Println(arr1)
  fmt.Println(arr2)
  fmt.Println(arr3)
}
```

**Output:**
```
[0 0 0 0 0]
[1 2 0 0 0]
[1 2 3 4 5]
```

---

## Initialize Only Specific Elements

You can initialize only specific elements using the `index:value` syntax.

```go
package main
import ("fmt")

func main() {
  arr1 := [5]int{1: 10, 2: 40}

  fmt.Println(arr1)
}
```

**Output:**
```
[0 10 40 0 0]
```

**Explanation:**
- `1:10` → assigns `10` to index `1` (second element)
- `2:40` → assigns `40` to index `2` (third element)
- All other elements default to `0`

---

## Finding the Length of an Array

Use the `len()` function to get the length of an array.

```go
package main
import ("fmt")

func main() {
  arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  arr2 := [...]int{1, 2, 3, 4, 5, 6}

  fmt.Println(len(arr1))  // 4
  fmt.Println(len(arr2))  // 6
}
```

**Output:**
```
4
6
```

---

## Summary

| Feature | Description |
|---------|-------------|
| Fixed size | Array length cannot change after declaration |
| Zero-indexed | First element is at index `0` |
| Default values | Uninitialized elements get the zero value of their type |
| `len()` | Returns the number of elements in the array |
| `[...]` | Lets the compiler infer the array length |

---

*Source: [W3Schools – Go Arrays](https://www.w3schools.com/go/go_arrays.php)*