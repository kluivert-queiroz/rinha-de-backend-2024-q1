package entities

import "errors"

var ErrNotEnoughLimit = errors.New("user doesn't have enough limit to proceed")
