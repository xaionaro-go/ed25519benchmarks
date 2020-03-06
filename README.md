```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/ed25519benchmarks/gostandard
BenchmarkEd25519Sign_standard-8   	    8072	    197749 ns/op	     512 B/op	       6 allocs/op
PASS
ok  	github.com/xaionaro-go/ed25519benchmarks/gostandard	1.661s
?   	github.com/xaionaro-go/ed25519benchmarks/helpers	[no test files]
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/ed25519benchmarks/oasislabs
BenchmarkEd25519Sign_oasislabs-8   	   15153	    105726 ns/op	     480 B/op	       5 allocs/op
PASS
ok  	github.com/xaionaro-go/ed25519benchmarks/oasislabs	2.432s
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/ed25519benchmarks/sodium
BenchmarkEd25519Sign_sodium-8   	   10000	    122964 ns/op	     104 B/op	       3 allocs/op
PASS
ok  	github.com/xaionaro-go/ed25519benchmarks/sodium	1.265s
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/ed25519benchmarks/supercop_cgo
BenchmarkEd25519Sign_supercop_cgo-8   	   23888	     64777 ns/op	      96 B/op	       2 allocs/op
PASS
ok  	github.com/xaionaro-go/ed25519benchmarks/supercop_cgo	2.104s
```
