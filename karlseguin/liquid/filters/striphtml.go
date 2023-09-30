package filters

import (
	"regexp"

	"github.com/unix-world/go-modules/karlseguin/liquid/core"
)

var stripHtml = &ReplacePattern{regexp.MustCompile("(?i)<script.*?</script>|<!--.*?-->|<style.*?</style>|<.*?>"), ""}

func StripHtmlFactory(parameters []core.Value) core.Filter {
	return stripHtml.Replace
}
