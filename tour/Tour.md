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

### For

Go has only one looping construct, the `for` loop. 

```go
package main

import "fmt"

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
	sum += i
    }
    fmt.Println(sum)
}
```
The for loop contains three sections:
- The first that is executed in the first time the loop runs in the example: `i := 0`.
- The second is the condition expression that is evaluated every interation, in the example: `i < 10`.
- The third is the post statement, executed at the end of every iteration, in the example: `i++`.

### For continued

The first and third statements are optional.

### For is Go's "while"

Since the first and third statements are optional, if we drop the semicolons we can make the for contruct our while loop.

```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```
### Forever

If you omit the second statements what is left is infinite loop.

```go
package main

func main() {
	for {
	}
}
```

### If

Go's if statements don't need surrouding `()`.

```go
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```

### If with a short statement

Like `for`, the `if` statement can start with a short statement to execute before the condition.

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```
### If and else

Variables declared inside an `if` short statement are also available inside any of the `else` blocks.

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

### Exercise: Loops and Functions

- exercise-loops-and-function.go

### Switch

Switch statements in Go are like in any other language, the only difference being that only the matched case ir run therefore not needing the `break` keyword.

### Switch with no condition

Switch without condition is the same as `switch true`.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

### Defer

A `defer` statement defers the execution of a function until the surrouding function returns.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

### Stacking defers

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

## More types: structs, slices and maps

### Pointers

Go has pointers. A pointer holds the memory address of a value.

The `&` operator generates a points to its operand.

```go
i := 42
p = &i
```

The `*` operator denotes the pointer's undelying value.

```go
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

### Structs

A `struct` is a collection of fields.

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

### Struct Fields

Fields are accessed with `.` (dot) notation.

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```

### Pointers to structs

Struct fields can be accessed through a struct pointer.

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

```

### Struct Literals

A struct literal denotes a newly allocated struct value by listing the values of its fields.

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```

### Arrays

The type `[n]T` is an array of `n` values of type `T`

The number of items in the array is a part of it's type.

### Slices

An array is of fixed sizes, but slices are dinamically-sized, flexible view into the contents of the fixed array.

The type `[]T` is an slice with elements of type `T`.

`a[low : high]`

### Slices are like references to arrays

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

### Slice literals

A slice literal is like an array literal without the length.

This is an array literal:

```go
[3]bool{true, true, false}
```
And this creates the same array as above, then builds a slice that references it:

```go
[]bool{true, true, false}
```

### Slice defaults

When slicing, you may omit the high or low bounds to use their defaults instead.

### Slice lenght and capacity

Slices have lenght and capacity.

- Capacity: the capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
- Lenght: the length of a slice is the number of elements it contains.

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

### Nil slices

The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.

### Creating a slice with make

The `make` function allocates a zeroed array and returns  a slice that refers  to that array.

`a := make([]int, 5)`

### Slices of slices

Slices can contain any type, including other slices.

### Appending to slices

It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.

```go
func append(s []T, vs ...T) []T
```
The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

### Range

The `range` form of the `for` loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

### Range continued

You can skip the index or value by assigning to `_`.

```go
for i, _ := range pow
for _, value := range pow
```

### Exercise: Slices

Implement `Pic`. It should return 
