package helpers

import (
	"crypto/ed25519"
	"testing"

	"github.com/stretchr/testify/require"
)

func getStuff() (privKey, message []byte) {
	_, priv, err := ed25519.GenerateKey(ZeroReader{})
	if err != nil {
		panic(err)
	}
	return priv, GenMessage()
}

func DoTest(t *testing.T, impl Implementation) {
	privateKey, message := getStuff()
	referenceSignature := ed25519.Sign(privateKey, message)
	resultSignature := impl.Sign(privateKey, message)
	require.Equal(t, referenceSignature, resultSignature)
}
