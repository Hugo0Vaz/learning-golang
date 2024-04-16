# Basics

## Packages, variables and functions

### Packages

Every Go program starts running in package `main`.

### Imports

The best way to use imports is convention of the factored import statements.

```go
import (
    "fmt"
    "math"
)
```

### Exported names

Names that begin with capital letters are expected to be exported from the package.

### Functions

A function can contain zero or more parameters.

### Multiple results

A function can return multiple values.

```go
func swap(x, y string) (string, string) {
	return y, x
}
```

### Names return values

Go's return values can be named, therefore the return may be implicit

```go
func split(sum int) (x, y int){
    x = sum * 4 / 9
    y = sum - x
    return
}
```

### Variables

The `var` keyword declares a list of variables of the same type.

```go
var c, python, java bool
```

### Variables with initializers

You can declare varibles and initialize all in one go.

```go
var i, j int = 1, 2
```

### Short variable declarations

Inside a function we can use a short form of variable declaration with implicit type.

```go
func main() {
    var i, j int = 1, 2
    k := 3
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
}
```

### Basic types

Go's basic types are

- bool
- string
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32
- uint64
- uintptr
- byte -> alias of uint8
- rune -> alias for int32: represents a Unicode code point
- float32
- float64
- complex64
- complex128

### Zero values

To declare variables with an explicit initial value are given their zero value.

- 0 for numeric types, `false` for the boolean type and "" (empty string) for strings.

### Type conversions

The expression `T(v)` converts the value `v` to the type `T`

```go
i := 42
f := float64(i)
u := uint(f)
```

### Type inference

When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

### Constants

Constants are declared like variables, but with the `const` keyword.

```go
const Pi = 3.14
```

### Numeric Constants

Numeric constants are high-precision values.

```go
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)
```

## Flow control statements: for , if, else, switch and defer

## More types: structs, slices and maps

# Methods and Interfaces

# Generics

# Concurrency
