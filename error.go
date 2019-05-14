package dgoflake

import "errors"

func newError(err string, cause error) error {
	if cause != nil {
		return errors.New("dgoflake: " + err + cause.Error())
	}

	return errors.New("dgoflake: " + err)
}
