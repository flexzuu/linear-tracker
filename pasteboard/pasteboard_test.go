package pasteboard_test

import (
	"strings"
	"testing"

	"github.com/flexzuu/linear-tracker/pasteboard"
)

func TestCopyPaste(t *testing.T) {
	pasteboard.Copy(strings.NewReader("test-string"))
	b := &strings.Builder{}

	pasteboard.Paste(b)
	if b.String() != "test-string" {
		t.Fail()
	}
}
