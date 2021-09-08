package client

import "fmt"

func wrapError(msg string, oErr error) error {
	return fmt.Errorf("%s : %v", msg, oErr)
}
