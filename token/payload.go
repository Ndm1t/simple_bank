package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.ExpiredAt), nil
}

func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.IssuedAt), nil
}

func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	//TODO implement me
	return nil, nil
}

func (payload *Payload) GetIssuer() (string, error) {
	//TODO implement me
	return "", nil
}

func (payload *Payload) GetSubject() (string, error) {
	//TODO implement me
	return "", nil
}

func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	//TODO implement me
	return nil, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenId,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
