package auth

import (
	"admin-server/types"
	"math/rand"
)

type DAO struct {
}

type DAOInterface interface {
	GetUserFromToken(token string) (user *types.User, err error)
}

func NewDAO() DAOInterface {
	return &DAO{}
}

func (d *DAO) GetUserFromToken(token string) (user *types.User, err error) {
	// TODO
	return
}

func (d *DAO) GenerateTokenForUser(user types.LoginUser) (token *string, err error) {
	// TODO

	// TODO validate user

	var str string
	for {
		str = GenerateRandomString(128)

		// TODO check if string is already in use

		// TODO insert in db

		token = &str
		return token, nil
	}
}

// TODO remove

const ValidChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateRandomString(length uint32) (out string) {
	out = ""
	for i := uint32(0); i < length; i++ {
		out += string(ValidChars[int(rand.Uint32())%len(ValidChars)])
	}
	return out
}
