package token

import "time"

// Switch between JWT and PASETO tokens
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)

	VeriftToken(token string) (*Payload, error)
}
