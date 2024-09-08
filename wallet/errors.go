package wallet

import "errors"

var (
	ErrNoSuchUser          = errors.New("user doesn't exists")
	ErrWalletAlreadyExists = errors.New("wallet already exists")
)
