# github.com/Guest-615695028/comp
Go support for programmers familiar with other common programming languages, such as C/C++, Java, Javascript, Python.
## missing operators
1. `Cond(a, b, c)`: Conditional operator `a ? b : c` 
2. `Null(a, b)`: Nil-checking operator `a ?? b` *i. e.* `a!=nil ? a : b`
3. `Value(value, error)`: Returns the error if so, value if not.
4. Case to Boolean type: `Bool(v)`
   - Numbers: whether `v` is non-zero;
   - Strings: whether `v` is non-empty;
   - `nil`-Testible Objects *viz.* pointer, channel, function, interface, map, and slice types: `nil` to `false`, others to `true`;
   - Arrays and structures: always true as for JavaScript empty array `[]` and empty object `{}`.
