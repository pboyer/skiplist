# Benchmarks

On a mid-2015 Macbook Pro:

```
λ go test -run=NO_TEST -bench . -benchmem -benchtime 1s ./...
BenchmarkPut100-8                  30000             70578 ns/op            6532 B/op        200 allocs/op
BenchmarkPut1000-8                  5000            903507 ns/op           65330 B/op       2000 allocs/op
BenchmarkPut10000-8                  300           6021168 ns/op          653359 B/op      20000 allocs/op
BenchmarkPutRemoveAll100-8         50000             34872 ns/op            8750 B/op        210 allocs/op
BenchmarkPutRemoveAll1000-8         5000            336889 ns/op           81887 B/op       2013 allocs/op
BenchmarkPutRemoveAll10000-8         500           3462277 ns/op         1039832 B/op      20022 allocs/op
BenchmarkPutRemove-8             5000000               253 ns/op              65 B/op          2 allocs/op
PASS
ok      github.com/pboyer/skiplist      16.902s
```