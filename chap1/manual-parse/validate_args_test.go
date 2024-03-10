package main

import (
	"errors"
	"testing"
)

type validateArgsTestConfig struct {
	c   config
	err error
}

func TestValidateArgs(t *testing.T) {
	tests := []validateArgsTestConfig{
		{
			c:   config{},
			err: errors.New("Must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: -1},
			err: errors.New("Must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: 10},
			err: nil,
		},
	}

	for _, tc := range tests {
		err := validateArgs(tc.c)

		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Errorf("Expected error to be: %v, got: %v\n", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Errorf("Expected nillerror, got: %v\n", err)
		}
	}
}
