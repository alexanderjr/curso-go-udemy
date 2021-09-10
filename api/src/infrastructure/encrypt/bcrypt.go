package encrypt

import "golang.org/x/crypto/bcrypt"

type BCryptConcrete struct{}

func (b BCryptConcrete) ToHash(value string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
}

func (b BCryptConcrete) IsHashEqualsTo(hashedValue, value string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value))
}
