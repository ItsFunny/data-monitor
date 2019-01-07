package utils

import "errors"

func WrapBaseError(baseError error, msg string) error {
	msg += baseError.Error()
	return errors.New(msg)
}
