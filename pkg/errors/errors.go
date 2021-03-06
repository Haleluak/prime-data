package errors

import (
	"github.com/pkg/errors"
	"fmt"
)

type CustomError struct {
	errType      ErrorType
	wrappedError error
	context      errorContext
}

func (err CustomError) Error() string {
	return err.wrappedError.Error()
}

func (err CustomError) Stacktrace() string {
	return fmt.Sprintf("%+v\n", err.wrappedError)
}

// New creates a no type error
func New(msg string) error {
	return CustomError{errType: Error, wrappedError: errors.New(msg)}
}

// Newf creates a no type error with formatted message
func Newf(msg string, args ...interface{}) error {
	return CustomError{errType: Error, wrappedError: errors.New(fmt.Sprintf(msg, args...))}
}

// Wrap wrans an error with a string
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Cause gives the original error
func Cause(err error) error {
	return errors.Cause(err)
}

// Wrapf wraps an error with format string
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(CustomError); ok {
		return CustomError{
			errType:      customErr.errType,
			wrappedError: wrappedError,
			context:      customErr.context,
		}
	}

	return CustomError{errType: Error, wrappedError: wrappedError}
}

// Get Stacktrace of error
func Stack(err error) string {
	if customErr, ok := err.(CustomError); ok {
		return fmt.Sprintf("%+v\n", customErr.wrappedError)
	}
	return fmt.Sprintf("%+v\n", errors.WithStack(err))
}