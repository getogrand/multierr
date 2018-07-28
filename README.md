# multierr

[![GoDoc](https://godoc.org/github.com/getogrand/multierr?status.svg)](https://godoc.org/github.com/getogrand/multierr)
[![Build Status](https://travis-ci.org/getogrand/multierr.svg?branch=master)](https://travis-ci.org/getogrand/multierr)

Package multierr introduce a simple way to join multiple errors as an error.

## Installation

Run `go get github.com/getogrand/multierr`.

## Usage

The `multierr.Join()` combine multiple errors to an error.

You can report multiple errors as an joined error to the caller using `multierr.Join()`.

```go
import "github.com/getogrand/multierr"

var ee []error

up := ReservationUserPush{}
n := ReservationAlarmNoti{}

if err := up.Send(); err != nil {
  ee = append(ee,
    fmt.Errorf("send reservation %v push to user: %v", n.Reservation.ID, err))
}

if err := n.sendShopAlarm(); err != nil {
  ee = append(ee, fmt.Errorf(
    "send shop alarm message of reservation %v: %v", n.Reservation.ID, err))
}

if err := n.sendUserAlarm(); err != nil {
  ee = append(ee, fmt.Errorf(
    "send user alarm message of reservation %v: %v", n.Reservation.ID, err))
}

return multierr.Join(ee) // error{"3 errors occured: send reservation 1 push to user: connection fail, send shop alarm message of reservation 1: connection fail, send user alarm message of reservation 1: connection fail"}
```

It is really useful when you run errorable operation in for-loop.

```go
import "github.com/getogrand/multierr"

func SliceConvAtoi32(aa []string) ([]int32, error) {
	errs := []error{}
	ii := []int32{}
	for _, a := range aa {
		i, err := strconv.Atoi(a)
		if err != nil {
			errs = append(errs, fmt.Errorf("convert %q to int: %v", a, err))
			continue
		}
		ii = append(ii, int32(i))
	}

	if len(errs) > 0 {
		return []int32{}, multierr.Join(errs)
	}
	return ii, nil
}
```

## Extracted From [heybeauty](https://heybeauty.me)

## License

Released under the [MIT License](https://github.com/getogrand/multierr/blob/master/License)
