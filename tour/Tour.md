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
