package helper

import "errors"

var ErrUserAlreadyExists = errors.New("User already exists")
var ErrEmailOrPasswordIncorrect = errors.New("Incorrect Username or Email")
