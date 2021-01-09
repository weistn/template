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