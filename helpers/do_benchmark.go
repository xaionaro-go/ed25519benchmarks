package helpers

import (
	"crypto/ed25519"
	"testing"
)

func DoBenchmarkFunc(
	b *testing.B,
	signFn func(privateKey []byte, message []byte) []byte,
) {
	b.ReportAllocs()

	_, priv, err := ed25519.GenerateKey(ZeroReader{})
	if err != nil {
		b.Fatal(err)
	}
	message := GenMessage()

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
