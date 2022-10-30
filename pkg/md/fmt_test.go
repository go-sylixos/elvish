package md_test

import (
	"html"
	"testing"

	"github.com/google/go-cmp/cmp"
	"src.elv.sh/pkg/md"
	"src.elv.sh/pkg/testutil"
)

func TestFmtPreservesHTMLRender(t *testing.T) {
	testutil.Set(t, &md.UnescapeEntities, html.UnescapeString)
	for _, tc := range testCases {
		t.Run(tc.testName(), func(t *testing.T) {
			tc.skipIfNotSupported(t)
			if tc.Name != "" {
				t.Skip("TODO supplemental cases")
			}
			switch tc.Example {
			case 39, 40:
				t.Skip("TODO escape sequence")
			case 488, 491, 497, 498, 499, 519:
				t.Skip("TODO link and image formatting")
			}
			testFmtPreservesHTMLRender(t, tc.Markdown)
		})
	}
}

func testFmtPreservesHTMLRender(t *testing.T, original string) {
	formatted := render(original, &md.FmtCodec{})
	formattedRender := render(formatted, &htmlCodec{})
	originalRender := render(original, &htmlCodec{})
	if formattedRender != originalRender {
		t.Errorf("original:\n%s\nformatted:\n%s\n"+
			"HTML diff (-original +formatted):\n%sops diff (-original +formatted):\n%s",
			hr+"\n"+original+hr, hr+"\n"+formatted+hr,
			cmp.Diff(originalRender, formattedRender),
			cmp.Diff(render(original, &md.OpTraceCodec{}), render(formatted, &md.OpTraceCodec{})))
	}
}
