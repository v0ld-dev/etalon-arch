package otp


type Generator interface {
	RandomSecret(length int) string
}
