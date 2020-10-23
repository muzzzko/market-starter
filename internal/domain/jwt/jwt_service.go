package jwt

import (
	"time"
)

type HMAC struct {
	secret []byte
	expiredAt time.Duration
}



func NewHMAC(secret []byte, expiredAt time.Duration) *HMAC {
	return &HMAC{
		secret: secret,
		expiredAt: expiredAt,
	}
}

