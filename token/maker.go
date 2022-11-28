package token

import "time"

type Maker interface {
	// CreateToken creates new token with specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//verifyToken verifies the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
