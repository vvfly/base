//go:build linux
// +build linux

package program_rlimit_tuning

import (
	_ "go.uber.org/automaxprocs"
	"golang.org/x/sys/unix"
)

func init() {
	if err := Tuning(); err != nil {
		panic(err)
	}
}

const (
	defaultFileMax = 1048576
)

func Tuning() error {
	var (
		rLimit unix.Rlimit
		err    error
	)

	if err = unix.Getrlimit(unix.RLIMIT_NOFILE, &rLimit); err != nil {
		return err
	}

	if rLimit.Max < defaultFileMax || rLimit.Cur < defaultFileMax {
		rLimit.Max = defaultFileMax
		rLimit.Cur = defaultFileMax
	}

	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rLimit)
}
