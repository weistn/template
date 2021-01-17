# template
Extended Go template package

This package is a clone of the template package of the Go standard library. 
Thus, this code uses the same license as the original Go code from which it is derived.

# Lambda Functions and Currying

This version adds support for functional programming.
It is possible to define lambda functions in the template language and to use map, reduce and filter to handle lists easily.
In addition, the extension allows for currying as a second way of creating custom functions.

The following example creates list of all initial runes of a list of strings.
Thus, if the input is `["Hi", "Hello", "Greetings"]`, the output is `["G", "h"]`.

```
map (prefix 1) | uniq | sort
```

In this example, the function `prefix` expects two arguments: the length of the prefix and the string.
The term `prefix 1` provides only one argument. Thus, the result is a new function that expects a string as input.
This function is now applied by `map` to all elements of the list. The intermediate result is `["H", "H", "G"]`.
The remaining pipeline is used to form a sorted unique set of strings resulting in `["G", "H"]`.

The same result can be achieved without currying, but currying allows for more readable code.

```
map &(prefix 1 .) | uniq | sort
```

Here `&(...)` indicates that the following expression should be treated like a new function.
This new function has only one parameter that is accessed via the `dot` operator.
Functions defined this way can accept any number of arguments, but at least one.
The first argument is passed to the dot as seen in the above example.
In addition, arguments are accessible via `$0`, `$1` etc.
Hence, the above example can be rewritten to>

```
map &(prefix 1 $0) | uniq | sort
```

The differences between the `&(...)` syntax and currying is more than syntax.
In the case of currying, the arguments are evaluated when the function is defined.
In the case of a function, the arguments are evaluated when the function is executed.

```Go
$v := "A"
// The following line uses currying, because hasPrefix expects two arguments.
// Due to currying, `$c` captures the value of `$v`.
$c := hasPrefix $v

$v = "B"
// The following line creates a new function.
// It does not capture the value of `$v`.
$f := `&(hasPrefix $v .)

// Now the following conditions hold:

$f "Adam"  // false, because the first argumnent to hasPrefix is $v and that is "B" currently.
$c "Adam"  // true, because the first argument to hasPrefix is "A".
```

Another example of currying uses `reduce`.
Here we assume the dot is a list of integers and the `add` function adds two integers.

```
reduce add .
```

The template engine detects that `add` is a function and will immediately invoke it with zero parameters
(this is default behavior of the original Go template engine).
The original Go template engine will raise an error, because arguments are missing for `add`.
With the extended version, no error is raised.
Instead, currying means that `add` remains a function, which expects two arguments.
This function is called by `reduce` repeatedly to sum up all elements of the list.

The same can be achieved without currying, but with a functions instead.

```
reduce &(add $0 $1) .
```

Here, `&(add $0 $1)` is a lambda function that consumes two parameters.
The `&(...)` operator will keep the template engine from evaluating the following term immediately and
wraps it in a function instead.