package domain

import "github.com/pkg/errors"

var ErrHashNotValid = errors.New("hash is not valid")
var ErrHashReadError = errors.New("hash read error")
var ErrHashStringNotRegistered = errors.New("hash string not registered")
var ErrUnknownCommand = errors.New("unknown command")
var ErrAlreadyExists = errors.New("already exists")
var ErrDifficultyNotMatchWithRegistered = errors.New("difficulty not match with registered")
