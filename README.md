```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/ed25519benchmarks/gostandard
BenchmarkEd25519Sign_standard-8   	   26454	     46023 ns/op	     512 B/op	       6 allocs/op
PASS
ok  	github.com/xaionaro-go/ed25519benchmarks/gostandard	1.678s
?   	github.com/xaionaro-go/ed25519benchmarks/helpers	[no test files]
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/ed25519benchmarks/oasislabs
BenchmarkEd25519Sign_oasislabs-8   	   58632	     21660 ns/op	     480 B/op	       5 allocs/op
PASS
ok  	github.com/xaionaro-go/ed25519benchmarks/oasislabs	1.479s
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/ed25519benchmarks/sodium
BenchmarkEd25519Sign_sodium-8   	   39229	     30414 ns/op	     104 B/op	       3 allocs/op
PASS
ok  	github.com/xaionaro-go/ed25519benchmarks/sodium	1.505s
```
