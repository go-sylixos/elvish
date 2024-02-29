module src.elv.sh/website

go 1.18

require (
	github.com/BurntSushi/toml v1.2.1
	github.com/creack/pty v1.1.18
	github.com/google/go-cmp v0.5.9
	src.elv.sh v0.18.0
)

require (
	github.com/mattn/go-isatty v0.0.17 // indirect
	golang.org/x/sys v0.6.0 // indirect
)

replace golang.org/x/sys v0.6.0 => ../../golang.org/x/sys

replace github.com/mattn/go-isatty v0.0.17 => ../../github.com/mattn/go-isatty

replace src.elv.sh => ../
