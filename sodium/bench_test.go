package gostandard

import (
	"testing"

	"github.com/jamesruan/sodium"
	"github.com/xaionaro-go/ed25519benchmarks/helpers"
)

type implementation struct{}
func (impl *implementation) Sign(privateKey []byte, message []byte) []byte {
	state := sodium.MakeSignState()
	state.Update(message)
	return state.Sign(sodium.SignSecretKey{privateKey}).Bytes
}

func BenchmarkEd25519Sign_sodium(b *testing.B) {
	helpers.DoBenchmark(b, &implementation{})
}
