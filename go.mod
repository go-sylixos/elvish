module src.elv.sh

require (
	github.com/creack/pty v1.1.15
	github.com/google/go-cmp v0.5.9
	github.com/mattn/go-isatty v0.0.17
	github.com/sourcegraph/go-lsp v0.0.0-20200429204803-219e11d77f5d
	github.com/sourcegraph/jsonrpc2 v0.2.0
	go.etcd.io/bbolt v1.3.7
	golang.org/x/sys v0.18.0
)

replace golang.org/x/sys v0.18.0 => ../golang.org/x/sys

replace github.com/mattn/go-isatty v0.0.17 => ../github.com/mattn/go-isatty

replace go.etcd.io/bbolt v1.3.7 => ../bbolt

go 1.23
