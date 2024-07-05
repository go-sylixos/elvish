//go:build sylixos

package eval

import (
	"syscall"
)

// Process control functions in Unix.

// Nop on SylixOS.
func putSelfInFg() error { return nil }

func makeSysProcAttr(bg bool) *syscall.SysProcAttr {
	return &syscall.SysProcAttr{Setpgid: bg}
}
