package helpers

import (
	"testing"
)

func DoBenchmarkFunc(
	b *testing.B,
	signFn func(privateKey []byte, message []byte) []byte,
) {
	b.ReportAllocs()

	priv, message := getStuff()

	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		signFn(priv, message)
	}
}

func DoBenchmark(
	b *testing.B,
	implFuncs Implementation,
){
	DoBenchmarkFunc(b, implFuncs.Sign)
}
