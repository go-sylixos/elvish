// The macros program implements an ad-hoc preprocessor for Markdown files, used
// in Elvish's website.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"src.elv.sh/pkg/mods/doc"
)

func filter(in io.Reader, out io.Writer) {
	f := filterer{}
	f.filter(in, out)
}

type filterer struct {
	module string
}

var macros = map[string]func(*filterer, string) string{
	"@module ":  (*filterer).expandModule,
	"@ttyshot ": (*filterer).expandTtyshot,
	"@cf ":      (*filterer).expandCf,
	"@dl ":      (*filterer).expandDl,
}

func (f *filterer) filter(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		for leader, expand := range macros {
			i := strings.Index(line, leader)
			if i >= 0 {
				line = line[:i] + expand(f, line[i+len(leader):])
				break
			}
		}
		fmt.Fprintln(out, line)
	}
	if f.module != "" {
		ns := f.module + ":"
		if f.module == "builtin" {
			ns = ""
		}
		docs := doc.Docs()[ns]

		var buf bytes.Buffer
		writeElvdocSections(&buf, ns, docs)
		filter(&buf, out)
	}
}

func (f *filterer) expandModule(rest string) string {
	f.module = rest
	// Module doc will be added at end of file
	return fmt.Sprintf(
		"<a name='//apple_ref/cpp/Module/%s' class='dashAnchor'></a>", f.module)
}

func (f *filterer) expandTtyshot(name string) string {
	content, err := os.ReadFile(name + ".ttyshot.html")
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf(`<pre class="ttyshot"><code>%s</code></pre>`,
		bytes.Replace(content, []byte("\n"), []byte("<br>"), -1))
}

func (f *filterer) expandCf(rest string) string {
	targets := strings.Split(rest, " ")
	var buf strings.Builder
	buf.WriteString("See also")
	for i, target := range targets {
		if i == 0 {
			buf.WriteString(" ")
		} else if i == len(targets)-1 {
			buf.WriteString(" and ")
		} else {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "[`%s`](%s)", target, cfHref(target, f.module))
	}
	buf.WriteString(".")
	return buf.String()
}

// Returns the href for a `@cf` reference.
func cfHref(symbol, currentModule string) string {
	i := strings.IndexRune(symbol, ':')
	if i == -1 {
		// An internal link in the builtin module's doc.
		return "#" + symbol
	}

	var module, unqualified string
	if strings.HasPrefix(symbol, "$") {
		module, unqualified = symbol[1:i], "$"+symbol[i+1:]
	} else {
		module, unqualified = symbol[:i], symbol[i+1:]
	}
	switch module {
	case "builtin":
		// A link from a non-builtin module's doc to the builtin module. Use
		// unqualified name (like #put or #$paths, instead of #builtin:put or
		// #$builtin:paths).
		return "builtin.html#" + unqualified
	case currentModule:
		// An internal link in a non-builtin module's doc.
		return "#" + unqualified
	default:
		// A link to a non-builtin module.
		return module + ".html#" + symbol
	}
}

func (f *filterer) expandDl(rest string) string {
	fields := strings.SplitN(rest, " ", 2)
	name := fields[0]
	url := name
	if len(fields) == 2 {
		url = fields[1]
	}
	return fmt.Sprintf(`<a href="https://dl.elv.sh/%s">%s</a>`, url, name)
}