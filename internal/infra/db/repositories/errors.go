package repositories

import (
	"errors"

	"github.com/avast/retry-go"
)

var ErrCustomerNotUpdated = errors.New("customer not updated")

var ErrCustomerNotFound = retry.Unrecoverable(errors.New("customer not found"))
