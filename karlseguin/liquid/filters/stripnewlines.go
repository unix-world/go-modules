package filters

import (
	"regexp"

	"github.com/unix-world/go-modules/karlseguin/liquid/core"
)

var stripNewLines = &ReplacePattern{regexp.MustCompile("(\n|\r)"), ""}

func StripNewLinesFactory(parameters []core.Value) core.Filter {
	return stripNewLines.Replace
}
