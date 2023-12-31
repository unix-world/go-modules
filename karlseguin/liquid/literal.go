package liquid

import (
	"io"

	"github.com/unix-world/go-modules/karlseguin/liquid/core"
)

type Literal struct {
	Value []byte
}

// Creates a literal (just plain text)
func newLiteral(data []byte) core.Code {
	return &Literal{Value: data}
}

func (l *Literal) Execute(writer io.Writer, data map[string]interface{}) core.ExecuteState {
	writer.Write(l.Value)
	return core.Normal
}
