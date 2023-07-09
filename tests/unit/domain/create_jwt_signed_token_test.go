package auth_test

import (
	"testing"

	"github.com/tsadamori/kondate-api-golang/auth"
)

func TestCreateJWTSignedToken(t *testing.T) {
	var expUnix int64 = 1672498800
	email := "sample@example.com"
	exptected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbXBsZUBleGFtcGxlLmNvbSIsImV4cCI6MTY3MjQ5ODgwMH0.wyGNcrKTSyadbnjC8y9Jb_wJJMC7vv8HFhfJpKLqWG8"

	token, err := auth.CreateJWTSignedToken(email, expUnix)
	if token != exptected || err != nil {
		t.Errorf("%s", t.Name())
	}
}