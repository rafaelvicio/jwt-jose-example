package main

import (
	"fmt"
	"time"

	"github.com/SermoDigital/jose/jws"
)

func main() {
	expires := time.Now().Add(time.Duration(10) * time.Second)

	claims := jws.Claims{}
	claims.SetExpiration(expires)
	claims.SetIssuedAt(time.Now())

	fmt.Println(claims)
}
