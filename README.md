# Select Rand

Implements the `math/rand.Source` and `math/rand.Source64` interfaces using
Go's [select](https://golang.org/ref/spec#Select_statements) statement.

While the random number generator is deterministic, it cannot be seeded and
the sequence cannot be controlled or reproduced.

The performance is poor. Comparison with the `math/rand` implementation:

```
$ go test -bench=.
BenchmarkSelectRand-4             200000              8093 ns/op
BenchmarkMathRand-4             300000000                4.67 ns/op
PASS
ok      github.com/peterstace/selectrand        3.597s
```
