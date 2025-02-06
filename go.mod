module src.elv.sh

require (
	github.com/creack/pty v1.1.21
	github.com/google/go-cmp v0.6.0
	github.com/mattn/go-isatty v0.0.20
	github.com/sourcegraph/jsonrpc2 v0.2.0
	go.etcd.io/bbolt v1.3.8
	golang.org/x/sync v0.10.0
	golang.org/x/sys v0.28.0
	pkg.nimblebun.works/go-lsp v1.1.0
)

replace golang.org/x/sys v0.28.0 => github.com/go-sylixos/sys v0.28.1-sylixos.1

replace github.com/mattn/go-isatty v0.0.20 => github.com/go-sylixos/go-isatty v0.0.21-sylixos.1

replace go.etcd.io/bbolt v1.3.8 => ../bbolt

go 1.23
