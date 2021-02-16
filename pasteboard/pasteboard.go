package pasteboard

import (
	"io"
	"strings"

	"github.com/progrium/macdriver/cocoa"
)

func Paste(dst io.Writer) error {
	r := strings.NewReader(cocoa.NSPasteboard_GeneralPasteboard().StringForType(cocoa.NSPasteboardTypeString))
	io.Copy(dst, r)
	return nil
}

func Copy(src io.Reader) error {
	b := &strings.Builder{}
	io.Copy(b, src)

	cocoa.NSPasteboard_GeneralPasteboard().ClearContents()
	cocoa.NSPasteboard_GeneralPasteboard().SetStringForType(b.String(), cocoa.NSPasteboardTypeString)
	return nil
}
