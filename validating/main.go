package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

func generatedToken() []byte {
	bytes, _ := ioutil.ReadFile("./sample_key.priv")

	claims := jws.Claims{}
	claims.SetExpiration(time.Now().Add(time.Duration(10) * time.Second))

	rsaPrivate, _ := crypto.ParseRSAPrivateKeyFromPEM(bytes)
	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)

	b, _ := jwt.Serialize(rsaPrivate)
	return b
}

func main() {

	bytes, _ := ioutil.ReadFile("./sample_key.pub")
	rsaPublic, _ := crypto.ParseRSAPublicKeyFromPEM(bytes)

	accessToken := generatedToken()
	jwt, err := jws.ParseJWT([]byte(accessToken))
	if err != nil {
		log.Fatal(err)
	}

	// Validate token
	if err = jwt.Validate(rsaPublic, crypto.SigningMethodRS256); err != nil {
		log.Fatal(err)
	}
}
