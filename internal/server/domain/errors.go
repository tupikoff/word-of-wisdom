package domain

import "github.com/pkg/errors"

var ErrHashNotValid = errors.New("hash is not valid")
var ErrHashReadError = errors.New("hash read error")
var ErrHashStringNotRegistered = errors.New("hash string not registered")
var ErrUnknownCommand = errors.New("unknown command")
var ErrRecordAlreadyExists = errors.New("record already exists")
var ErrDifficultyNotMatch = errors.New("difficulty not match")
