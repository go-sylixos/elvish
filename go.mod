module src.elv.sh

require (
	github.com/creack/pty v1.1.21
	github.com/google/go-cmp v0.6.0
	github.com/mattn/go-isatty v0.0.20
	github.com/sourcegraph/jsonrpc2 v0.2.0
	go.etcd.io/bbolt v1.3.8
	golang.org/x/sync v0.7.0
	golang.org/x/sys v0.19.0
	pkg.nimblebun.works/go-lsp v1.1.0
)

replace golang.org/x/sys v0.19.0 => ../golang.org/x/sys

replace github.com/mattn/go-isatty v0.0.20 => ../github.com/mattn/go-isatty

replace go.etcd.io/bbolt v1.3.8 => ../bbolt

go 1.23
