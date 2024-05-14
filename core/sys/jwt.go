package sys

import (
	"cos/core/service"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/square/go-jose"
)

type Scope string

const (
	ScopeAuth  Scope = "auth"
	ScopeAdmin Scope = "admin"
)

var PrivateKey, _ = rsa.GenerateKey(rand.Reader, 2048)

type AuthContext struct {
	User   int64     `json:"user"`
	Scope  Scope     `json:"scope"`
	Expire time.Time `json:"expire"`
}

func CreateAuthKey(context AuthContext) (key string, err error) {

	if context.User == 0 {
		return "", errors.New("expected user not empty")
	}

	if !context.Expire.After(time.Now()) {
		return "", errors.New("expected valid expiration")
	}

	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.PS512, Key: PrivateKey}, nil)
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(context)

	if err != nil {
		return "", err
	}

	object, err := signer.Sign(data)
	if err != nil {
		return "", err
	}

	return object.CompactSerialize()
}

func AuthContextFromRequest(c service.Context) (context *AuthContext, err error) {

	key := c.Query("access_token")

	if key == "" {

		bearer := strings.Split(c.Header("Authorization"), "Bearer ")

		if len(bearer) == 2 {
			key = bearer[1]
		}
	}

	if key == "" {
		return nil, errors.New("not authorized")
	}

	return AuthContextFromKey(key)
}

func AuthContextFromKey(key string) (context *AuthContext, err error) {

	object, err := jose.ParseSigned(key)
	if err != nil {
		return nil, err
	}

	data, err := object.Verify(&PrivateKey.PublicKey)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &context)

	return
}
