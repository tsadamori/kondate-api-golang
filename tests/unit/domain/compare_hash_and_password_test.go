package crypto_test

import (
	"testing"

	"github.com/tsadamori/kondate-api-golang/auth"
)

func TestCompareHashAndPassword(t *testing.T) {
	t.Run("Match password and hash", func(t *testing.T) {
		hash := "$2a$10$40tl6HzxmljVgDtiUKp9gO2DUHQ8vV1xGAdH9uchwmVUg8gmffaq2"
		password := "test"

		err := auth.CompareHashAndPassword(hash, password)
		if err != nil {
			t.Errorf("%s; hash=%v, password=%v", t.Name(), hash, password)
		}
	})
	t.Run("Does not match password and hash", func(t *testing.T) {
		hash := "$2a$10$40tl6HzxmljVgDtiUKp9gO2DUHQ8vV1xGAdH9uchwmVUg8gmffaq2"
		password := "test1"

		err := auth.CompareHashAndPassword(hash, password)
		if err == nil {
			t.Errorf("%s; hash=%v, password=%v",t.Name(), hash, password)
		}
	})
}