package errors

import (
	"errors"
	"fmt"
)

// New implements errros.New
func New(text string) error {
	return errors.New(text)
}

// Is implements errors.Is
func Is(err error, target error) bool {
	return errors.Is(err, target)
}

// As implements errors.As
func As(err error, target any) bool {
	return errors.As(err, target)
}

// Wrap returns new error wrap with message in format "%s: %w"
func Wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

// Unwrap implements errros.Unwrap
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Join implements errors.Join
func Join(errs ...error) error {
	return errors.Join(errs...)
}

// IsUnsupported implements short cut for errors.Is(err, errros.ErrUnsupported)
func IsUnsupported(err error) bool {
	return Is(err, errors.ErrUnsupported)
}
