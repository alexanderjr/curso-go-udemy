package contract

type Encryptor interface {
	ToHash(value string) ([]byte, error)
	IsHashEqualsTo(hashValue, value string) error
}
