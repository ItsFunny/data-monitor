package utils

import (
	"errors"
	"fmt"
)

func MissingArgumentError(key, value string) error {
	return errors.New(fmt.Sprintf("missing %v of %v ", value, key))
}

func MissingArgument(v string) error {
	return errors.New(fmt.Sprintf("missing argument:%v", v))
}

func ParseError(key, value interface{}) error {
	return errors.New(fmt.Sprintf("when parse %v to %v occur error", value, key))
}
