package helpers

type Implementation interface {
	Sign(privateKey []byte, message []byte) []byte
}
