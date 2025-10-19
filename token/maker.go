package token

import "time"

// Switch between JWT and PASETO tokens
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}
