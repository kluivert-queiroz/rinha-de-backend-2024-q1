package entities

import (
	"errors"

	"github.com/avast/retry-go"
)

var ErrNotEnoughLimit = retry.Unrecoverable(errors.New("user doesn't have enough limit to proceed"))
