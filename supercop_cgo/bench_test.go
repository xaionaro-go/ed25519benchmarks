package supercop_cgo

import (
	"testing"

	"github.com/oasislabs/ed25519"
	"github.com/xaionaro-go/ed25519benchmarks/helpers"
	supercop "github.com/xaionaro-go/supercop/crypto_sign/ed25519/amd64-51-30k"
)

type implementation struct{}
func (impl *implementation) Sign(privateKey []byte, message []byte) []byte {
	result := make([]byte, ed25519.SignatureSize + len(message))
	err := supercop.CryptoSign(result, message, privateKey)
	if err != nil {
		panic(err)
	}
	return result[:ed25519.SignatureSize]
}

func BenchmarkEd25519Sign_supercop_cgo(b *testing.B) {
	helpers.DoBenchmark(b, &implementation{})
}

func TestSign(t *testing.T) {
	helpers.DoTest(t, &implementation{})
}