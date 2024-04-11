# github.com/Guest-615695028/comp
Go support for programmers familiar with other common programming languages, such as C/C++, Java, Javascript, Python.
## missing operators
1. Conditional operator `a ? b : c` using `Cond(a, b, c)`
2. Null-checking operator `a ?? b` *i. e.* `a!=null ? a : b` using `Cond(a, b, c)`
3. Case to Boolean type: `Bool(v)`
   - Numbers: whether `v` is non-zero;
   - Strings: whether `v` is non-empty;
   - `nil`-Testible Objects *viz.* pointer, channel, function, interface, map, and slice types: `nil` to `false`, others to `true`;
   - Arrays and structures: always true as for JavaScript empty array `[]` and empty object `{}`.
