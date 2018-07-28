package multierr_test

import (
	"fmt"
	"testing"

	"github.com/getogrand/multierr"
	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	errors := []error{
		fmt.Errorf("this is error 1"),
		fmt.Errorf("this is error 2"),
		fmt.Errorf("this is error 3"),
	}

	joinderr := multierr.Join(errors)
	assert.Error(t, joinderr)
	assert.Equal(t,
		"3 errors occured: this is error 1, this is error 2, this is error 3",
		joinderr.Error())
}
