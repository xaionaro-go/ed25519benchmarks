package gostandard

import (
	"testing"

	"github.com/oasislabs/ed25519"
	"github.com/xaionaro-go/ed25519benchmarks/helpers"
)

type implementation struct{}
func (impl *implementation) Sign(privateKey []byte, message []byte) []byte {
	return ed25519.Sign(privateKey, message)
}

func BenchmarkEd25519Sign_oasislabs(b *testing.B) {
	helpers.DoBenchmark(b, &implementation{})
}
