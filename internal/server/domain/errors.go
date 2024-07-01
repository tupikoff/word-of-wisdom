package domain

import "github.com/pkg/errors"

var ErrHashNotValid = errors.New("hash is not valid")
var ErrHashReadError = errors.New("hash read error")
var ErrRecordAlreadyExists = errors.New("record already exists")
var ErrDifficultyNotMatch = errors.New("difficulties not match")
var ErrRandomStringNotMatch = errors.New("random strings not match")
