package liquid

import (
	"bytes"
	"testing"

	"github.com/unix-world/go-modules/karlseguin/liquid/core"
)

func TestLiteralRendersItself(t *testing.T) {
	literal := newLiteral([]byte("it's over 9001"))
	assertRender(t, literal, nil, "it's over 9001")
}

func assertRender(t *testing.T, code core.Code, d map[string]interface{}, expected string) {
	writer := new(bytes.Buffer)
	code.Execute(writer, d)
	if writer.String() != expected {
		t.Errorf("Expecting %q, got %q", expected, writer.String())
	}
}
