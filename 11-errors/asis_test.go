package main

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ErrInvalidValue = errors.New("invalid value")

func TestIs(t *testing.T) {
	var err error
	err = ErrInvalidValue
	err = fmt.Errorf("wrapped: %w", err)

	// ---

	equalsWorks := err == ErrInvalidValue
	assert.False(t, equalsWorks)

	isWorks := errors.Is(err, ErrInvalidValue)
	assert.True(t, isWorks)
}

type ErrorWithTimeout interface {
	error
	Timeout() bool
}

func TestIsAs_CustomType(t *testing.T) {
	var err error
	err = &os.PathError{Op: "open", Path: "file.txt", Err: errors.New("file not found")}
	err = fmt.Errorf("wrapped: %w", err)

	// ---

	isWorks := errors.Is(err, &os.PathError{})
	assert.False(t, isWorks)

	var pathErr *os.PathError
	asWorks := errors.As(err, &pathErr)
	assert.True(t, asWorks)

	var ifaceWorks ErrorWithTimeout
	asWorksWithIface := errors.As(err, &ifaceWorks)
	assert.True(t, asWorksWithIface)

}
