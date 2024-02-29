//go:build sylixos

package unix

import "golang.org/x/sys/unix"

var rlimitKeys = map[int]string{
	// The following are defined by POSIX
	// (https://pubs.opengroup.org/onlinepubs/9699919799/functions/getrlimit.html).
	//
	// Note: RLIMIT_AS is defined by POSIX, but missing on OpenBSD
	// (https://man.openbsd.org/getrlimit.2#BUGS); it is defined on Darwin, but
	// it's an undocumented alias of RLIMIT_RSS there.
	unix.RLIMIT_CORE:   "core",
	unix.RLIMIT_CPU:    "cpu",
	unix.RLIMIT_DATA:   "data",
	unix.RLIMIT_FSIZE:  "fsize",
	unix.RLIMIT_NOFILE: "nofile",
	unix.RLIMIT_STACK:  "stack",
}

//lint:ignore U1000 used on some OS
func addRlimitKeys(m map[int]string) {
	for k, v := range m {
		rlimitKeys[k] = v
	}
}
