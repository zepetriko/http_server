package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestHashAndCheckPassword(t *testing.T) {
	password := "supersecret123"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error %v", err)
	}

	match, err := CheckPasswordHash(password, hash)
	if err != nil {
		t.Fatalf("CheckPasswordHash returned error: %v", err)
	}

	if !match {
		t.Fatal("Expected password to match hash")
	}
}

func TestWrongPassword(t *testing.T) {
	password := "123456"
	wrongPassword := "000011112222"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	match, err := CheckPasswordHash(wrongPassword, hash)
	if err != nil {
		t.Fatalf("CheckPasswordHash returned error: %v", err)
	}

	if match {
		t.Fatal("expected password not to match hash")
	}
}

func TestHashesAreDifferent(t *testing.T) {
	password := "supersecret123"

	hash1, err := HashPassword(password)
	if err != nil {
		t.Fatal(err)
	}

	hash2, err := HashPassword(password)
	if err != nil {
		t.Fatal(err)
	}

	if hash1 == hash2 {
		t.Fatal("expected hashses to be different due to random salt")
	}
}

func TestJWT(t *testing.T) {
	userID := uuid.New()
	secret := "my-secret"

	token, err := MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	gotID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatal(err)
	}

	if gotID != userID {
		t.Fatalf("expected %s, got %s", userID, gotID)
	}
}

func TestJWTWrongSecret(t *testing.T) {
	userID := uuid.New()

	token, err := MakeJWT(userID, "secret1", time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ValidateJWT(token, "secret2")
	if err == nil {
		t.Fatal("this returns means no error occured, but should")
	}
}

func TestJWTExpired(t *testing.T) {
	userID := uuid.New()

	token, err := MakeJWT(userID, "secret", -time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ValidateJWT(token, "secret")
	if err == nil {
		t.Fatal("this returns means no error occured, but should")
	}

}

func TestGetBearerToken(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer my-token")

	token, err := GetBearerToken(headers)
	if err != nil {
		t.Fatal(err)
	}

	if token != "my-token" {
		t.Fatalf("expected my-token, got %s", token)
	}
}

func TestGetBearerTokenMissingHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetBearerToken(headers)
	if err == nil {
		t.Fatal("token should be missing header")
	}
}

func TestGetBearerTokenBadFormat(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "my-token")

	_, err := GetBearerToken(headers)
	if err == nil {
		t.Fatal("token should be in a bad format")
	}
}
