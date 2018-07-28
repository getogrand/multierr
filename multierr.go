// Package multierr introduce a simple way to join multiple errors as an error.
package multierr

import (
	"fmt"
	"strings"
)

// Join combine multiple errors to an error.
//
// Error that joined have error string follow below format.
//
// `2 errors occured: open file abc.png, close connection of db 'test_db'`
func Join(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	var msgs []string
	for _, e := range errs {
		msgs = append(msgs, e.Error())
	}
	joinedmsg := strings.Join(msgs, ", ")
	return fmt.Errorf("%d errors occured: %s", len(errs), joinedmsg)
}
