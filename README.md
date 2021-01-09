# template
Extended Go template package

This package is a clone of the template package of the Go standard library. 
Thus, this code uses the same license as the original Go code from which it is derived.

# Features

This version adds support for functional programming.
It is possible to define custom functions in the template language and to use map, reduce and filter to handle lists easily.
In addition, the extension allows for currying as a second way of creating custom functions.

The following example creates list of all initial runes of a list of strings.
Thus, if the input is `["Hi", "Hello", "Greetings"]`, the output is `["G", "h"]`.

```
map (prefix 1) | uniq | sort
```

In this example, the function `prefix` expects two arguments: the length of the prefix and the string.
The term `prefix 1` provides only one argument. Thus, the result is a new function that expects a string as input.
This function is now applied by `map` to all elements of the list. The intermediate result is `["H", "H", "G"]`.
The remaining pipeline is used to form a sorted unique set of strings.

The same result can be achieved without currying, but currying allows for more readable code.

```
map `(prefix 1 .) | uniq | sort
```

Here ``\``` indicates that the following expression should be treated like a new function.
This new function has only one parameter that is accessed via the `dot` operator.

However, the differences between the ``\`(...)`` syntax and currying is more than syntax.
In the case of currying, the arguments are evaluated when the function is defined.
In the case of a function, the arguments are evaliated when the function is executed.

```Go
$v := "A"
$c := hasPrefix $v

$v = "B"
$f := `(hasPrefix $v .)

/* Now the following conditions hold */

$f "A"  // false, because the first argumnent to hasPrefix is $v and that is "B" currently.
$c "A"  // true, because the first argument to hasPrefix is "A".
```
